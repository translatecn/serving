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

	// The set of controllers this controller process runs.
	"flag"
	"log"

	versioned "github.com/cert-manager/cert-manager/pkg/client/clientset/versioned"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	netcfg "knative.dev/serving/networking/pkg/config"
	"knative.dev/serving/pkg/client/certmanager/injection/informers/acme/v1/challenge"
	v1certificate "knative.dev/serving/pkg/client/certmanager/injection/informers/certmanager/v1/certificate"
	"knative.dev/serving/pkg/client/certmanager/injection/informers/certmanager/v1/certificaterequest"
	"knative.dev/serving/pkg/client/certmanager/injection/informers/certmanager/v1/clusterissuer"
	"knative.dev/serving/pkg/client/certmanager/injection/informers/certmanager/v1/issuer"
	"knative.dev/serving/pkg/injection"
	"knative.dev/serving/pkg/injection/sharedmain"
	"knative.dev/serving/pkg/overnetworking"
	"knative.dev/serving/pkg/oversignals"
	"knative.dev/serving/pkg/oversystem"
	"knative.dev/serving/pkg/reconciler"
	"knative.dev/serving/pkg/reconciler/overcertificate"
	"knative.dev/serving/pkg/reconciler/overconfiguration"
	"knative.dev/serving/pkg/reconciler/overdomainmapping"
	"knative.dev/serving/pkg/reconciler/overgc"
	"knative.dev/serving/pkg/reconciler/overlabeler"
	"knative.dev/serving/pkg/reconciler/overnscert"
	"knative.dev/serving/pkg/reconciler/overrevision"
	"knative.dev/serving/pkg/reconciler/overroute"
	"knative.dev/serving/pkg/reconciler/overserverlessservice"
	"knative.dev/serving/pkg/reconciler/overservice"
)

var ctors = []injection.ControllerConstructor{
	overconfiguration.NewController,
	overlabeler.NewController,
	overrevision.NewController,
	overroute.NewController,
	overserverlessservice.NewController,
	overservice.NewController,
	overgc.NewController,
	overnscert.NewController,
	overdomainmapping.NewController,
}

func main() {
	flag.DurationVar(&reconciler.DefaultTimeout,
		"reconciliation-timeout", reconciler.DefaultTimeout,
		"The amount of time to give each reconciliation of a resource to complete before its context is canceled.")

	ctx := oversignals.NewContext()

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
		ctors = append(ctors, overcertificate.NewController)
	}

	sharedmain.MainWithConfig(ctx, "controller", cfg, ctors...)
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

func shouldEnableNetCertManagerController(ctx context.Context, client *kubernetes.Clientset) bool {
	var cm *v1.ConfigMap
	var err error
	if cm, err = client.CoreV1().ConfigMaps(oversystem.Namespace()).Get(ctx, "config-network", metav1.GetOptions{}); err != nil {
		log.Fatalf("Failed to get cm config-network: %v", err)
	}
	netCfg, err := netcfg.NewConfigFromMap(cm.Data)
	if err != nil {
		log.Fatalf("Failed to construct network config: %v", err)
	}

	return overnetworking.IsNetCertManagerControllerRequired(netCfg)
}
