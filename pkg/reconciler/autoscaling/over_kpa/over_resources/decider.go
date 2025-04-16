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

package over_resources

import (
	"context"
	"knative.dev/serving/pkg/reconciler/autoscaling/over_resources"

	"k8s.io/apimachinery/pkg/types"
	autoscalingv1alpha1 "knative.dev/serving/pkg/apis/autoscaling/v1alpha1"
	"knative.dev/serving/pkg/autoscaler/config/autoscalerconfig"
	"knative.dev/serving/pkg/autoscaler/scaling"
)

// Deciders is an interface for notifying the presence or absence of autoscaling deciders.
type Deciders interface {
	// Get accesses the ReversionReplicasByLoad resource for this key, returning any errors.
	Get(ctx context.Context, namespace, name string) (*scaling.ReversionReplicasByLoad, error)

	// Create adds a ReversionReplicasByLoad resource for a given key, returning any errors.
	Create(ctx context.Context, decider *scaling.ReversionReplicasByLoad) (*scaling.ReversionReplicasByLoad, error)

	// Delete removes the ReversionReplicasByLoad resource for a given key, returning any errors.
	Delete(ctx context.Context, namespace, name string)

	// Watch registers a function to call when ReversionReplicasByLoad change.
	Watch(watcher func(types.NamespacedName))

	// Update update the ReversionReplicasByLoad resource, return the new ReversionReplicasByLoad or any errors.
	Update(ctx context.Context, decider *scaling.ReversionReplicasByLoad) (*scaling.ReversionReplicasByLoad, error)
}

// GetInitialScale returns the calculated initial scale based on the autoscaler
// ConfigMap and PA initial scale annotation value.
func GetInitialScale(asConfig *autoscalerconfig.Config, pa *autoscalingv1alpha1.PodAutoscaler) int32 {
	initialScale := asConfig.InitialScale
	revisionInitialScale, ok := pa.InitialScale()
	if !ok || (revisionInitialScale == 0 && !asConfig.AllowZeroInitialScale) {
		return initialScale
	}
	return revisionInitialScale
}

func MakeReversionReplicasByLoadDecider(pa *autoscalingv1alpha1.PodAutoscaler, config *autoscalerconfig.Config) *scaling.ReversionReplicasByLoad {
	panicThresholdPercentage := config.PanicThresholdPercentage // 200
	if x, ok := pa.PanicThresholdPercentage(); ok {
		panicThresholdPercentage = x
	}

	target, total := over_resources.ResolveMetricTarget(pa, config) // 70,100  ✅
	panicThreshold := panicThresholdPercentage / 100.0

	tbc := config.TargetBurstCapacity // 211
	if x, ok := pa.TargetBC(); ok {
		tbc = x
	}

	scaleDownDelay := config.ScaleDownDelay // 0
	if sdd, ok := pa.ScaleDownDelay(); ok {
		scaleDownDelay = sdd
	}

	var activationScale int32
	if mnzr, ok := pa.ActivationScale(); ok {
		activationScale = mnzr
	}

	return &scaling.ReversionReplicasByLoad{
		ObjectMeta: *pa.ObjectMeta.DeepCopy(),
		Spec: scaling.DeciderSpec{
			MaxScaleUpRate:      config.MaxScaleUpRate,
			MaxScaleDownRate:    config.MaxScaleDownRate,
			ScalingMetric:       pa.Metric(),
			TargetValue:         target,
			TotalValue:          total,
			TargetBurstCapacity: tbc,
			ActivatorCapacity:   config.ActivatorCapacity,
			PanicThreshold:      panicThreshold,
			StableWindow:        over_resources.StableWindow(pa, config),
			ScaleDownDelay:      scaleDownDelay,
			InitialScale:        GetInitialScale(config, pa),
			Reachable:           pa.Spec.Reachability != autoscalingv1alpha1.ReachabilityUnreachable,
			ActivationScale:     activationScale,
		},
	}
}
