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

package configmaps

import (
	"context"

	"reflect"

	// Injection stuff
	kubeclient "knative.dev/serving/pkg/client/injection/kube/client"
	vwhinformer "knative.dev/serving/pkg/client/injection/kube/informers/admissionregistration/v1/validatingwebhookconfiguration"
	secretinformer "knative.dev/serving/pkg/injection/clients/namespacedkube/informers/core/v1/secret"
	"knative.dev/serving/pkg/over_logging"
	pkgreconciler "knative.dev/serving/pkg/reconciler"

	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/cache"
	"knative.dev/serving/pkg/configmap"
	"knative.dev/serving/pkg/controller"
	"knative.dev/serving/pkg/over_system"
	"knative.dev/serving/pkg/over_webhook"
)

// NewAdmissionController constructs a reconciler
func NewAdmissionController(
	ctx context.Context,
	name, path string,
	constructors configmap.Constructors,
) *controller.Impl {
	client := kubeclient.Get(ctx)
	vwhInformer := vwhinformer.Get(ctx)
	secretInformer := secretinformer.Get(ctx)
	options := over_webhook.GetOptions(ctx)

	// if this environment variable is set, it overrides the value in the Options
	disableNamespaceOwnership := over_webhook.DisableNamespaceOwnershipFromEnv()
	if disableNamespaceOwnership != nil {
		options.DisableNamespaceOwnership = *disableNamespaceOwnership
	}

	key := types.NamespacedName{Name: name}

	wh := &reconciler{
		LeaderAwareFuncs: pkgreconciler.LeaderAwareFuncs{
			// Have this reconciler enqueue our singleton whenever it becomes leader.
			PromoteFunc: func(bkt pkgreconciler.Bucket, enq func(pkgreconciler.Bucket, types.NamespacedName)) error {
				enq(bkt, key)
				return nil
			},
		},

		key:  key,
		path: path,

		constructors:              make(map[string]reflect.Value),
		secretName:                options.SecretName,
		disableNamespaceOwnership: options.DisableNamespaceOwnership,

		client:       client,
		vwhlister:    vwhInformer.Lister(),
		secretlister: secretInformer.Lister(),
	}

	for configName, constructor := range constructors {
		wh.registerConfig(configName, constructor)
	}

	const queueName = "ConfigMapWebhook"
	c := controller.NewContext(ctx, wh, controller.ControllerOptions{WorkQueueName: queueName, Logger: over_logging.FromContext(ctx).Named(queueName)})

	// Reconcile when the named ValidatingWebhookConfiguration changes.
	vwhInformer.Informer().AddEventHandler(cache.FilteringResourceEventHandler{
		FilterFunc: controller.FilterWithName(name),
		// It doesn't matter what we enqueue because we will always Reconcile
		// the named VWH resource.
		Handler: controller.HandleAll(c.Enqueue),
	})

	// Reconcile when the cert bundle changes.
	secretInformer.Informer().AddEventHandler(cache.FilteringResourceEventHandler{
		FilterFunc: controller.FilterWithNameAndNamespace(over_system.Namespace(), wh.secretName),
		// It doesn't matter what we enqueue because we will always Reconcile
		// the named VWH resource.
		Handler: controller.HandleAll(c.Enqueue),
	})

	return c
}
