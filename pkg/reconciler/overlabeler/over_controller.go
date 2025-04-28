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

package overlabeler

import (
	"context"

	"k8s.io/client-go/tools/cache"
	"k8s.io/utils/clock"

	v1 "knative.dev/serving/pkg/apis/serving/v1"
	servingclient "knative.dev/serving/pkg/client/injection/client"
	configurationinformer "knative.dev/serving/pkg/client/injection/informers/serving/v1/configuration"
	revisioninformer "knative.dev/serving/pkg/client/injection/informers/serving/v1/revision"
	routeinformer "knative.dev/serving/pkg/client/injection/informers/serving/v1/route"
	routereconciler "knative.dev/serving/pkg/client/injection/reconciler/serving/v1/route"
	"knative.dev/serving/pkg/reconciler/overconfiguration/config"

	"knative.dev/serving/pkg/overcontroller"
	"knative.dev/serving/pkg/overconfigmap"
	"knative.dev/serving/pkg/overlogging"
)

// NewController wraps a new instance of the labeler that labels
// Configurations with Routes in a controller.
func NewController(ctx context.Context, cmw overconfigmap.Watcher) *overcontroller.Impl {
	logger := overlogging.FromContext(ctx)
	routeInformer := routeinformer.Get(ctx)
	configInformer := configurationinformer.Get(ctx)
	revisionInformer := revisioninformer.Get(ctx)

	configStore := config.NewStore(logger.Named("config-store"))
	configStore.WatchConfigs(cmw)

	c := &Reconciler{}
	_ = c.ReconcileKind
	impl := routereconciler.NewImpl(ctx, c, func(*overcontroller.Impl) overcontroller.Options {
		return overcontroller.Options{
			ConfigStore: configStore,
			// The labeler shouldn't mutate the route's status.
			SkipStatusUpdates: true,
		}
	})

	routeInformer.Informer().AddEventHandler(overcontroller.HandleAll(impl.Enqueue))

	// Make sure trackers are deleted once the observers are removed.
	routeInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		DeleteFunc: impl.Tracker.OnDeletedObserver,
	})

	configInformer.Informer().AddEventHandler(overcontroller.HandleAll(
		overcontroller.EnsureTypeMeta(
			impl.Tracker.OnChanged,
			v1.SchemeGroupVersion.WithKind("Configuration"),
		),
	))

	revisionInformer.Informer().AddEventHandler(overcontroller.HandleAll(
		overcontroller.EnsureTypeMeta(
			impl.Tracker.OnChanged,
			v1.SchemeGroupVersion.WithKind("Revision"),
		),
	))

	client := servingclient.Get(ctx)
	clock := &clock.RealClock{}
	c.caccV2 = newConfigurationAccessor(client, impl.Tracker, configInformer.Lister(), configInformer.Informer().GetIndexer(), clock)
	c.raccV2 = newRevisionAccessor(client, impl.Tracker, revisionInformer.Lister(), revisionInformer.Informer().GetIndexer(), clock)

	return impl
}
