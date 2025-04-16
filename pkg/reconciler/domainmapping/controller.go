/*
Copyright 2020 The Knative Authors

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

package domainmapping

import (
	"context"

	"k8s.io/client-go/tools/cache"
	netclient "knative.dev/serving/networking/pkg/client/injection/client"
	certificateinformer "knative.dev/serving/networking/pkg/client/injection/informers/networking/v1alpha1/certificate"
	domainclaiminformer "knative.dev/serving/networking/pkg/client/injection/informers/networking/v1alpha1/clusterdomainclaim"
	ingressinformer "knative.dev/serving/networking/pkg/client/injection/informers/networking/v1alpha1/ingress"
	netcfg "knative.dev/serving/networking/pkg/config"
	"knative.dev/serving/pkg/apis/serving/v1beta1"
	"knative.dev/serving/pkg/client/injection/informers/serving/v1beta1/domainmapping"
	kindreconciler "knative.dev/serving/pkg/client/injection/reconciler/serving/v1beta1/domainmapping"
	"knative.dev/serving/pkg/configmap"
	"knative.dev/serving/pkg/controller"
	"knative.dev/serving/pkg/over_logging"
	"knative.dev/serving/pkg/reconciler/domainmapping/config"
	"knative.dev/serving/pkg/resolver"
)

// NewController creates a new DomainMapping controller.
func NewController(ctx context.Context, cmw configmap.Watcher) *controller.Impl {
	logger := over_logging.FromContext(ctx)
	certificateInformer := certificateinformer.Get(ctx)
	domainmappingInformer := domainmapping.Get(ctx)
	ingressInformer := ingressinformer.Get(ctx)
	domainClaimInformer := domainclaiminformer.Get(ctx)

	r := &Reconciler{
		certificateLister: certificateInformer.Lister(),
		ingressLister:     ingressInformer.Lister(),
		domainClaimLister: domainClaimInformer.Lister(),
		netclient:         netclient.Get(ctx),
	}

	_ = r.ReconcileKind
	impl := kindreconciler.NewImpl(ctx, r, func(impl *controller.Impl) controller.Options {
		configsToResync := []interface{}{
			&netcfg.Config{},
		}
		resync := configmap.TypeFilter(configsToResync...)(func(string, interface{}) {
			impl.GlobalResync(domainmappingInformer.Informer())
		})
		configStore := config.NewStore(over_logging.WithLogger(ctx, logger.Named("config-store")), resync)
		configStore.WatchConfigs(cmw)
		return controller.Options{ConfigStore: configStore}
	})

	domainmappingInformer.Informer().AddEventHandler(controller.HandleAll(impl.Enqueue))

	certificateInformer.Informer().AddEventHandler(cache.FilteringResourceEventHandler{
		FilterFunc: controller.FilterController(&v1beta1.DomainMapping{}),
		Handler:    controller.HandleAll(impl.EnqueueControllerOf),
	})
	ingressInformer.Informer().AddEventHandler(cache.FilteringResourceEventHandler{
		FilterFunc: controller.FilterController(&v1beta1.DomainMapping{}),
		Handler:    controller.HandleAll(impl.EnqueueControllerOf),
	})

	r.resolver = resolver.NewURIResolverFromTracker(ctx, impl.Tracker)

	return impl
}
