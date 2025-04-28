/*
Copyright 2019 The Knative Authors

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

package overkpa

import (
	"context"

	"k8s.io/client-go/tools/cache"

	networkingclient "knative.dev/serving/networking/pkg/client/injection/client"
	sksinformer "knative.dev/serving/networking/pkg/client/injection/informers/networking/v1alpha1/serverlessservice"
	servingclient "knative.dev/serving/pkg/client/injection/client"
	"knative.dev/serving/pkg/client/injection/ducks/autoscaling/v1alpha1/podscalable"
	metricinformer "knative.dev/serving/pkg/client/injection/informers/autoscaling/v1alpha1/metric"
	painformer "knative.dev/serving/pkg/client/injection/informers/autoscaling/v1alpha1/podautoscaler"
	filteredpodinformer "knative.dev/serving/pkg/client/injection/kube/informers/core/v1/pod/filtered"
	pareconciler "knative.dev/serving/pkg/client/injection/reconciler/autoscaling/v1alpha1/podautoscaler"

	"knative.dev/serving/pkg/apis/autoscaling"
	autoscalingv1alpha1 "knative.dev/serving/pkg/apis/autoscaling/v1alpha1"
	"knative.dev/serving/pkg/apis/serving"
	"knative.dev/serving/pkg/autoscaler/config/autoscalerconfig"
	"knative.dev/serving/pkg/overcontroller"
	"knative.dev/serving/pkg/overdeployment"
	"knative.dev/serving/pkg/overconfigmap"
	"knative.dev/serving/pkg/overlogging"
	pkgreconciler "knative.dev/serving/pkg/reconciler"
	areconciler "knative.dev/serving/pkg/reconciler/autoscaling"
	"knative.dev/serving/pkg/reconciler/autoscaling/config"
	"knative.dev/serving/pkg/reconciler/autoscaling/overkpa/overresources"
)

func NewController(ctx context.Context, cmw overconfigmap.Watcher, deciders overresources.Deciders) *overcontroller.Impl {
	logger := overlogging.FromContext(ctx)
	paInformer := painformer.Get(ctx)
	sksInformer := sksinformer.Get(ctx)
	podsInformer := filteredpodinformer.Get(ctx, serving.RevisionUID)
	metricInformer := metricinformer.Get(ctx)
	psInformerFactory := podscalable.Get(ctx)

	onlyKPAClass := pkgreconciler.AnnotationFilterFunc(autoscaling.ClassAnnotationKey, autoscaling.KPA, false /*allowUnset*/) // autoscaling.knative.dev/class
	//kpa.autoscaling.knative.dev
	c := &Reconciler{
		Base: &areconciler.Base{
			Client:           servingclient.Get(ctx),
			NetworkingClient: networkingclient.Get(ctx),
			SKSLister:        sksInformer.Lister(),
			MetricLister:     metricInformer.Lister(),
		},
		podsLister: podsInformer.Lister(),
		deciders:   deciders,
	}
	_ = c.ReconcileKind

	impl := pareconciler.NewImpl(ctx, c, autoscaling.KPA, func(impl *overcontroller.Impl) overcontroller.Options {
		logger.Info("Setting up ConfigMap receivers")
		configsToResync := []interface{}{
			&autoscalerconfig.Config{},
			&overdeployment.Config{},
		}
		resync := overconfigmap.TypeFilter(configsToResync...)(func(string, interface{}) {
			impl.FilteredGlobalResync(onlyKPAClass, paInformer.Informer())
		})
		configStore := config.NewStore(logger.Named("config-store"), resync)
		configStore.WatchConfigs(cmw)
		return overcontroller.Options{ConfigStore: configStore}
	})

	c.scaler = newScaler(ctx, psInformerFactory, impl.EnqueueAfter)

	logger.Info("Setting up KPA-Class event handlers")

	paInformer.Informer().AddEventHandler(cache.FilteringResourceEventHandler{
		FilterFunc: onlyKPAClass,
		Handler:    overcontroller.HandleAll(impl.Enqueue),
	})
	// PodAutoscalers、ServerlessServices、Metrics、Pod
	onlyPAControlled := overcontroller.FilterController(&autoscalingv1alpha1.PodAutoscaler{})

	sksInformer.Informer().AddEventHandler(cache.FilteringResourceEventHandler{
		FilterFunc: pkgreconciler.ChainFilterFuncs(onlyKPAClass, onlyPAControlled),
		Handler:    overcontroller.HandleAll(impl.EnqueueControllerOf),
	})

	metricInformer.Informer().AddEventHandler(cache.FilteringResourceEventHandler{
		FilterFunc: pkgreconciler.ChainFilterFuncs(onlyKPAClass, onlyPAControlled),
		Handler:    overcontroller.HandleAll(impl.EnqueueControllerOf),
	})

	// Watch the knative pods.
	podsInformer.Informer().AddEventHandler(cache.FilteringResourceEventHandler{
		FilterFunc: pkgreconciler.LabelExistsFilterFunc(serving.RevisionLabelKey),
		Handler:    overcontroller.HandleAll(impl.EnqueueLabelOfNamespaceScopedResource("", serving.RevisionLabelKey)),
	})

	// Have the Deciders enqueue the PAs whose decisions have changed.
	deciders.Watch(impl.EnqueueKey)

	return impl
}
