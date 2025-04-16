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

package main

import (
	"context"
	"k8s.io/apimachinery/pkg/runtime/schema"
	servingv1beta1 "knative.dev/serving/pkg/apis/serving/v1beta1"
	"knative.dev/serving/pkg/configmap"
	"knative.dev/serving/pkg/controller"
	"knative.dev/serving/pkg/injection/sharedmain"
	"knative.dev/serving/pkg/leaderelection"
	"knative.dev/serving/pkg/metrics"
	"knative.dev/serving/pkg/over_logging"
	certconfig "knative.dev/serving/pkg/reconciler/certificate/config"
	"knative.dev/serving/pkg/signals"
	"knative.dev/serving/pkg/webhook"
	"knative.dev/serving/pkg/webhook/certificates"
	"knative.dev/serving/pkg/webhook/configmaps"
	"knative.dev/serving/pkg/webhook/resourcesemantics"
	"knative.dev/serving/pkg/webhook/resourcesemantics/defaulting"
	"knative.dev/serving/pkg/webhook/resourcesemantics/validation"

	// resource validation types
	net "knative.dev/serving/networking/pkg/apis/networking/v1alpha1"
	autoscalingv1alpha1 "knative.dev/serving/pkg/apis/autoscaling/v1alpha1"
	servingv1 "knative.dev/serving/pkg/apis/serving/v1"
	extravalidation "knative.dev/serving/pkg/webhook"

	// config validation constructors
	network "knative.dev/serving/networking/pkg"
	netcfg "knative.dev/serving/networking/pkg/config"
	apisconfig "knative.dev/serving/pkg/apis/config"
	autoscalerconfig "knative.dev/serving/pkg/autoscaler/config"
	"knative.dev/serving/pkg/deployment"
	"knative.dev/serving/pkg/gc"
	domainconfig "knative.dev/serving/pkg/reconciler/route/config"
	tracingconfig "knative.dev/serving/pkg/tracing/config"
)

var types = map[schema.GroupVersionKind]resourcesemantics.GenericCRD{
	servingv1.SchemeGroupVersion.WithKind("Revision"):      &servingv1.Revision{},
	servingv1.SchemeGroupVersion.WithKind("Configuration"): &servingv1.Configuration{},
	servingv1.SchemeGroupVersion.WithKind("Route"):         &servingv1.Route{},
	servingv1.SchemeGroupVersion.WithKind("Service"):       &servingv1.Service{},

	autoscalingv1alpha1.SchemeGroupVersion.WithKind("PodAutoscaler"): &autoscalingv1alpha1.PodAutoscaler{},
	autoscalingv1alpha1.SchemeGroupVersion.WithKind("Metric"):        &autoscalingv1alpha1.Metric{},

	net.SchemeGroupVersion.WithKind("Certificate"):       &net.Certificate{},
	net.SchemeGroupVersion.WithKind("Ingress"):           &net.Ingress{},
	net.SchemeGroupVersion.WithKind("ServerlessService"): &net.ServerlessService{},

	servingv1beta1.SchemeGroupVersion.WithKind("DomainMapping"): &servingv1beta1.DomainMapping{},
}

var serviceValidation = validation.NewCallback(
	extravalidation.ValidateService, webhook.Create, webhook.Update)

var configValidation = validation.NewCallback(
	extravalidation.ValidateConfiguration, webhook.Create, webhook.Update)

var callbacks = map[schema.GroupVersionKind]validation.Callback{
	servingv1.SchemeGroupVersion.WithKind("Service"):       serviceValidation,
	servingv1.SchemeGroupVersion.WithKind("Configuration"): configValidation,
}

func newDefaultingAdmissionController(ctx context.Context, cmw configmap.Watcher) *controller.Impl {
	// Decorate contexts with the current state of the config.
	store := apisconfig.NewStore(over_logging.FromContext(ctx).Named("config-store"))
	store.WatchConfigs(cmw)

	return defaulting.NewAdmissionController(ctx,

		// Name of the resource webhook.
		"webhook.serving.knative.dev",

		// The path on which to serve the webhook.
		"/defaulting",

		// The resources to default.
		types,

		// A function that infuses the context passed to Validate/SetDefaults with custom metadata.
		store.ToContext,

		// Whether to disallow unknown fields. We set this to 'false' since
		// our CRDs have schemas
		false,
	)
}

func newValidationAdmissionController(ctx context.Context, cmw configmap.Watcher) *controller.Impl {
	// Decorate contexts with the current state of the config.
	store := apisconfig.NewStore(over_logging.FromContext(ctx).Named("config-store"))
	store.WatchConfigs(cmw)

	return validation.NewAdmissionController(ctx,

		// Name of the resource webhook.
		"validation.webhook.serving.knative.dev",

		// The path on which to serve the webhook.
		"/resource-validation",

		// The resources to validate.
		types,

		// A function that infuses the context passed to Validate/SetDefaults with custom metadata.
		store.ToContext,

		// Whether to disallow unknown fields. We set this to 'false' since
		// our CRDs have schemas
		false,

		// Extra validating callbacks to be applied to resources.
		callbacks,
	)
}

func newConfigValidationController(ctx context.Context, cmw configmap.Watcher) *controller.Impl {
	return configmaps.NewAdmissionController(ctx,

		// Name of the configmap webhook.
		"config.webhook.serving.knative.dev",

		// The path on which to serve the webhook.
		"/config-validation",

		// The configmaps to validate.
		configmap.Constructors{
			tracingconfig.ConfigName:         tracingconfig.NewTracingConfigFromConfigMap,
			autoscalerconfig.ConfigName:      autoscalerconfig.NewConfigFromConfigMap,
			gc.ConfigName:                    gc.NewConfigFromConfigMapFunc(ctx),
			netcfg.ConfigMapName:             network.NewConfigFromConfigMap,
			deployment.ConfigName:            deployment.NewConfigFromConfigMap,
			apisconfig.FeaturesConfigName:    apisconfig.NewFeaturesConfigFromConfigMap,
			metrics.ConfigMapName():          metrics.NewObservabilityConfigFromConfigMap,
			over_logging.ConfigMapName():     over_logging.NewConfigFromConfigMap,
			leaderelection.ConfigMapName():   leaderelection.NewConfigFromConfigMap,
			domainconfig.DomainConfigName:    domainconfig.NewDomainFromConfigMap,
			apisconfig.DefaultsConfigName:    apisconfig.NewDefaultsConfigFromConfigMap,
			certconfig.CertManagerConfigName: certconfig.NewCertManagerConfigFromConfigMap,
		},
	)
}

func main() {
	// Set up a signal context with our webhook options
	ctx := webhook.WithOptions(signals.NewContext(), webhook.Options{
		ServiceName: webhook.NameFromEnv(),
		Port:        webhook.PortFromEnv(8443),
		SecretName:  "webhook-certs",
	})

	ctx = sharedmain.WithHealthProbesDisabled(ctx)
	sharedmain.WebhookMainWithContext(ctx, "webhook",
		certificates.NewController,
		newDefaultingAdmissionController,
		newValidationAdmissionController,
		newConfigValidationController,
	)
}
