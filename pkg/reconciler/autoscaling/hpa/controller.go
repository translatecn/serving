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

package hpa

import (
	"context"

	networkingclient "knative.dev/serving/networking/pkg/client/injection/client"
	sksinformer "knative.dev/serving/networking/pkg/client/injection/informers/networking/v1alpha1/serverlessservice"
	servingclient "knative.dev/serving/pkg/client/injection/client"
	metricinformer "knative.dev/serving/pkg/client/injection/informers/autoscaling/v1alpha1/metric"
	painformer "knative.dev/serving/pkg/client/injection/informers/autoscaling/v1alpha1/podautoscaler"
	kubeclient "knative.dev/serving/pkg/client/injection/kube/client"
	hpainformer "knative.dev/serving/pkg/client/injection/kube/informers/autoscaling/v2/horizontalpodautoscaler"
	pareconciler "knative.dev/serving/pkg/client/injection/reconciler/autoscaling/v1alpha1/podautoscaler"
	"knative.dev/serving/pkg/overdeployment"
	"knative.dev/serving/pkg/overlogging"

	"k8s.io/client-go/tools/cache"
	"knative.dev/serving/pkg/apis/autoscaling"
	autoscalingv1alpha1 "knative.dev/serving/pkg/apis/autoscaling/v1alpha1"
	"knative.dev/serving/pkg/autoscaler/config/autoscalerconfig"
	"knative.dev/serving/pkg/overcontroller"
	"knative.dev/serving/pkg/overconfigmap"
	pkgreconciler "knative.dev/serving/pkg/reconciler"
	areconciler "knative.dev/serving/pkg/reconciler/autoscaling"
	"knative.dev/serving/pkg/reconciler/autoscaling/config"
)

// NewController returns a new HPA reconcile controller.
func NewController(
	ctx context.Context,
	cmw overconfigmap.Watcher,
) *overcontroller.Impl {
	logger := overlogging.FromContext(ctx)
	paInformer := painformer.Get(ctx)
	sksInformer := sksinformer.Get(ctx)
	hpaInformer := hpainformer.Get(ctx)
	metricInformer := metricinformer.Get(ctx)

	onlyHPAClass := pkgreconciler.AnnotationFilterFunc(autoscaling.ClassAnnotationKey, autoscaling.HPA, false)

	c := &Reconciler{
		Base: &areconciler.Base{
			Client:           servingclient.Get(ctx),
			NetworkingClient: networkingclient.Get(ctx),
			SKSLister:        sksInformer.Lister(),
			MetricLister:     metricInformer.Lister(),
		},

		kubeClient: kubeclient.Get(ctx),
		hpaLister:  hpaInformer.Lister(),
	}
	_ = c.ReconcileKind

	impl := pareconciler.NewImpl(ctx, c, autoscaling.HPA, func(impl *overcontroller.Impl) overcontroller.Options {
		logger.Info("Setting up ConfigMap receivers")
		configsToResync := []interface{}{
			&autoscalerconfig.Config{},
			&overdeployment.Config{},
		}
		resync := overconfigmap.TypeFilter(configsToResync...)(func(string, interface{}) {
			impl.FilteredGlobalResync(onlyHPAClass, paInformer.Informer())
		})
		configStore := config.NewStore(logger.Named("config-store"), resync)
		configStore.WatchConfigs(cmw)
		return overcontroller.Options{ConfigStore: configStore}
	})

	logger.Info("Setting up hpa-class event handlers")

	paInformer.Informer().AddEventHandler(cache.FilteringResourceEventHandler{
		FilterFunc: onlyHPAClass,
		Handler:    overcontroller.HandleAll(impl.Enqueue),
	})

	onlyPAControlled := overcontroller.FilterController(&autoscalingv1alpha1.PodAutoscaler{})

	hpaInformer.Informer().AddEventHandler(cache.FilteringResourceEventHandler{
		FilterFunc: pkgreconciler.ChainFilterFuncs(onlyHPAClass, onlyPAControlled),
		Handler:    overcontroller.HandleAll(impl.EnqueueControllerOf),
	})
	sksInformer.Informer().AddEventHandler(cache.FilteringResourceEventHandler{
		FilterFunc: pkgreconciler.ChainFilterFuncs(onlyHPAClass, onlyPAControlled),
		Handler:    overcontroller.HandleAll(impl.EnqueueControllerOf),
	})
	metricInformer.Informer().AddEventHandler(cache.FilteringResourceEventHandler{
		FilterFunc: pkgreconciler.ChainFilterFuncs(onlyHPAClass, onlyPAControlled),
		Handler:    overcontroller.HandleAll(impl.EnqueueControllerOf),
	})

	return impl
}
