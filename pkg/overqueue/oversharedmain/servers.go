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

package oversharedmain

import (
	"net/http"
	"strconv"
	"time"

	"knative.dev/serving/pkg/overnetworking"
	pkgnet "knative.dev/serving/pkg/overnetwork"
	"knative.dev/serving/pkg/overqueue"
)

func mainServer(addr string, handler http.Handler) *http.Server {
	return pkgnet.NewServer(addr, handler)
}

func adminServer(addr string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: handler,
		// https://medium.com/a-journey-with-go/go-understand-and-mitigate-slowloris-attack-711c1b1403f6
		ReadHeaderTimeout: time.Minute,
	}
}

func metricsServer(reporter *overqueue.ProtobufStatsReporter) *http.Server {
	metricsMux := http.NewServeMux()
	metricsMux.Handle("/metrics", overqueue.NewStatsHandler(reporter))

	return &http.Server{
		Addr:              ":" + strconv.Itoa(overnetworking.AutoscalingQueueMetricsPort),
		Handler:           metricsMux,
		ReadHeaderTimeout: time.Minute, // https://medium.com/a-journey-with-go/go-understand-and-mitigate-slowloris-attack-711c1b1403f6
	}
}
