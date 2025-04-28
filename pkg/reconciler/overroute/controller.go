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

package overroute

import (
	"context"

	netclient "knative.dev/serving/networking/pkg/client/injection/client"
	certificateinformer "knative.dev/serving/networking/pkg/client/injection/informers/networking/v1alpha1/certificate"
	ingressinformer "knative.dev/serving/networking/pkg/client/injection/informers/networking/v1alpha1/ingress"
	servingclient "knative.dev/serving/pkg/client/injection/client"
	configurationinformer "knative.dev/serving/pkg/client/injection/informers/serving/v1/configuration"
	revisioninformer "knative.dev/serving/pkg/client/injection/informers/serving/v1/revision"
	routeinformer "knative.dev/serving/pkg/client/injection/informers/serving/v1/route"
	kubeclient "knative.dev/serving/pkg/client/injection/kube/client"
	endpointsinformer "knative.dev/serving/pkg/client/injection/kube/informers/core/v1/endpoints"
	serviceinformer "knative.dev/serving/pkg/client/injection/kube/informers/core/v1/service"
	routereconciler "knative.dev/serving/pkg/client/injection/reconciler/serving/v1/route"

	"k8s.io/client-go/tools/cache"
	"k8s.io/utils/clock"
	netcfg "knative.dev/serving/networking/pkg/config"
	v1 "knative.dev/serving/pkg/apis/serving/v1"
	"knative.dev/serving/pkg/overcontroller"
	"knative.dev/serving/pkg/overconfigmap"
	"knative.dev/serving/pkg/overlogging"
	"knative.dev/serving/pkg/reconciler/overroute/overconfig"
)

// NewController initializes the controller and is called by the generated code
// Registers eventhandlers to enqueue events
func NewController(
	ctx context.Context,
	cmw overconfigmap.Watcher,
) *overcontroller.Impl {
	return newController(ctx, cmw, clock.RealClock{})
}

type reconcilerOption func(*Reconciler)

func newController(
	ctx context.Context,
	cmw overconfigmap.Watcher,
	clock clock.Clock,
	opts ...reconcilerOption,
) *overcontroller.Impl {
	logger := overlogging.FromContext(ctx)
	serviceInformer := serviceinformer.Get(ctx)
	endpointsInformer := endpointsinformer.Get(ctx)
	routeInformer := routeinformer.Get(ctx)
	configInformer := configurationinformer.Get(ctx)
	revisionInformer := revisioninformer.Get(ctx)
	ingressInformer := ingressinformer.Get(ctx)
	certificateInformer := certificateinformer.Get(ctx)

	c := &Reconciler{
		kubeclient:          kubeclient.Get(ctx),
		client:              servingclient.Get(ctx),
		netclient:           netclient.Get(ctx),
		configurationLister: configInformer.Lister(),
		revisionLister:      revisionInformer.Lister(),
		serviceLister:       serviceInformer.Lister(),
		endpointsLister:     endpointsInformer.Lister(),
		ingressLister:       ingressInformer.Lister(),
		certificateLister:   certificateInformer.Lister(),
		clock:               clock,
	}
	_ = c.ReconcileKind
	impl := routereconciler.NewImpl(ctx, c, func(impl *overcontroller.Impl) overcontroller.Options {
		configsToResync := []interface{}{
			&netcfg.Config{},
			&overconfig.Domain{},
		}
		resync := overconfigmap.TypeFilter(configsToResync...)(func(string, interface{}) {
			impl.GlobalResync(routeInformer.Informer())
		})
		configStore := overconfig.NewStore(overlogging.WithLogger(ctx, logger.Named("config-store")), resync)
		configStore.WatchConfigs(cmw)
		return overcontroller.Options{ConfigStore: configStore}
	})
	c.enqueueAfter = impl.EnqueueAfter

	routeInformer.Informer().AddEventHandler(overcontroller.HandleAll(impl.Enqueue))

	serviceInformer.Informer().AddEventHandler(cache.FilteringResourceEventHandler{
		FilterFunc: overcontroller.FilterController(&v1.Route{}),
		Handler:    overcontroller.HandleAll(impl.EnqueueControllerOf),
	})
	certificateInformer.Informer().AddEventHandler(cache.FilteringResourceEventHandler{
		FilterFunc: overcontroller.FilterController(&v1.Route{}),
		Handler:    overcontroller.HandleAll(impl.EnqueueControllerOf),
	})
	ingressInformer.Informer().AddEventHandler(cache.FilteringResourceEventHandler{
		FilterFunc: overcontroller.FilterController(&v1.Route{}),
		Handler:    overcontroller.HandleAll(impl.EnqueueControllerOf),
	})

	c.tracker = impl.Tracker

	// Make sure trackers are deleted once the observers are removed.
	routeInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		DeleteFunc: c.tracker.OnDeletedObserver,
	})

	configInformer.Informer().AddEventHandler(overcontroller.HandleAll(
		// Call the tracker's OnChanged method, but we've seen the objects
		// coming through this path missing TypeMeta, so ensure it is properly
		// populated.
		overcontroller.EnsureTypeMeta(
			c.tracker.OnChanged,
			v1.SchemeGroupVersion.WithKind("Configuration"),
		),
	))

	revisionInformer.Informer().AddEventHandler(overcontroller.HandleAll(
		// Call the tracker's OnChanged method, but we've seen the objects
		// coming through this path missing TypeMeta, so ensure it is properly
		// populated.
		overcontroller.EnsureTypeMeta(
			c.tracker.OnChanged,
			v1.SchemeGroupVersion.WithKind("Revision"),
		),
	))

	for _, opt := range opts {
		opt(c)
	}
	return impl
}
