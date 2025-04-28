/*
Copyright 2018 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Multitenant autoscaler executable.
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	corev1listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/rest"
	kubeclient "knative.dev/serving/pkg/client/injection/kube/client"

	netcfg "knative.dev/serving/networking/pkg/config"
	autoscalingv1alpha1 "knative.dev/serving/pkg/apis/autoscaling/v1alpha1"
	"knative.dev/serving/pkg/apis/serving"
	"knative.dev/serving/pkg/autoscaler/bucket"
	asmetrics "knative.dev/serving/pkg/autoscaler/metrics"
	"knative.dev/serving/pkg/autoscaler/overstatserver"
	"knative.dev/serving/pkg/autoscaler/scaling"
	"knative.dev/serving/pkg/autoscaler/statforwarder"
	filteredpodinformer "knative.dev/serving/pkg/client/injection/kube/informers/core/v1/pod/filtered"
	filteredinformerfactory "knative.dev/serving/pkg/client/injection/kube/informers/factory/filtered"
	"knative.dev/serving/pkg/overcontroller"
	"knative.dev/serving/pkg/injection"
	"knative.dev/serving/pkg/injection/sharedmain"
	"knative.dev/serving/pkg/overleaderelection"
	"knative.dev/serving/pkg/metrics"
	smetrics "knative.dev/serving/pkg/metrics"
	configmap "knative.dev/serving/pkg/overconfigmap/informer"
	"knative.dev/serving/pkg/overlogging"
	"knative.dev/serving/pkg/overprofiling"
	"knative.dev/serving/pkg/oversignals"
	"knative.dev/serving/pkg/oversystem"
	"knative.dev/serving/pkg/overversion"
	"knative.dev/serving/pkg/reconciler/autoscaling/overkpa"
	"knative.dev/serving/pkg/reconciler/overmetric"
	"knative.dev/serving/pkg/resources"
)

const (
	statsServerAddr = ":8080"
	statsBufferLen  = 1000
	component       = "autoscaler"
	controllerNum   = 2
)

func main() {
	// Set up signals so we handle the first shutdown signal gracefully.
	ctx := oversignals.NewContext()

	// Report stats on Go memory usage every 30 seconds.
	metrics.MemStatsOrDie(ctx)

	cfg := injection.ParseAndGetRESTConfigOrDie()

	log.Printf("Registering %d clients", len(injection.Default.GetClients()))                      // 4
	log.Printf("Registering %d informer factories", len(injection.Default.GetInformerFactories())) // 5
	log.Printf("Registering %d informers", len(injection.Default.GetInformers()))                  // 5
	log.Printf("Registering %d filtered informers", len(injection.Default.GetFilteredInformers())) // 1
	log.Printf("Registering %d controllers", controllerNum)                                        // 2

	// Adjust our client's rate limits based on the number of controller's we are running.
	cfg.QPS = controllerNum * rest.DefaultQPS
	cfg.Burst = controllerNum * rest.DefaultBurst
	ctx = filteredinformerfactory.WithSelectors(ctx, serving.RevisionUID)
	ctx, informers := injection.Default.SetupInformers(ctx, cfg)

	kubeClient := kubeclient.Get(ctx)

	// We sometimes startup faster than we can reach kube-api. Poll on failure to prevent us terminating
	var err error
	if perr := wait.PollUntilContextTimeout(ctx, time.Second, 60*time.Second, true, func(context.Context) (bool, error) {
		if err = overversion.CheckMinimumVersion(kubeClient.Discovery()); err != nil {
			log.Print("Failed to get k8s version ", err)
		}
		return err == nil, nil
	}); perr != nil {
		log.Fatal("Timed out attempting to get k8s version: ", err)
	}

	// Set up our logger.
	loggingConfig, err := sharedmain.GetLoggingConfig(ctx)

	if err != nil {
		log.Fatal("Error loading/parsing logging configuration: ", err)
	}
	loggingConfig.LoggingLevel[component] = zapcore.DebugLevel
	logger, atomicLevel := overlogging.NewLoggerFromConfig(loggingConfig, component)
	defer flush(logger)
	ctx = overlogging.WithLogger(ctx, logger)

	// statsCh is the main communication channel between the stats server and multiscaler.
	statsCh := make(chan asmetrics.StatMessage, statsBufferLen)
	defer close(statsCh)

	profilingHandler := overprofiling.NewHandler(logger, false)

	cmw := configmap.NewInformedWatcher(kubeclient.Get(ctx), oversystem.Namespace())
	// Watch the logging config map and dynamically update logging levels.
	cmw.Watch(overlogging.ConfigMapName(), overlogging.UpdateLevelFromConfigMap(logger, atomicLevel, component)) // ✅
	// Watch the observability config map
	cmw.Watch(
		metrics.ConfigMapName(),
		metrics.ConfigMapWatcher(ctx, component, nil /* SecretFetcher */, logger), // ✅
		profilingHandler.UpdateFromConfigMap,                                      // ✅
	)

	podLister := filteredpodinformer.Get(ctx, serving.RevisionUID).Lister()
	networkCM, err := kubeclient.Get(ctx).CoreV1().ConfigMaps(oversystem.Namespace()).Get(ctx, netcfg.ConfigMapName, metav1.GetOptions{})
	if err != nil {
		logger.Fatalw("Failed to fetch network config", zap.Error(err))
	}
	networkConfig, err := netcfg.NewConfigFromMap(networkCM.Data)
	if err != nil {
		logger.Fatalw("Failed to construct network config", zap.Error(err))
	}

	collector := asmetrics.NewMetricCollector(
		statsScraperFactoryFunc(podLister,
			networkConfig.EnableMeshPodAddressability, //false
			networkConfig.MeshCompatibilityMode,       // auto
		),
		logger,
	)

	// Set up scalers.
	multiScaler := scaling.NewMultiScaler(ctx.Done(), uniScalerFactoryFunc(podLister, collector), logger)

	controllers := []*overcontroller.Impl{
		overkpa.NewController(ctx, cmw, multiScaler), // 这里很重要 ✈️ ✈️ ✈️ ✈️ ✈️ ✈️ ✈️ ✈️ ✈️
		overmetric.NewController(ctx, cmw, collector),
	}

	// Start watching the configs.
	if err := cmw.Start(ctx.Done()); err != nil {
		logger.Fatalw("Failed to start watching configs", zap.Error(err))
	}

	// Start all of the informers and wait for them to sync.
	if err := overcontroller.StartInformers(ctx.Done(), informers...); err != nil {
		logger.Fatalw("Failed to start informers", zap.Error(err))
	}

	// 当此 Pod 拥有此 StatMessage 的修订版本时，调用 accept 函数。
	accept := func(sm asmetrics.StatMessage) {
		collector.Record(sm.Key, time.Unix(sm.Stat.Timestamp, 0), sm.Stat)
		multiScaler.Poke(sm.Key, sm.Stat)
	}

	cc := componentConfigAndIP(ctx)

	// We don't want an elector on the controller context
	// since they will be sharing an elector
	var electorCtx context.Context

	var f *statforwarder.Forwarder
	if b, bs, err := overleaderelection.NewStatefulSetBucketAndSet(int(cc.Buckets)); err == nil {
		logger.Info("Running with StatefulSet leader election")
		electorCtx = overleaderelection.WithStatefulSetElectorBuilder(ctx, cc, b)
		f = statforwarder.New(ctx, bs)
		if err := statforwarder.StatefulSetBasedProcessor(ctx, f, accept); err != nil {
			logger.Fatalw("Failed to set up statefulset processors", zap.Error(err))
		}
	} else {
		logger.Info("Running with Standard leader election")
		electorCtx = overleaderelection.WithStandardLeaderElectorBuilder(ctx, kubeClient, cc)
		f = statforwarder.New(ctx, bucket.AutoscalerBucketSet(cc.Buckets))
		if err := statforwarder.LeaseBasedProcessor(ctx, f, accept); err != nil {
			logger.Fatalw("Failed to set up lease tracking", zap.Error(err))
		}
	}

	elector, err := setupSharedElector(electorCtx, controllers) // 也很重要
	if err != nil {
		logger.Fatalw("Failed to setup elector", zap.Error(err))
	}

	// Set up a statserver.
	statsServer := overstatserver.New(statsServerAddr, statsCh, logger, f.IsBucketOwner) // ✅
	defer f.Cancel()

	go func() {
		for sm := range statsCh {
			// Set the timestamp when first receiving the stat.
			if sm.Stat.Timestamp == 0 {
				sm.Stat.Timestamp = time.Now().Unix()
			}
			f.Process(sm) // 重要
		}
	}()

	profilingServer := overprofiling.NewServer(profilingHandler) // ✅

	eg, egCtx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		elector.Run(egCtx)
		return nil
	})
	eg.Go(statsServer.ListenAndServe)
	eg.Go(profilingServer.ListenAndServe)
	eg.Go(func() error {
		return overcontroller.StartAll(egCtx, controllers...)
	})

	// This will block until either a signal arrives or one of the grouped functions
	// returns an error.
	<-egCtx.Done()

	statsServer.Shutdown(5 * time.Second)
	profilingServer.Shutdown(context.Background())
	// Don't forward ErrServerClosed as that indicates we're already shutting down.
	if err := eg.Wait(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Errorw("Error while running server", zap.Error(err))
	}
}

