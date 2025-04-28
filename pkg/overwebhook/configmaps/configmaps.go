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
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"go.uber.org/zap"
	admissionv1 "k8s.io/api/admission/v1"
	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	admissionlisters "k8s.io/client-go/listers/admissionregistration/v1"
	corelisters "k8s.io/client-go/listers/core/v1"

	"knative.dev/serving/pkg/overcontroller"
	"knative.dev/serving/pkg/overconfigmap"
	"knative.dev/serving/pkg/overkmp"
	"knative.dev/serving/pkg/overlogging"
	"knative.dev/serving/pkg/overptr"
	"knative.dev/serving/pkg/oversystem"
	"knative.dev/serving/pkg/overwebhook"
	certresources "knative.dev/serving/pkg/overwebhook/overcertificates/resources"
	pkgreconciler "knative.dev/serving/pkg/reconciler"
)

// reconciler implements the AdmissionController for ConfigMaps
type reconciler struct {
	overwebhook.StatelessAdmissionImpl
	pkgreconciler.LeaderAwareFuncs

	key          types.NamespacedName
	path         string
	constructors map[string]reflect.Value

	client       kubernetes.Interface
	vwhlister    admissionlisters.ValidatingWebhookConfigurationLister
	secretlister corelisters.SecretLister

	secretName                string
	disableNamespaceOwnership bool
}

var (
	_ overcontroller.Reconciler                = (*reconciler)(nil)
	_ pkgreconciler.LeaderAware                = (*reconciler)(nil)
	_ overwebhook.AdmissionController          = (*reconciler)(nil)
	_ overwebhook.StatelessAdmissionController = (*reconciler)(nil)
)

// Reconcile implements overcontroller.Reconciler
func (ac *reconciler) Reconcile(ctx context.Context, key string) error {
	logger := overlogging.FromContext(ctx)

	if !ac.IsLeaderFor(ac.key) {
		return overcontroller.NewSkipKey(key)
	}

	secret, err := ac.secretlister.Secrets(oversystem.Namespace()).Get(ac.secretName)
	if err != nil {
		logger.Errorw("Error fetching secret ", zap.Error(err))
		return err
	}

	caCert, ok := secret.Data[certresources.CACert]
	if !ok {
		return fmt.Errorf("secret %q is missing %q key", ac.secretName, certresources.CACert)
	}

	return ac.reconcileValidatingWebhook(ctx, caCert)
}

// Path implements AdmissionController
func (ac *reconciler) Path() string {
	return ac.path
}

// Admit implements AdmissionController
func (ac *reconciler) Admit(ctx context.Context, request *admissionv1.AdmissionRequest) *admissionv1.AdmissionResponse {
	logger := overlogging.FromContext(ctx)
	switch request.Operation {
	case admissionv1.Create, admissionv1.Update:
	default:
		logger.Info("Unhandled webhook operation, letting it through ", request.Operation)
		return &admissionv1.AdmissionResponse{Allowed: true}
	}

	if err := ac.validate(ctx, request); err != nil {
		return overwebhook.MakeErrorStatus("validation failed: %v", err)
	}

	return &admissionv1.AdmissionResponse{
		Allowed: true,
	}
}

func (ac *reconciler) reconcileValidatingWebhook(ctx context.Context, caCert []byte) error {
	logger := overlogging.FromContext(ctx)

	ruleScope := admissionregistrationv1.NamespacedScope
	rules := []admissionregistrationv1.RuleWithOperations{{
		Operations: []admissionregistrationv1.OperationType{
			admissionregistrationv1.Create,
			admissionregistrationv1.Update,
		},
		Rule: admissionregistrationv1.Rule{
			APIGroups:   []string{""},
			APIVersions: []string{"v1"},
			Resources:   []string{"configmaps/*"},
			Scope:       &ruleScope,
		},
	}}

	configuredWebhook, err := ac.vwhlister.Get(ac.key.Name)
	if err != nil {
		return fmt.Errorf("error retrieving webhook: %w", err)
	}

	webhook := configuredWebhook.DeepCopy()

	if !ac.disableNamespaceOwnership {
		// Set the owner to namespace.
		ns, err := ac.client.CoreV1().Namespaces().Get(ctx, oversystem.Namespace(), metav1.GetOptions{})
		if err != nil {
			return fmt.Errorf("failed to fetch namespace: %w", err)
		}
		nsRef := *metav1.NewControllerRef(ns, corev1.SchemeGroupVersion.WithKind("Namespace"))
		webhook.OwnerReferences = []metav1.OwnerReference{nsRef}
	}

	for i, wh := range webhook.Webhooks {
		if wh.Name != webhook.Name {
			continue
		}
		webhook.Webhooks[i].Rules = rules
		webhook.Webhooks[i].ClientConfig.CABundle = caCert
		if webhook.Webhooks[i].ClientConfig.Service == nil {
			return errors.New("missing service reference for webhook: " + wh.Name)
		}
		webhook.Webhooks[i].ClientConfig.Service.Path = overptr.String(ac.Path())
	}

	if ok, err := overkmp.SafeEqual(configuredWebhook, webhook); err != nil {
		return fmt.Errorf("error diffing webhooks: %w", err)
	} else if !ok {
		logger.Info("Updating webhook")
		vwhclient := ac.client.AdmissionregistrationV1().ValidatingWebhookConfigurations()
		if _, err := vwhclient.Update(ctx, webhook, metav1.UpdateOptions{}); err != nil {
			return fmt.Errorf("failed to update webhook: %w", err)
		}
	} else {
		logger.Info("Webhook is valid")
	}

	return nil
}

func (ac *reconciler) validate(ctx context.Context, req *admissionv1.AdmissionRequest) error {
	logger := overlogging.FromContext(ctx)
	kind := req.Kind
	newBytes := req.Object.Raw

	// Why, oh why are these different types...
	gvk := schema.GroupVersionKind{
		Group:   kind.Group,
		Version: kind.Version,
		Kind:    kind.Kind,
	}

	resourceGVK := corev1.SchemeGroupVersion.WithKind("ConfigMap")
	if gvk != resourceGVK {
		logger.Error("Unhandled kind: ", gvk)
		return fmt.Errorf("unhandled kind: %v", gvk)
	}

	var newObj corev1.ConfigMap
	if len(newBytes) != 0 {
		if err := json.Unmarshal(newBytes, &newObj); err != nil {
			return fmt.Errorf("cannot decode incoming new object: %w", err)
		}
	}

	if constructor, ok := ac.constructors[newObj.Name]; ok {
		// Only validate example data if this is a configMap we know about.
		exampleData, hasExampleData := newObj.Data[overconfigmap.ExampleKey]
		exampleChecksum, hasExampleChecksumAnnotation := newObj.Annotations[overconfigmap.ExampleChecksumAnnotation]
		if hasExampleData && hasExampleChecksumAnnotation &&
			exampleChecksum != overconfigmap.Checksum(exampleData) {
			return fmt.Errorf(
				"the update modifies a key in %q which is probably not what you want. Instead, copy the respective setting to the top-level of the ConfigMap, directly below %q",
				overconfigmap.ExampleKey, "data")
		}

		inputs := []reflect.Value{
			reflect.ValueOf(&newObj),
		}

		outputs := constructor.Call(inputs)
		errVal := outputs[1]

		if !errVal.IsNil() {
			return errVal.Interface().(error)
		}
	}

	return nil
}

func (ac *reconciler) registerConfig(name string, constructor interface{}) {
	if err := overconfigmap.ValidateConstructor(constructor); err != nil {
		panic(err)
	}

	ac.constructors[name] = reflect.ValueOf(constructor)
}
