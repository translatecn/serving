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
	"fmt"
	"os"

	// The set of controllers this controller process runs.
	"flag"
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	netcfg "knative.dev/serving/networking/pkg/config"
	"knative.dev/serving/pkg/injection"
	"knative.dev/serving/pkg/injection/sharedmain"
	"knative.dev/serving/pkg/networking"
	"knative.dev/serving/pkg/over_signals"
	"knative.dev/serving/pkg/over_system"
	"knative.dev/serving/pkg/reconciler"
	"knative.dev/serving/pkg/reconciler/configuration"
	"knative.dev/serving/pkg/reconciler/domainmapping"
	"knative.dev/serving/pkg/reconciler/labeler"
	"knative.dev/serving/pkg/reconciler/over_certificate"
	"knative.dev/serving/pkg/reconciler/over_gc"
	"knative.dev/serving/pkg/reconciler/over_nscert"
	"knative.dev/serving/pkg/reconciler/over_revision"
	"knative.dev/serving/pkg/reconciler/over_serverlessservice"
	"knative.dev/serving/pkg/reconciler/over_service"
	"knative.dev/serving/pkg/reconciler/route"

	versioned "github.com/cert-manager/cert-manager/pkg/client/clientset/versioned"
	"knative.dev/serving/pkg/client/certmanager/injection/informers/acme/v1/challenge"
	v1certificate "knative.dev/serving/pkg/client/certmanager/injection/informers/certmanager/v1/certificate"
	"knative.dev/serving/pkg/client/certmanager/injection/informers/certmanager/v1/certificaterequest"
	"knative.dev/serving/pkg/client/certmanager/injection/informers/certmanager/v1/clusterissuer"
	"knative.dev/serving/pkg/client/certmanager/injection/informers/certmanager/v1/issuer"
)

var ctors = []injection.ControllerConstructor{
	configuration.NewController,
	labeler.NewController,
	over_revision.NewController,
	route.NewController,
	over_serverlessservice.NewController,
	over_service.NewController,
	over_gc.NewController,
	over_nscert.NewController,
	domainmapping.NewController,
}

func init() {
	os.Setenv("CONFIG_LOGGING_NAME", "logging")
	os.Setenv("CONFIG_OBSERVABILITY_NAME", "config-observability")
	os.Setenv("POD_NAME", "ace")
	os.Setenv("SYSTEM_NAMESPACE", "knative-serving")
	os.Setenv("KUBECONFIG", "/Users/acejilam/.kube/koord")
}

func main() {
	flag.DurationVar(&reconciler.DefaultTimeout,
		"reconciliation-timeout", reconciler.DefaultTimeout,
		"The amount of time to give each reconciliation of a resource to complete before its context is canceled.")

	ctx := over_signals.NewContext()

	// HACK: This parses flags, so the above should be set once this runs.
	cfg := injection.ParseAndGetRESTConfigOrDie()

	// If nil it panics
	client := kubernetes.NewForConfigOrDie(cfg)

	if shouldEnableNetCertManagerController(ctx, client) {
		v := versioned.NewForConfigOrDie(cfg)
		if ok, err := certManagerCRDsExist(v); !ok {
			log.Fatalf("Please install cert-manager: %v", err)
		}
		for _, inf := range []injection.InformerInjector{
			challenge.WithInformer,
			v1certificate.WithInformer,
			certificaterequest.WithInformer,
			clusterissuer.WithInformer,
			issuer.WithInformer} {
			injection.Default.RegisterInformer(inf)
		}
		ctors = append(ctors, over_certificate.NewController)
	}

	sharedmain.MainWithConfig(ctx, "controller", cfg, ctors...)
}

func shouldEnableNetCertManagerController(ctx context.Context, client *kubernetes.Clientset) bool {
	var cm *v1.ConfigMap
	var err error
	if cm, err = client.CoreV1().ConfigMaps(over_system.Namespace()).Get(ctx, "config-network", metav1.GetOptions{}); err != nil {
		log.Fatalf("Failed to get cm config-network: %v", err)
	}
	netCfg, err := netcfg.NewConfigFromMap(cm.Data)
	if err != nil {
		log.Fatalf("Failed to construct network config: %v", err)
	}

	return networking.IsNetCertManagerControllerRequired(netCfg)
}

func certManagerCRDsExist(client *versioned.Clientset) (bool, error) {
	if ok, err := findCRD(client, "cert-manager.io/v1", []string{"certificaterequests", "certificates", "clusterissuers", "issuers"}); !ok {
		return false, err
	}
	if ok, err := findCRD(client, "acme.cert-manager.io/v1", []string{"challenges"}); !ok {
		return false, err
	}
	return true, nil
}

func findCRD(client *versioned.Clientset, groupVersion string, crds []string) (bool, error) {
	resourceList, err := client.Discovery().ServerResourcesForGroupVersion(groupVersion)
	if err != nil {
		return false, err
	}
	for _, crdName := range crds {
		isCRDPresent := false
		for _, resource := range resourceList.APIResources {
			if resource.Name == crdName {
				isCRDPresent = true
			}
		}
		if !isCRDPresent {
			return false, fmt.Errorf("cert manager crds are missing: %s", crdName)
		}
	}
	return true, nil
}
