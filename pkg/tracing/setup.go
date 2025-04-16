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

package tracing

import (
	"context"
)

type Tracer interface {
	// Shutdown allows for proper cleanup/flush.
	Shutdown(ctx context.Context) error
}

// SetupPublishingWithStaticConfig sets up trace publishing for the process. The caller should call .Shutdown() on
// the returned Tracer before shutdown to make sure all traces are properly flushed.
// Note that other pieces still need to generate the traces, this just ensures that if generated, they are collected
// appropriately. This is normally done by using tracing.HTTPSpanMiddleware as a middleware HTTP
// handler. The configuration will not be dynamically updated.

// Deprecated: Use SetupPublishingWithStaticConfig.

// SetupPublishingWithDynamicConfig sets up trace publishing for the process, by watching a
// ConfigMap for the configuration. The caller should call .Shutdown() on
// the returned OpenCensusTracer before shutdown to make sure all traces are properly flushed.
// Note that other pieces still need to generate the traces, this
// just ensures that if generated, they are collected appropriately. This is normally done by using
// tracing.HTTPSpanMiddleware as a middleware HTTP handler. The configuration will be dynamically
// updated when the ConfigMap is updated.

// Deprecated: Use SetupPublishingWithDynamicConfig.

// SetupPublishingWithDynamicConfigAndInitialValue sets up the trace publishing for the process with an
// initial value, by watching a ConfigMap for the configuration. Note that other pieces still
// need to generate the traces, this just ensures that if generated, they are collected
// appropriately. This is normally done by using tracing.HTTPSpanMiddleware as a middleware
// HTTP handler. The configuration will be dynamically updated when the ConfigMap is updated.
