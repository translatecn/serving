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

package overmetric

import (
	"context"

	"knative.dev/serving/pkg/autoscaler/metrics"
	metricinformer "knative.dev/serving/pkg/client/injection/informers/autoscaling/v1alpha1/metric"
	metricreconciler "knative.dev/serving/pkg/client/injection/reconciler/autoscaling/v1alpha1/metric"

	"knative.dev/serving/pkg/overcontroller"
	"knative.dev/serving/pkg/overconfigmap"
)

// NewController initializes the controller and is called by the generated code.
// Registers eventhandlers to enqueue events.
func NewController(
	ctx context.Context,
	cmw overconfigmap.Watcher,
	collector metrics.Collector,
) *overcontroller.Impl {
	metricInformer := metricinformer.Get(ctx)

	c := &reconciler{
		collector: collector,
	}
	impl := metricreconciler.NewImpl(ctx, c)

	// Watch all the Metric objects.
	metricInformer.Informer().AddEventHandler(overcontroller.HandleAll(impl.Enqueue))

	collector.Watch(impl.EnqueueKey)

	return impl
}
