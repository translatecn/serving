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

package overnscert

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/cache"
	"knative.dev/serving/networking/pkg/client/injection/client"
	kcertinformer "knative.dev/serving/networking/pkg/client/injection/informers/networking/v1alpha1/certificate"
	nsinformer "knative.dev/serving/pkg/client/injection/kube/informers/core/v1/namespace"
	namespacereconciler "knative.dev/serving/pkg/client/injection/kube/reconciler/core/v1/namespace"
	"knative.dev/serving/pkg/overcontroller"
	"knative.dev/serving/pkg/overconfigmap"
	"knative.dev/serving/pkg/overlogging"
	routecfg "knative.dev/serving/pkg/reconciler/overroute/overconfig"

	netcfg "knative.dev/serving/networking/pkg/config"
	"knative.dev/serving/pkg/reconciler/overnscert/config"
)

// NewController initializes the controller and is called by the generated code
// Registers eventhandlers to enqueue events.
func NewController(ctx context.Context, cmw overconfigmap.Watcher) *overcontroller.Impl {
	logger := overlogging.FromContext(ctx)
	nsInformer := nsinformer.Get(ctx)
	knCertificateInformer := kcertinformer.Get(ctx)

	c := &reconciler{
		client:              client.Get(ctx),
		knCertificateLister: knCertificateInformer.Lister(),
	}

	_ = c.ReconcileKind
	impl := namespacereconciler.NewImpl(ctx, c, func(impl *overcontroller.Impl) overcontroller.Options {
		nsInformer.Informer().AddEventHandler(overcontroller.HandleAll(impl.Enqueue))

		knCertificateInformer.Informer().AddEventHandler(cache.FilteringResourceEventHandler{
			FilterFunc: overcontroller.FilterControllerGK(corev1.SchemeGroupVersion.WithKind("Namespace").GroupKind()),
			Handler:    overcontroller.HandleAll(impl.EnqueueControllerOf),
		})

		configsToResync := []interface{}{
			&netcfg.Config{},
			&routecfg.Domain{},
		}
		resync := overconfigmap.TypeFilter(configsToResync...)(func(string, interface{}) {
			impl.GlobalResync(nsInformer.Informer())
		})
		configStore := config.NewStore(logger.Named("config-store"), resync)
		configStore.WatchConfigs(cmw)
		return overcontroller.Options{ConfigStore: configStore}
	})

	return impl
}
