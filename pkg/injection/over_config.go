/*
Copyright 2020 The Knative Authors

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

package injection

import (
	"flag"
	"log"
	"net/http"
	"strings"

	"knative.dev/serving/pkg/overenvironment"

	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
)

// ParseAndGetRESTConfigOrDie parses the rest config flags and creates a client or
// dies by calling log.Fatalf.
func ParseAndGetRESTConfigOrDie() *rest.Config {
	env := new(overenvironment.ClientConfig)
	env.InitFlags(flag.CommandLine)
	klog.InitFlags(flag.CommandLine)
	flag.Parse()
	cfg, err := env.GetRESTConfig()
	if err != nil {
		log.Fatal("Error building kubeconfig: ", err)
	}
	cfg.WrapTransport = func(rt http.RoundTripper) http.RoundTripper {
		return &LoggingTransport{rt: rt}
	}
	return cfg
}

type LoggingTransport struct {
	rt http.RoundTripper
}

func (l *LoggingTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	if !strings.Contains(request.URL.String(), "/leases/") {
		klog.Infoln(request.URL.String(), request.Method)
	}
	return l.rt.RoundTrip(request)
}
