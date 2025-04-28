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

package overconfiguration

import (
	"context"

	"k8s.io/client-go/tools/cache"
	"k8s.io/utils/clock"
	v1 "knative.dev/serving/pkg/apis/serving/v1"
	servingclient "knative.dev/serving/pkg/client/injection/client"
	configurationinformer "knative.dev/serving/pkg/client/injection/informers/serving/v1/configuration"
	revisioninformer "knative.dev/serving/pkg/client/injection/informers/serving/v1/revision"
	configreconciler "knative.dev/serving/pkg/client/injection/reconciler/serving/v1/configuration"
	"knative.dev/serving/pkg/overcontroller"
	"knative.dev/serving/pkg/overconfigmap"
	"knative.dev/serving/pkg/overlogging"
	"knative.dev/serving/pkg/reconciler/overconfiguration/config"
)

// NewController creates a new Configuration controller
func NewController(
	ctx context.Context,
	cmw overconfigmap.Watcher,
) *overcontroller.Impl {
	logger := overlogging.FromContext(ctx)
	configurationInformer := configurationinformer.Get(ctx)
	revisionInformer := revisioninformer.Get(ctx)

	configStore := config.NewStore(logger.Named("config-store"))
	configStore.WatchConfigs(cmw)

	c := &Reconciler{
		client:         servingclient.Get(ctx),
		revisionLister: revisionInformer.Lister(),
		clock:          &clock.RealClock{},
	}
	_ = c.ReconcileKind
	impl := configreconciler.NewImpl(ctx, c, func(*overcontroller.Impl) overcontroller.Options {
		return overcontroller.Options{ConfigStore: configStore}
	})

	configurationInformer.Informer().AddEventHandler(overcontroller.HandleAll(impl.Enqueue))

	revisionInformer.Informer().AddEventHandler(cache.FilteringResourceEventHandler{
		FilterFunc: overcontroller.FilterController(&v1.Configuration{}),
		Handler:    overcontroller.HandleAll(impl.EnqueueControllerOf),
	})

	return impl
}