func flush(logger *zap.SugaredLogger) {
	logger.Sync()
	metrics.FlushExporter()
}

func statsScraperFactoryFunc(podLister corev1listers.PodLister, usePassthroughLb bool, meshMode netcfg.MeshCompatibilityMode) asmetrics.StatsScraperFactory {
	return func(metric *autoscalingv1alpha1.Metric, logger *zap.SugaredLogger) (asmetrics.StatsScraper, error) {
		if metric.Spec.ScrapeTarget == "" {
			return nil, nil
		}

		revisionName := metric.Labels[serving.RevisionLabelKey]
		if revisionName == "" {
			return nil, fmt.Errorf("label %q not found or empty in Metric %s", serving.RevisionLabelKey, metric.Name)
		}

		podAccessor := resources.NewPodAccessor(podLister, metric.Namespace, revisionName)
		return asmetrics.NewStatsScraper(metric, revisionName, podAccessor,
			usePassthroughLb, meshMode, logger), nil
	}
}

func uniScalerFactoryFunc(podLister corev1listers.PodLister,
	metricClient asmetrics.MetricClient,
) scaling.UniScalerFactory {
	return func(decider *scaling.ReversionReplicasByLoad) (scaling.UniScaler, error) {
		// decider 是对 pa 的 封装
		configName := decider.Labels[serving.ConfigurationLabelKey]
		if configName == "" {
			return nil, fmt.Errorf("label %q not found or empty in ReversionReplicasByLoad %s", serving.ConfigurationLabelKey, decider.Name)
		}
		revisionName := decider.Labels[serving.RevisionLabelKey]
		if revisionName == "" {
			return nil, fmt.Errorf("label %q not found or empty in ReversionReplicasByLoad %s", serving.RevisionLabelKey, decider.Name)
		}
		serviceName := decider.Labels[serving.ServiceLabelKey] // This can be empty.

		// Create a stats reporter which tags statistics by PA namespace, configuration name, and PA name.
		ctx := smetrics.RevisionContext(decider.Namespace, serviceName, configName, revisionName)

		podAccessor := resources.NewPodAccessor(podLister, decider.Namespace, revisionName)
		return scaling.New(ctx, decider.Namespace, decider.Name, metricClient, podAccessor, &decider.Spec), nil
	}
}

func componentConfigAndIP(ctx context.Context) overleaderelection.ComponentConfig {
	id, err := bucket.Identity()
	if err != nil {
		overlogging.FromContext(ctx).Fatalw("Failed to generate Lease holder identity", zap.Error(err))
	}

	// Set up leader election config
	leaderElectionConfig, err := sharedmain.GetLeaderElectionConfig(ctx)
	if err != nil {
		overlogging.FromContext(ctx).Fatalw("Error loading leader election configuration", zap.Error(err))
	}

	cc := leaderElectionConfig.GetComponentConfig(component)
	cc.LeaseName = func(i uint32) string {
		return bucket.AutoscalerBucketName(i, cc.Buckets)
	}
	cc.Identity = id

	return cc
}
