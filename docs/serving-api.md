<p>Packages:</p>
<ul>
<li>
<a href="#autoscaling.internal.knative.dev%2fv1alpha1">autoscaling.internal.knative.dev/v1alpha1</a>
</li>
<li>
<a href="#duck.knative.dev%2fv1">duck.knative.dev/v1</a>
</li>
<li>
<a href="#duck.knative.dev%2fv1beta1">duck.knative.dev/v1beta1</a>
</li>
<li>
<a href="#serving.knative.dev%2fv1">serving.knative.dev/v1</a>
</li>
<li>
<a href="#serving.knative.dev%2fv1alpha1">serving.knative.dev/v1alpha1</a>
</li>
<li>
<a href="#serving.knative.dev%2fv1beta1">serving.knative.dev/v1beta1</a>
</li>
</ul>
<h2 id="autoscaling.internal.knative.dev/v1alpha1">autoscaling.internal.knative.dev/v1alpha1</h2>
<div>
<p>Package v1alpha1 contains the Autoscaling v1alpha1 API types.</p>
</div>
Resource Types:
<ul><li>
<a href="#autoscaling.internal.knative.dev/v1alpha1.PodAutoscaler">PodAutoscaler</a>
</li></ul>
<h3 id="autoscaling.internal.knative.dev/v1alpha1.PodAutoscaler">PodAutoscaler
</h3>
<div>
<p>PodAutoscaler is a Knative abstraction that encapsulates the interface by which Knative
components instantiate autoscalers.  This definition is an abstraction that may be backed
by multiple definitions.  For more information, see the Knative Pluggability presentation:
<a href="https://docs.google.com/presentation/d/19vW9HFZ6Puxt31biNZF3uLRejDmu82rxJIk1cWmxF7w/edit">https://docs.google.com/presentation/d/19vW9HFZ6Puxt31biNZF3uLRejDmu82rxJIk1cWmxF7w/edit</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>
autoscaling.internal.knative.dev/v1alpha1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>PodAutoscaler</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
<em>(Optional)</em>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#autoscaling.internal.knative.dev/v1alpha1.PodAutoscalerSpec">
PodAutoscalerSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Spec holds the desired state of the PodAutoscaler (from the client).</p>
<br/>
<br/>
<table>
<tr>
<td>
<code>containerConcurrency</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>ContainerConcurrency specifies the maximum allowed
in-flight (concurrent) requests per container of the Revision.
Defaults to <code>0</code> which means unlimited concurrency.</p>
</td>
</tr>
<tr>
<td>
<code>scaleTargetRef</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#objectreference-v1-core">
Kubernetes core/v1.ObjectReference
</a>
</em>
</td>
<td>
<p>ScaleTargetRef defines the /scale-able resource that this PodAutoscaler
is responsible for quickly right-sizing.</p>
</td>
</tr>
<tr>
<td>
<code>reachability</code><br/>
<em>
<a href="#autoscaling.internal.knative.dev/v1alpha1.ReachabilityType">
ReachabilityType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Reachability specifies whether or not the <code>ScaleTargetRef</code> can be reached (ie. has a route).
Defaults to <code>ReachabilityUnknown</code></p>
</td>
</tr>
<tr>
<td>
<code>protocolType</code><br/>
<em>
knative.dev/serving/networking/pkg/apis/networking.ProtocolType
</em>
</td>
<td>
<p>The application-layer protocol. Matches <code>ProtocolType</code> inferred from the revision spec.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#autoscaling.internal.knative.dev/v1alpha1.PodAutoscalerStatus">
PodAutoscalerStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Status communicates the observed state of the PodAutoscaler (from the controller).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="autoscaling.internal.knative.dev/v1alpha1.Metric">Metric
</h3>
<div>
<p>Metric represents a resource to configure the metric collector with.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
<em>(Optional)</em>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#autoscaling.internal.knative.dev/v1alpha1.MetricSpec">
MetricSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Spec holds the desired state of the Metric (from the client).</p>
<br/>
<br/>
<table>
<tr>
<td>
<code>stableWindow</code><br/>
<em>
<a href="https://golang.org/pkg/time/#Duration">
time.Duration
</a>
</em>
</td>
<td>
<p>StableWindow is the aggregation window for metrics in a stable state.</p>
</td>
</tr>
<tr>
<td>
<code>panicWindow</code><br/>
<em>
<a href="https://golang.org/pkg/time/#Duration">
time.Duration
</a>
</em>
</td>
<td>
<p>PanicWindow is the aggregation window for metrics where quick reactions are needed.</p>
</td>
</tr>
<tr>
<td>
<code>scrapeTarget</code><br/>
<em>
string
</em>
</td>
<td>
<p>ScrapeTarget is the K8s service that publishes the metric endpoint.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#autoscaling.internal.knative.dev/v1alpha1.MetricStatus">
MetricStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Status communicates the observed state of the Metric (from the controller).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="autoscaling.internal.knative.dev/v1alpha1.MetricSpec">MetricSpec
</h3>
<p>
(<em>Appears on:</em><a href="#autoscaling.internal.knative.dev/v1alpha1.Metric">Metric</a>)
</p>
<div>
<p>MetricSpec contains all values a metric collector needs to operate.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>stableWindow</code><br/>
<em>
<a href="https://golang.org/pkg/time/#Duration">
time.Duration
</a>
</em>
</td>
<td>
<p>StableWindow is the aggregation window for metrics in a stable state.</p>
</td>
</tr>
<tr>
<td>
<code>panicWindow</code><br/>
<em>
<a href="https://golang.org/pkg/time/#Duration">
time.Duration
</a>
</em>
</td>
<td>
<p>PanicWindow is the aggregation window for metrics where quick reactions are needed.</p>
</td>
</tr>
<tr>
<td>
<code>scrapeTarget</code><br/>
<em>
string
</em>
</td>
<td>
<p>ScrapeTarget is the K8s service that publishes the metric endpoint.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="autoscaling.internal.knative.dev/v1alpha1.MetricStatus">MetricStatus
</h3>
<p>
(<em>Appears on:</em><a href="#autoscaling.internal.knative.dev/v1alpha1.Metric">Metric</a>)
</p>
<div>
<p>MetricStatus reflects the status of metric collection for this specific entity.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>Status</code><br/>
<em>
<a href="#duck.knative.dev/v1.Status">
Status
</a>
</em>
</td>
<td>
<p>
(Members of <code>Status</code> are embedded into this type.)
</p>
</td>
</tr>
</tbody>
</table>
<h3 id="autoscaling.internal.knative.dev/v1alpha1.PodAutoscalerSpec">PodAutoscalerSpec
</h3>
<p>
(<em>Appears on:</em><a href="#autoscaling.internal.knative.dev/v1alpha1.PodAutoscaler">PodAutoscaler</a>)
</p>
<div>
<p>PodAutoscalerSpec holds the desired state of the PodAutoscaler (from the client).</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>containerConcurrency</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>ContainerConcurrency specifies the maximum allowed
in-flight (concurrent) requests per container of the Revision.
Defaults to <code>0</code> which means unlimited concurrency.</p>
</td>
</tr>
<tr>
<td>
<code>scaleTargetRef</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#objectreference-v1-core">
Kubernetes core/v1.ObjectReference
</a>
</em>
</td>
<td>
<p>ScaleTargetRef defines the /scale-able resource that this PodAutoscaler
is responsible for quickly right-sizing.</p>
</td>
</tr>
<tr>
<td>
<code>reachability</code><br/>
<em>
<a href="#autoscaling.internal.knative.dev/v1alpha1.ReachabilityType">
ReachabilityType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Reachability specifies whether or not the <code>ScaleTargetRef</code> can be reached (ie. has a route).
Defaults to <code>ReachabilityUnknown</code></p>
</td>
</tr>
<tr>
<td>
<code>protocolType</code><br/>
<em>
knative.dev/serving/networking/pkg/apis/networking.ProtocolType
</em>
</td>
<td>
<p>The application-layer protocol. Matches <code>ProtocolType</code> inferred from the revision spec.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="autoscaling.internal.knative.dev/v1alpha1.PodAutoscalerStatus">PodAutoscalerStatus
</h3>
<p>
(<em>Appears on:</em><a href="#autoscaling.internal.knative.dev/v1alpha1.PodAutoscaler">PodAutoscaler</a>)
</p>
<div>
<p>PodAutoscalerStatus communicates the observed state of the PodAutoscaler (from the controller).</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>Status</code><br/>
<em>
<a href="#duck.knative.dev/v1.Status">
Status
</a>
</em>
</td>
<td>
<p>
(Members of <code>Status</code> are embedded into this type.)
</p>
</td>
</tr>
<tr>
<td>
<code>serviceName</code><br/>
<em>
string
</em>
</td>
<td>
<p>ServiceName is the K8s Service name that serves the revision, scaled by this PA.
The service is created and owned by the ServerlessService object owned by this PA.</p>
</td>
</tr>
<tr>
<td>
<code>metricsServiceName</code><br/>
<em>
string
</em>
</td>
<td>
<p>MetricsServiceName is the K8s Service name that provides revision metrics.
The service is managed by the PA object.</p>
</td>
</tr>
<tr>
<td>
<code>desiredScale</code><br/>
<em>
int32
</em>
</td>
<td>
<p>DesiredScale 显示了此次修订所需的当前期望副本数量。</p>
</td>
</tr>
<tr>
<td>
<code>actualScale</code><br/>
<em>
int32
</em>
</td>
<td>
<p>ActualScale 显示了该修订版本的实际副本数量。</p>
</td>
</tr>
</tbody>
</table>
<h3 id="autoscaling.internal.knative.dev/v1alpha1.PodScalable">PodScalable
</h3>
<div>
<p>PodScalable is a duck type that the resources referenced by the
PodAutoscaler&rsquo;s ScaleTargetRef must implement.  They must also
implement the <code>/scale</code> sub-resource for use with <code>/scale</code> based
implementations (e.g. HPA), but this further constrains the shape
the referenced resources may take.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#autoscaling.internal.knative.dev/v1alpha1.PodScalableSpec">
PodScalableSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>replicas</code><br/>
<em>
int32
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>selector</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#labelselector-v1-meta">
Kubernetes meta/v1.LabelSelector
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>template</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#podtemplatespec-v1-core">
Kubernetes core/v1.PodTemplateSpec
</a>
</em>
</td>
<td>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#autoscaling.internal.knative.dev/v1alpha1.PodScalableStatus">
PodScalableStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="autoscaling.internal.knative.dev/v1alpha1.PodScalableSpec">PodScalableSpec
</h3>
<p>
(<em>Appears on:</em><a href="#autoscaling.internal.knative.dev/v1alpha1.PodScalable">PodScalable</a>)
</p>
<div>
<p>PodScalableSpec is the specification for the desired state of a
PodScalable (or at least our shared portion).</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>replicas</code><br/>
<em>
int32
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>selector</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#labelselector-v1-meta">
Kubernetes meta/v1.LabelSelector
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>template</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#podtemplatespec-v1-core">
Kubernetes core/v1.PodTemplateSpec
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="autoscaling.internal.knative.dev/v1alpha1.PodScalableStatus">PodScalableStatus
</h3>
<p>
(<em>Appears on:</em><a href="#autoscaling.internal.knative.dev/v1alpha1.PodScalable">PodScalable</a>)
</p>
<div>
<p>PodScalableStatus is the observed state of a PodScalable (or at
least our shared portion).</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>replicas</code><br/>
<em>
int32
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="autoscaling.internal.knative.dev/v1alpha1.ReachabilityType">ReachabilityType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#autoscaling.internal.knative.dev/v1alpha1.PodAutoscalerSpec">PodAutoscalerSpec</a>)
</p>
<div>
<p>ReachabilityType is the enumeration type for the different states of reachability
to the <code>ScaleTarget</code> of a <code>PodAutoscaler</code></p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Reachable&#34;</p></td>
<td><p>ReachabilityReachable means the <code>ScaleTarget</code> is reachable, ie. it has an active route.</p>
</td>
</tr><tr><td><p>&#34;&#34;</p></td>
<td><p>ReachabilityUnknown means the reachability of the <code>ScaleTarget</code> is unknown.
Used when the reachability cannot be determined, eg. during activation.</p>
</td>
</tr><tr><td><p>&#34;Unreachable&#34;</p></td>
<td><p>ReachabilityUnreachable means the <code>ScaleTarget</code> is not reachable, ie. it does not have an active route.</p>
</td>
</tr></tbody>
</table>
<hr/>
<h2 id="duck.knative.dev/v1">duck.knative.dev/v1</h2>
<div>
</div>
Resource Types:
<ul></ul>
<h3 id="duck.knative.dev/v1.AddressStatus">AddressStatus
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1.AddressableType">AddressableType</a>)
</p>
<div>
<p>AddressStatus shows how we expect folks to embed Addressable in
their Status field.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>address</code><br/>
<em>
<a href="#duck.knative.dev/v1.Addressable">
Addressable
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Address is a single Addressable address.
If Addresses is present, Address will be ignored by clients.</p>
</td>
</tr>
<tr>
<td>
<code>addresses</code><br/>
<em>
<a href="#duck.knative.dev/v1.Addressable">
[]Addressable
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Addresses is a list of addresses for different protocols (HTTP and HTTPS)
If Addresses is present, Address must be ignored by clients.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1.Addressable">Addressable
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1.AddressStatus">AddressStatus</a>, <a href="#serving.knative.dev/v1.RouteStatusFields">RouteStatusFields</a>, <a href="#serving.knative.dev/v1beta1.DomainMappingStatus">DomainMappingStatus</a>)
</p>
<div>
<p>Addressable provides a generic mechanism for a custom resource
definition to indicate a destination for message delivery.</p>
<p>Addressable is the schema for the destination information. This is
typically stored in the object&rsquo;s <code>status</code>, as this information may
be generated by the controller.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Name is the name of the address.</p>
</td>
</tr>
<tr>
<td>
<code>url</code><br/>
<em>
knative.dev/serving/pkg/apis.URL
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>CACerts</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>CACerts is the Certification Authority (CA) certificates in PEM format
according to <a href="https://www.rfc-editor.org/rfc/rfc7468">https://www.rfc-editor.org/rfc/rfc7468</a>.</p>
</td>
</tr>
<tr>
<td>
<code>audience</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Audience is the OIDC audience for this address.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1.AddressableType">AddressableType
</h3>
<div>
<p>AddressableType is a skeleton type wrapping Addressable in the manner we expect
resource writers defining compatible resources to embed it.  We will
typically use this type to deserialize Addressable ObjectReferences and
access the Addressable data.  This is not a real resource.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#duck.knative.dev/v1.AddressStatus">
AddressStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1.AuthStatus">AuthStatus
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1.AuthenticatableStatus">AuthenticatableStatus</a>, <a href="#duck.knative.dev/v1.SourceStatus">SourceStatus</a>)
</p>
<div>
<p>AuthStatus is meant to provide the generated service account name
in the resource status.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>serviceAccountName</code><br/>
<em>
string
</em>
</td>
<td>
<p>ServiceAccountName is the name of the generated service account
used for this components OIDC authentication.</p>
</td>
</tr>
<tr>
<td>
<code>serviceAccountNames</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>ServiceAccountNames is the list of names of the generated service accounts
used for this components OIDC authentication. This list can have len() &gt; 1,
when the component uses multiple identities (e.g. in case of a Parallel).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1.AuthenticatableStatus">AuthenticatableStatus
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1.AuthenticatableType">AuthenticatableType</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>auth</code><br/>
<em>
<a href="#duck.knative.dev/v1.AuthStatus">
AuthStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Auth contains the service account name for the subscription</p>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1.AuthenticatableType">AuthenticatableType
</h3>
<div>
<p>AuthenticatableType is a skeleton type wrapping AuthStatus in the manner we expect
resource writers defining compatible resources to embed it.  We will
typically use this type to deserialize AuthenticatableType ObjectReferences and
access the AuthenticatableType data.  This is not a real resource.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#duck.knative.dev/v1.AuthenticatableStatus">
AuthenticatableStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1.Binding">Binding
</h3>
<div>
<p>Binding is a duck type that specifies the partial schema to which all
Binding implementations should adhere.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#duck.knative.dev/v1.BindingSpec">
BindingSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>subject</code><br/>
<em>
knative.dev/serving/pkg/tracker.Reference
</em>
</td>
<td>
<p>Subject references the resource(s) whose &ldquo;runtime contract&rdquo; should be
augmented by Binding implementations.</p>
</td>
</tr>
</table>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1.BindingSpec">BindingSpec
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1.Binding">Binding</a>)
</p>
<div>
<p>BindingSpec specifies the spec portion of the Binding partial-schema.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>subject</code><br/>
<em>
knative.dev/serving/pkg/tracker.Reference
</em>
</td>
<td>
<p>Subject references the resource(s) whose &ldquo;runtime contract&rdquo; should be
augmented by Binding implementations.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1.CloudEventAttributes">CloudEventAttributes
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1.SourceStatus">SourceStatus</a>)
</p>
<div>
<p>CloudEventAttributes specifies the attributes that a Source
uses as part of its CloudEvents.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type refers to the CloudEvent type attribute.</p>
</td>
</tr>
<tr>
<td>
<code>source</code><br/>
<em>
string
</em>
</td>
<td>
<p>Source is the CloudEvents source attribute.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1.CloudEventOverrides">CloudEventOverrides
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1.SourceSpec">SourceSpec</a>)
</p>
<div>
<p>CloudEventOverrides defines arguments for a Source that control the output
format of the CloudEvents produced by the Source.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>extensions</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Extensions specify what attribute are added or overridden on the
outbound event. Each <code>Extensions</code> key-value pair are set on the event as
an attribute extension independently.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1.Conditions">Conditions
(<code>[]knative.dev/serving/pkg/apis.Condition</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1.Status">Status</a>)
</p>
<div>
<p>Conditions is a simple wrapper around apis.Conditions to implement duck.Implementable.</p>
</div>
<h3 id="duck.knative.dev/v1.CronJob">CronJob
</h3>
<div>
<p>CronJob is a wrapper around CronJob resource, which supports our interfaces
for webhooks</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#cronjobspec-v1-batch">
Kubernetes batch/v1.CronJobSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>schedule</code><br/>
<em>
string
</em>
</td>
<td>
<p>The schedule in Cron format, see <a href="https://en.wikipedia.org/wiki/Cron">https://en.wikipedia.org/wiki/Cron</a>.</p>
</td>
</tr>
<tr>
<td>
<code>timeZone</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The time zone name for the given schedule, see <a href="https://en.wikipedia.org/wiki/List_of_tz_database_time_zones">https://en.wikipedia.org/wiki/List_of_tz_database_time_zones</a>.
If not specified, this will default to the time zone of the kube-controller-manager process.
The set of valid time zone names and the time zone offset is loaded from the system-wide time zone
database by the API server during CronJob validation and the controller manager during execution.
If no system-wide time zone database can be found a bundled version of the database is used instead.
If the time zone name becomes invalid during the lifetime of a CronJob or due to a change in host
configuration, the controller will stop creating new new Jobs and will create a system event with the
reason UnknownTimeZone.
More information can be found in <a href="https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/#time-zones">https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/#time-zones</a></p>
</td>
</tr>
<tr>
<td>
<code>startingDeadlineSeconds</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>Optional deadline in seconds for starting the job if it misses scheduled
time for any reason.  Missed jobs executions will be counted as failed ones.</p>
</td>
</tr>
<tr>
<td>
<code>concurrencyPolicy</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#concurrencypolicy-v1-batch">
Kubernetes batch/v1.ConcurrencyPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies how to treat concurrent executions of a Job.
Valid values are:</p>
<ul>
<li>&ldquo;Allow&rdquo; (default): allows CronJobs to run concurrently;</li>
<li>&ldquo;Forbid&rdquo;: forbids concurrent runs, skipping next run if previous run hasn&rsquo;t finished yet;</li>
<li>&ldquo;Replace&rdquo;: cancels currently running job and replaces it with a new one</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>suspend</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>This flag tells the controller to suspend subsequent executions, it does
not apply to already started executions.  Defaults to false.</p>
</td>
</tr>
<tr>
<td>
<code>jobTemplate</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#jobtemplatespec-v1-batch">
Kubernetes batch/v1.JobTemplateSpec
</a>
</em>
</td>
<td>
<p>Specifies the job that will be created when executing a CronJob.</p>
</td>
</tr>
<tr>
<td>
<code>successfulJobsHistoryLimit</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>The number of successful finished jobs to retain. Value must be non-negative integer.
Defaults to 3.</p>
</td>
</tr>
<tr>
<td>
<code>failedJobsHistoryLimit</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>The number of failed finished jobs to retain. Value must be non-negative integer.
Defaults to 1.</p>
</td>
</tr>
</table>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1.CronJobDefaulter">CronJobDefaulter
</h3>
<div>
<p>CronJobDefaulter is a callback to validate a CronJob.</p>
</div>
<h3 id="duck.knative.dev/v1.CronJobValidator">CronJobValidator
</h3>
<div>
<p>CronJobValidator is a callback to validate a CronJob.</p>
</div>
<h3 id="duck.knative.dev/v1.Destination">Destination
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1.SourceSpec">SourceSpec</a>)
</p>
<div>
<p>Destination represents a target of an invocation over HTTP.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ref</code><br/>
<em>
<a href="#duck.knative.dev/v1.KReference">
KReference
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Ref points to an Addressable.</p>
</td>
</tr>
<tr>
<td>
<code>uri</code><br/>
<em>
knative.dev/serving/pkg/apis.URL
</em>
</td>
<td>
<em>(Optional)</em>
<p>URI can be an absolute URL(non-empty scheme and non-empty host) pointing to the target or a relative URI. Relative URIs will be resolved using the base URI retrieved from Ref.</p>
</td>
</tr>
<tr>
<td>
<code>CACerts</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>CACerts are Certification Authority (CA) certificates in PEM format
according to <a href="https://www.rfc-editor.org/rfc/rfc7468">https://www.rfc-editor.org/rfc/rfc7468</a>.
If set, these CAs are appended to the set of CAs provided
by the Addressable target, if any.</p>
</td>
</tr>
<tr>
<td>
<code>audience</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Audience is the OIDC audience.
This need only be set, if the target is not an Addressable
and thus the Audience can&rsquo;t be received from the Addressable itself.
In case the Addressable specifies an Audience too, the Destinations
Audience takes preference.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1.KRShaped">KRShaped
</h3>
<div>
<p>KRShaped is an interface for retrieving the duck elements of an arbitrary resource.</p>
</div>
<h3 id="duck.knative.dev/v1.KReference">KReference
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1.Destination">Destination</a>, <a href="#serving.knative.dev/v1beta1.DomainMappingSpec">DomainMappingSpec</a>)
</p>
<div>
<p>KReference contains enough information to refer to another object.
It&rsquo;s a trimmed down version of corev1.ObjectReference.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>kind</code><br/>
<em>
string
</em>
</td>
<td>
<p>Kind of the referent.
More info: <a href="https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds">https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds</a></p>
</td>
</tr>
<tr>
<td>
<code>namespace</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Namespace of the referent.
More info: <a href="https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/">https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/</a>
This is optional field, it gets defaulted to the object holding it if left out.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name of the referent.
More info: <a href="https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names">https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names</a></p>
</td>
</tr>
<tr>
<td>
<code>apiVersion</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>API version of the referent.</p>
</td>
</tr>
<tr>
<td>
<code>group</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Group of the API, without the version of the group. This can be used as an alternative to the APIVersion, and then resolved using ResolveGroup.
Note: This API is EXPERIMENTAL and might break anytime. For more details: <a href="https://github.com/knative/eventing/issues/5086">https://github.com/knative/eventing/issues/5086</a></p>
</td>
</tr>
<tr>
<td>
<code>address</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Address points to a specific Address Name.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1.KResource">KResource
</h3>
<div>
<p>KResource is a skeleton type wrapping Conditions in the manner we expect
resource writers defining compatible resources to embed it.  We will
typically use this type to deserialize Conditions ObjectReferences and
access the Conditions data.  This is not a real resource.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#duck.knative.dev/v1.Status">
Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1.Pod">Pod
</h3>
<div>
<p>Pod is a wrapper around Pod-like resource, which supports our interfaces
for webhooks</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#podspec-v1-core">
Kubernetes core/v1.PodSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>volumes</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#volume-v1-core">
[]Kubernetes core/v1.Volume
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>List of volumes that can be mounted by containers belonging to the pod.
More info: <a href="https://kubernetes.io/docs/concepts/storage/volumes">https://kubernetes.io/docs/concepts/storage/volumes</a></p>
</td>
</tr>
<tr>
<td>
<code>initContainers</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#container-v1-core">
[]Kubernetes core/v1.Container
</a>
</em>
</td>
<td>
<p>List of initialization containers belonging to the pod.
Init containers are executed in order prior to containers being started. If any
init container fails, the pod is considered to have failed and is handled according
to its restartPolicy. The name for an init container or normal container must be
unique among all containers.
Init containers may not have Lifecycle actions, Readiness probes, Liveness probes, or Startup probes.
The resourceRequirements of an init container are taken into account during scheduling
by finding the highest request/limit for each resource type, and then using the max of
of that value or the sum of the normal containers. Limits are applied to init containers
in a similar fashion.
Init containers cannot currently be added or removed.
Cannot be updated.
More info: <a href="https://kubernetes.io/docs/concepts/workloads/pods/init-containers/">https://kubernetes.io/docs/concepts/workloads/pods/init-containers/</a></p>
</td>
</tr>
<tr>
<td>
<code>containers</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#container-v1-core">
[]Kubernetes core/v1.Container
</a>
</em>
</td>
<td>
<p>List of containers belonging to the pod.
Containers cannot currently be added or removed.
There must be at least one container in a Pod.
Cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>ephemeralContainers</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#ephemeralcontainer-v1-core">
[]Kubernetes core/v1.EphemeralContainer
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>List of ephemeral containers run in this pod. Ephemeral containers may be run in an existing
pod to perform user-initiated actions such as debugging. This list cannot be specified when
creating a pod, and it cannot be modified by updating the pod spec. In order to add an
ephemeral container to an existing pod, use the pod&rsquo;s ephemeralcontainers subresource.</p>
</td>
</tr>
<tr>
<td>
<code>restartPolicy</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#restartpolicy-v1-core">
Kubernetes core/v1.RestartPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Restart policy for all containers within the pod.
One of Always, OnFailure, Never. In some contexts, only a subset of those values may be permitted.
Default to Always.
More info: <a href="https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy">https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy</a></p>
</td>
</tr>
<tr>
<td>
<code>terminationGracePeriodSeconds</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete request.
Value must be non-negative integer. The value zero indicates stop immediately via
the kill signal (no opportunity to shut down).
If this value is nil, the default grace period will be used instead.
The grace period is the duration in seconds after the processes running in the pod are sent
a termination signal and the time when the processes are forcibly halted with a kill signal.
Set this value longer than the expected cleanup time for your process.
Defaults to 30 seconds.</p>
</td>
</tr>
<tr>
<td>
<code>activeDeadlineSeconds</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>Optional duration in seconds the pod may be active on the node relative to
StartTime before the system will actively try to mark it failed and kill associated containers.
Value must be a positive integer.</p>
</td>
</tr>
<tr>
<td>
<code>dnsPolicy</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#dnspolicy-v1-core">
Kubernetes core/v1.DNSPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Set DNS policy for the pod.
Defaults to &ldquo;ClusterFirst&rdquo;.
Valid values are &lsquo;ClusterFirstWithHostNet&rsquo;, &lsquo;ClusterFirst&rsquo;, &lsquo;Default&rsquo; or &lsquo;None&rsquo;.
DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy.
To have DNS options set along with hostNetwork, you have to specify DNS policy
explicitly to &lsquo;ClusterFirstWithHostNet&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>nodeSelector</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>NodeSelector is a selector which must be true for the pod to fit on a node.
Selector which must match a node&rsquo;s labels for the pod to be scheduled on that node.
More info: <a href="https://kubernetes.io/docs/concepts/configuration/assign-pod-node/">https://kubernetes.io/docs/concepts/configuration/assign-pod-node/</a></p>
</td>
</tr>
<tr>
<td>
<code>serviceAccountName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ServiceAccountName is the name of the ServiceAccount to use to run this pod.
More info: <a href="https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/">https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/</a></p>
</td>
</tr>
<tr>
<td>
<code>serviceAccount</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>DeprecatedServiceAccount is a deprecated alias for ServiceAccountName.
Deprecated: Use serviceAccountName instead.</p>
</td>
</tr>
<tr>
<td>
<code>automountServiceAccountToken</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>AutomountServiceAccountToken indicates whether a service account token should be automatically mounted.</p>
</td>
</tr>
<tr>
<td>
<code>nodeName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>NodeName indicates in which node this pod is scheduled.
If empty, this pod is a candidate for scheduling by the scheduler defined in schedulerName.
Once this field is set, the kubelet for this node becomes responsible for the lifecycle of this pod.
This field should not be used to express a desire for the pod to be scheduled on a specific node.
<a href="https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#nodename">https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#nodename</a></p>
</td>
</tr>
<tr>
<td>
<code>hostNetwork</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Host networking requested for this pod. Use the host&rsquo;s network namespace.
If this option is set, the ports that will be used must be specified.
Default to false.</p>
</td>
</tr>
<tr>
<td>
<code>hostPID</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Use the host&rsquo;s pid namespace.
Optional: Default to false.</p>
</td>
</tr>
<tr>
<td>
<code>hostIPC</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Use the host&rsquo;s ipc namespace.
Optional: Default to false.</p>
</td>
</tr>
<tr>
<td>
<code>shareProcessNamespace</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Share a single process namespace between all of the containers in a pod.
When this is set containers will be able to view and signal processes from other containers
in the same pod, and the first process in each container will not be assigned PID 1.
HostPID and ShareProcessNamespace cannot both be set.
Optional: Default to false.</p>
</td>
</tr>
<tr>
<td>
<code>securityContext</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#podsecuritycontext-v1-core">
Kubernetes core/v1.PodSecurityContext
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>SecurityContext holds pod-level security attributes and common container settings.
Optional: Defaults to empty.  See type description for default values of each field.</p>
</td>
</tr>
<tr>
<td>
<code>imagePullSecrets</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#localobjectreference-v1-core">
[]Kubernetes core/v1.LocalObjectReference
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.
If specified, these secrets will be passed to individual puller implementations for them to use.
More info: <a href="https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod">https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod</a></p>
</td>
</tr>
<tr>
<td>
<code>hostname</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the hostname of the Pod
If not specified, the pod&rsquo;s hostname will be set to a system-defined value.</p>
</td>
</tr>
<tr>
<td>
<code>subdomain</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>If specified, the fully qualified Pod hostname will be &ldquo;<hostname>.<subdomain>.<pod namespace>.svc.<cluster domain>&rdquo;.
If not specified, the pod will not have a domainname at all.</p>
</td>
</tr>
<tr>
<td>
<code>affinity</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#affinity-v1-core">
Kubernetes core/v1.Affinity
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>If specified, the pod&rsquo;s scheduling constraints</p>
</td>
</tr>
<tr>
<td>
<code>schedulerName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>If specified, the pod will be dispatched by specified scheduler.
If not specified, the pod will be dispatched by default scheduler.</p>
</td>
</tr>
<tr>
<td>
<code>tolerations</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#toleration-v1-core">
[]Kubernetes core/v1.Toleration
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>If specified, the pod&rsquo;s tolerations.</p>
</td>
</tr>
<tr>
<td>
<code>hostAliases</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#hostalias-v1-core">
[]Kubernetes core/v1.HostAlias
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>HostAliases is an optional list of hosts and IPs that will be injected into the pod&rsquo;s hosts
file if specified.</p>
</td>
</tr>
<tr>
<td>
<code>priorityClassName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>If specified, indicates the pod&rsquo;s priority. &ldquo;system-node-critical&rdquo; and
&ldquo;system-cluster-critical&rdquo; are two special keywords which indicate the
highest priorities with the former being the highest priority. Any other
name must be defined by creating a PriorityClass object with that name.
If not specified, the pod priority will be default or zero if there is no
default.</p>
</td>
</tr>
<tr>
<td>
<code>priority</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>The priority value. Various system components use this field to find the
priority of the pod. When Priority Admission Controller is enabled, it
prevents users from setting this field. The admission controller populates
this field from PriorityClassName.
The higher the value, the higher the priority.</p>
</td>
</tr>
<tr>
<td>
<code>dnsConfig</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#poddnsconfig-v1-core">
Kubernetes core/v1.PodDNSConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the DNS parameters of a pod.
Parameters specified here will be merged to the generated DNS
configuration based on DNSPolicy.</p>
</td>
</tr>
<tr>
<td>
<code>readinessGates</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#podreadinessgate-v1-core">
[]Kubernetes core/v1.PodReadinessGate
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>If specified, all readiness gates will be evaluated for pod readiness.
A pod is ready when all its containers are ready AND
all conditions specified in the readiness gates have status equal to &ldquo;True&rdquo;
More info: <a href="https://git.k8s.io/enhancements/keps/sig-network/580-pod-readiness-gates">https://git.k8s.io/enhancements/keps/sig-network/580-pod-readiness-gates</a></p>
</td>
</tr>
<tr>
<td>
<code>runtimeClassName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>RuntimeClassName refers to a RuntimeClass object in the node.k8s.io group, which should be used
to run this pod.  If no RuntimeClass resource matches the named class, the pod will not be run.
If unset or empty, the &ldquo;legacy&rdquo; RuntimeClass will be used, which is an implicit class with an
empty definition that uses the default runtime handler.
More info: <a href="https://git.k8s.io/enhancements/keps/sig-node/585-runtime-class">https://git.k8s.io/enhancements/keps/sig-node/585-runtime-class</a></p>
</td>
</tr>
<tr>
<td>
<code>enableServiceLinks</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>EnableServiceLinks indicates whether information about services should be injected into pod&rsquo;s
environment variables, matching the syntax of Docker links.
Optional: Defaults to true.</p>
</td>
</tr>
<tr>
<td>
<code>preemptionPolicy</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#preemptionpolicy-v1-core">
Kubernetes core/v1.PreemptionPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>PreemptionPolicy is the Policy for preempting pods with lower priority.
One of Never, PreemptLowerPriority.
Defaults to PreemptLowerPriority if unset.</p>
</td>
</tr>
<tr>
<td>
<code>overhead</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#resourcelist-v1-core">
Kubernetes core/v1.ResourceList
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Overhead represents the resource overhead associated with running a pod for a given RuntimeClass.
This field will be autopopulated at admission time by the RuntimeClass admission controller. If
the RuntimeClass admission controller is enabled, overhead must not be set in Pod create requests.
The RuntimeClass admission controller will reject Pod create requests which have the overhead already
set. If RuntimeClass is configured and selected in the PodSpec, Overhead will be set to the value
defined in the corresponding RuntimeClass, otherwise it will remain unset and treated as zero.
More info: <a href="https://git.k8s.io/enhancements/keps/sig-node/688-pod-overhead/README.md">https://git.k8s.io/enhancements/keps/sig-node/688-pod-overhead/README.md</a></p>
</td>
</tr>
<tr>
<td>
<code>topologySpreadConstraints</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#topologyspreadconstraint-v1-core">
[]Kubernetes core/v1.TopologySpreadConstraint
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>TopologySpreadConstraints describes how a group of pods ought to spread across topology
domains. Scheduler will schedule pods in a way which abides by the constraints.
All topologySpreadConstraints are ANDed.</p>
</td>
</tr>
<tr>
<td>
<code>setHostnameAsFQDN</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>If true the pod&rsquo;s hostname will be configured as the pod&rsquo;s FQDN, rather than the leaf name (the default).
In Linux containers, this means setting the FQDN in the hostname field of the kernel (the nodename field of struct utsname).
In Windows containers, this means setting the registry value of hostname for the registry key HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Services\Tcpip\Parameters to FQDN.
If a pod does not have FQDN, this has no effect.
Default to false.</p>
</td>
</tr>
<tr>
<td>
<code>os</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#podos-v1-core">
Kubernetes core/v1.PodOS
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the OS of the containers in the pod.
Some pod and container fields are restricted if this is set.</p>
<p>If the OS field is set to linux, the following fields must be unset:
-securityContext.windowsOptions</p>
<p>If the OS field is set to windows, following fields must be unset:
- spec.hostPID
- spec.hostIPC
- spec.hostUsers
- spec.securityContext.appArmorProfile
- spec.securityContext.seLinuxOptions
- spec.securityContext.seccompProfile
- spec.securityContext.fsGroup
- spec.securityContext.fsGroupChangePolicy
- spec.securityContext.sysctls
- spec.shareProcessNamespace
- spec.securityContext.runAsUser
- spec.securityContext.runAsGroup
- spec.securityContext.supplementalGroups
- spec.securityContext.supplementalGroupsPolicy
- spec.containers[<em>].securityContext.appArmorProfile
- spec.containers[</em>].securityContext.seLinuxOptions
- spec.containers[<em>].securityContext.seccompProfile
- spec.containers[</em>].securityContext.capabilities
- spec.containers[<em>].securityContext.readOnlyRootFilesystem
- spec.containers[</em>].securityContext.privileged
- spec.containers[<em>].securityContext.allowPrivilegeEscalation
- spec.containers[</em>].securityContext.procMount
- spec.containers[<em>].securityContext.runAsUser
- spec.containers[</em>].securityContext.runAsGroup</p>
</td>
</tr>
<tr>
<td>
<code>hostUsers</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Use the host&rsquo;s user namespace.
Optional: Default to true.
If set to true or not present, the pod will be run in the host user namespace, useful
for when the pod needs a feature only available to the host user namespace, such as
loading a kernel module with CAP_SYS_MODULE.
When set to false, a new userns is created for the pod. Setting false is useful for
mitigating container breakout vulnerabilities even allowing users to run their
containers as root without actually having root privileges on the host.
This field is alpha-level and is only honored by servers that enable the UserNamespacesSupport feature.</p>
</td>
</tr>
<tr>
<td>
<code>schedulingGates</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#podschedulinggate-v1-core">
[]Kubernetes core/v1.PodSchedulingGate
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>SchedulingGates is an opaque list of values that if specified will block scheduling the pod.
If schedulingGates is not empty, the pod will stay in the SchedulingGated state and the
scheduler will not attempt to schedule the pod.</p>
<p>SchedulingGates can only be set at pod creation time, and be removed only afterwards.</p>
</td>
</tr>
<tr>
<td>
<code>resourceClaims</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#podresourceclaim-v1-core">
[]Kubernetes core/v1.PodResourceClaim
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>ResourceClaims defines which ResourceClaims must be allocated
and reserved before the Pod is allowed to start. The resources
will be made available to those containers which consume them
by name.</p>
<p>This is an alpha field and requires enabling the
DynamicResourceAllocation feature gate.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>resources</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#resourcerequirements-v1-core">
Kubernetes core/v1.ResourceRequirements
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Resources is the total amount of CPU and Memory resources required by all
containers in the pod. It supports specifying Requests and Limits for
&ldquo;cpu&rdquo; and &ldquo;memory&rdquo; resource names only. ResourceClaims are not supported.</p>
<p>This field enables fine-grained control over resource allocation for the
entire pod, allowing resource sharing among containers in a pod.
TODO: For beta graduation, expand this comment with a detailed explanation.</p>
<p>This is an alpha field and requires enabling the PodLevelResources feature
gate.</p>
</td>
</tr>
</table>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1.PodDefaulter">PodDefaulter
</h3>
<div>
<p>PodDefaulter is a callback to validate a Pod.</p>
</div>
<h3 id="duck.knative.dev/v1.PodSpecDefaulter">PodSpecDefaulter
</h3>
<div>
<p>PodSpecDefaulter is a callback to validate a PodSpecable.</p>
</div>
<h3 id="duck.knative.dev/v1.PodSpecValidator">PodSpecValidator
</h3>
<div>
<p>PodSpecValidator is a callback to validate a PodSpecable.</p>
</div>
<h3 id="duck.knative.dev/v1.PodSpecable">PodSpecable
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1.WithPodSpec">WithPodSpec</a>)
</p>
<div>
<p>PodSpecable is implemented by types containing a PodTemplateSpec
in the manner of ReplicaSet, Deployment, DaemonSet, StatefulSet.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Standard object&rsquo;s metadata.
More info: <a href="https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata">https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata</a></p>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#podspec-v1-core">
Kubernetes core/v1.PodSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specification of the desired behavior of the pod.
More info: <a href="https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status">https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status</a></p>
<br/>
<br/>
<table>
<tr>
<td>
<code>volumes</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#volume-v1-core">
[]Kubernetes core/v1.Volume
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>List of volumes that can be mounted by containers belonging to the pod.
More info: <a href="https://kubernetes.io/docs/concepts/storage/volumes">https://kubernetes.io/docs/concepts/storage/volumes</a></p>
</td>
</tr>
<tr>
<td>
<code>initContainers</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#container-v1-core">
[]Kubernetes core/v1.Container
</a>
</em>
</td>
<td>
<p>List of initialization containers belonging to the pod.
Init containers are executed in order prior to containers being started. If any
init container fails, the pod is considered to have failed and is handled according
to its restartPolicy. The name for an init container or normal container must be
unique among all containers.
Init containers may not have Lifecycle actions, Readiness probes, Liveness probes, or Startup probes.
The resourceRequirements of an init container are taken into account during scheduling
by finding the highest request/limit for each resource type, and then using the max of
of that value or the sum of the normal containers. Limits are applied to init containers
in a similar fashion.
Init containers cannot currently be added or removed.
Cannot be updated.
More info: <a href="https://kubernetes.io/docs/concepts/workloads/pods/init-containers/">https://kubernetes.io/docs/concepts/workloads/pods/init-containers/</a></p>
</td>
</tr>
<tr>
<td>
<code>containers</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#container-v1-core">
[]Kubernetes core/v1.Container
</a>
</em>
</td>
<td>
<p>List of containers belonging to the pod.
Containers cannot currently be added or removed.
There must be at least one container in a Pod.
Cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>ephemeralContainers</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#ephemeralcontainer-v1-core">
[]Kubernetes core/v1.EphemeralContainer
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>List of ephemeral containers run in this pod. Ephemeral containers may be run in an existing
pod to perform user-initiated actions such as debugging. This list cannot be specified when
creating a pod, and it cannot be modified by updating the pod spec. In order to add an
ephemeral container to an existing pod, use the pod&rsquo;s ephemeralcontainers subresource.</p>
</td>
</tr>
<tr>
<td>
<code>restartPolicy</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#restartpolicy-v1-core">
Kubernetes core/v1.RestartPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Restart policy for all containers within the pod.
One of Always, OnFailure, Never. In some contexts, only a subset of those values may be permitted.
Default to Always.
More info: <a href="https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy">https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy</a></p>
</td>
</tr>
<tr>
<td>
<code>terminationGracePeriodSeconds</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete request.
Value must be non-negative integer. The value zero indicates stop immediately via
the kill signal (no opportunity to shut down).
If this value is nil, the default grace period will be used instead.
The grace period is the duration in seconds after the processes running in the pod are sent
a termination signal and the time when the processes are forcibly halted with a kill signal.
Set this value longer than the expected cleanup time for your process.
Defaults to 30 seconds.</p>
</td>
</tr>
<tr>
<td>
<code>activeDeadlineSeconds</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>Optional duration in seconds the pod may be active on the node relative to
StartTime before the system will actively try to mark it failed and kill associated containers.
Value must be a positive integer.</p>
</td>
</tr>
<tr>
<td>
<code>dnsPolicy</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#dnspolicy-v1-core">
Kubernetes core/v1.DNSPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Set DNS policy for the pod.
Defaults to &ldquo;ClusterFirst&rdquo;.
Valid values are &lsquo;ClusterFirstWithHostNet&rsquo;, &lsquo;ClusterFirst&rsquo;, &lsquo;Default&rsquo; or &lsquo;None&rsquo;.
DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy.
To have DNS options set along with hostNetwork, you have to specify DNS policy
explicitly to &lsquo;ClusterFirstWithHostNet&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>nodeSelector</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>NodeSelector is a selector which must be true for the pod to fit on a node.
Selector which must match a node&rsquo;s labels for the pod to be scheduled on that node.
More info: <a href="https://kubernetes.io/docs/concepts/configuration/assign-pod-node/">https://kubernetes.io/docs/concepts/configuration/assign-pod-node/</a></p>
</td>
</tr>
<tr>
<td>
<code>serviceAccountName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ServiceAccountName is the name of the ServiceAccount to use to run this pod.
More info: <a href="https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/">https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/</a></p>
</td>
</tr>
<tr>
<td>
<code>serviceAccount</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>DeprecatedServiceAccount is a deprecated alias for ServiceAccountName.
Deprecated: Use serviceAccountName instead.</p>
</td>
</tr>
<tr>
<td>
<code>automountServiceAccountToken</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>AutomountServiceAccountToken indicates whether a service account token should be automatically mounted.</p>
</td>
</tr>
<tr>
<td>
<code>nodeName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>NodeName indicates in which node this pod is scheduled.
If empty, this pod is a candidate for scheduling by the scheduler defined in schedulerName.
Once this field is set, the kubelet for this node becomes responsible for the lifecycle of this pod.
This field should not be used to express a desire for the pod to be scheduled on a specific node.
<a href="https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#nodename">https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#nodename</a></p>
</td>
</tr>
<tr>
<td>
<code>hostNetwork</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Host networking requested for this pod. Use the host&rsquo;s network namespace.
If this option is set, the ports that will be used must be specified.
Default to false.</p>
</td>
</tr>
<tr>
<td>
<code>hostPID</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Use the host&rsquo;s pid namespace.
Optional: Default to false.</p>
</td>
</tr>
<tr>
<td>
<code>hostIPC</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Use the host&rsquo;s ipc namespace.
Optional: Default to false.</p>
</td>
</tr>
<tr>
<td>
<code>shareProcessNamespace</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Share a single process namespace between all of the containers in a pod.
When this is set containers will be able to view and signal processes from other containers
in the same pod, and the first process in each container will not be assigned PID 1.
HostPID and ShareProcessNamespace cannot both be set.
Optional: Default to false.</p>
</td>
</tr>
<tr>
<td>
<code>securityContext</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#podsecuritycontext-v1-core">
Kubernetes core/v1.PodSecurityContext
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>SecurityContext holds pod-level security attributes and common container settings.
Optional: Defaults to empty.  See type description for default values of each field.</p>
</td>
</tr>
<tr>
<td>
<code>imagePullSecrets</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#localobjectreference-v1-core">
[]Kubernetes core/v1.LocalObjectReference
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.
If specified, these secrets will be passed to individual puller implementations for them to use.
More info: <a href="https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod">https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod</a></p>
</td>
</tr>
<tr>
<td>
<code>hostname</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the hostname of the Pod
If not specified, the pod&rsquo;s hostname will be set to a system-defined value.</p>
</td>
</tr>
<tr>
<td>
<code>subdomain</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>If specified, the fully qualified Pod hostname will be &ldquo;<hostname>.<subdomain>.<pod namespace>.svc.<cluster domain>&rdquo;.
If not specified, the pod will not have a domainname at all.</p>
</td>
</tr>
<tr>
<td>
<code>affinity</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#affinity-v1-core">
Kubernetes core/v1.Affinity
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>If specified, the pod&rsquo;s scheduling constraints</p>
</td>
</tr>
<tr>
<td>
<code>schedulerName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>If specified, the pod will be dispatched by specified scheduler.
If not specified, the pod will be dispatched by default scheduler.</p>
</td>
</tr>
<tr>
<td>
<code>tolerations</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#toleration-v1-core">
[]Kubernetes core/v1.Toleration
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>If specified, the pod&rsquo;s tolerations.</p>
</td>
</tr>
<tr>
<td>
<code>hostAliases</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#hostalias-v1-core">
[]Kubernetes core/v1.HostAlias
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>HostAliases is an optional list of hosts and IPs that will be injected into the pod&rsquo;s hosts
file if specified.</p>
</td>
</tr>
<tr>
<td>
<code>priorityClassName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>If specified, indicates the pod&rsquo;s priority. &ldquo;system-node-critical&rdquo; and
&ldquo;system-cluster-critical&rdquo; are two special keywords which indicate the
highest priorities with the former being the highest priority. Any other
name must be defined by creating a PriorityClass object with that name.
If not specified, the pod priority will be default or zero if there is no
default.</p>
</td>
</tr>
<tr>
<td>
<code>priority</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>The priority value. Various system components use this field to find the
priority of the pod. When Priority Admission Controller is enabled, it
prevents users from setting this field. The admission controller populates
this field from PriorityClassName.
The higher the value, the higher the priority.</p>
</td>
</tr>
<tr>
<td>
<code>dnsConfig</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#poddnsconfig-v1-core">
Kubernetes core/v1.PodDNSConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the DNS parameters of a pod.
Parameters specified here will be merged to the generated DNS
configuration based on DNSPolicy.</p>
</td>
</tr>
<tr>
<td>
<code>readinessGates</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#podreadinessgate-v1-core">
[]Kubernetes core/v1.PodReadinessGate
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>If specified, all readiness gates will be evaluated for pod readiness.
A pod is ready when all its containers are ready AND
all conditions specified in the readiness gates have status equal to &ldquo;True&rdquo;
More info: <a href="https://git.k8s.io/enhancements/keps/sig-network/580-pod-readiness-gates">https://git.k8s.io/enhancements/keps/sig-network/580-pod-readiness-gates</a></p>
</td>
</tr>
<tr>
<td>
<code>runtimeClassName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>RuntimeClassName refers to a RuntimeClass object in the node.k8s.io group, which should be used
to run this pod.  If no RuntimeClass resource matches the named class, the pod will not be run.
If unset or empty, the &ldquo;legacy&rdquo; RuntimeClass will be used, which is an implicit class with an
empty definition that uses the default runtime handler.
More info: <a href="https://git.k8s.io/enhancements/keps/sig-node/585-runtime-class">https://git.k8s.io/enhancements/keps/sig-node/585-runtime-class</a></p>
</td>
</tr>
<tr>
<td>
<code>enableServiceLinks</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>EnableServiceLinks indicates whether information about services should be injected into pod&rsquo;s
environment variables, matching the syntax of Docker links.
Optional: Defaults to true.</p>
</td>
</tr>
<tr>
<td>
<code>preemptionPolicy</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#preemptionpolicy-v1-core">
Kubernetes core/v1.PreemptionPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>PreemptionPolicy is the Policy for preempting pods with lower priority.
One of Never, PreemptLowerPriority.
Defaults to PreemptLowerPriority if unset.</p>
</td>
</tr>
<tr>
<td>
<code>overhead</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#resourcelist-v1-core">
Kubernetes core/v1.ResourceList
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Overhead represents the resource overhead associated with running a pod for a given RuntimeClass.
This field will be autopopulated at admission time by the RuntimeClass admission controller. If
the RuntimeClass admission controller is enabled, overhead must not be set in Pod create requests.
The RuntimeClass admission controller will reject Pod create requests which have the overhead already
set. If RuntimeClass is configured and selected in the PodSpec, Overhead will be set to the value
defined in the corresponding RuntimeClass, otherwise it will remain unset and treated as zero.
More info: <a href="https://git.k8s.io/enhancements/keps/sig-node/688-pod-overhead/README.md">https://git.k8s.io/enhancements/keps/sig-node/688-pod-overhead/README.md</a></p>
</td>
</tr>
<tr>
<td>
<code>topologySpreadConstraints</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#topologyspreadconstraint-v1-core">
[]Kubernetes core/v1.TopologySpreadConstraint
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>TopologySpreadConstraints describes how a group of pods ought to spread across topology
domains. Scheduler will schedule pods in a way which abides by the constraints.
All topologySpreadConstraints are ANDed.</p>
</td>
</tr>
<tr>
<td>
<code>setHostnameAsFQDN</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>If true the pod&rsquo;s hostname will be configured as the pod&rsquo;s FQDN, rather than the leaf name (the default).
In Linux containers, this means setting the FQDN in the hostname field of the kernel (the nodename field of struct utsname).
In Windows containers, this means setting the registry value of hostname for the registry key HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Services\Tcpip\Parameters to FQDN.
If a pod does not have FQDN, this has no effect.
Default to false.</p>
</td>
</tr>
<tr>
<td>
<code>os</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#podos-v1-core">
Kubernetes core/v1.PodOS
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the OS of the containers in the pod.
Some pod and container fields are restricted if this is set.</p>
<p>If the OS field is set to linux, the following fields must be unset:
-securityContext.windowsOptions</p>
<p>If the OS field is set to windows, following fields must be unset:
- spec.hostPID
- spec.hostIPC
- spec.hostUsers
- spec.securityContext.appArmorProfile
- spec.securityContext.seLinuxOptions
- spec.securityContext.seccompProfile
- spec.securityContext.fsGroup
- spec.securityContext.fsGroupChangePolicy
- spec.securityContext.sysctls
- spec.shareProcessNamespace
- spec.securityContext.runAsUser
- spec.securityContext.runAsGroup
- spec.securityContext.supplementalGroups
- spec.securityContext.supplementalGroupsPolicy
- spec.containers[<em>].securityContext.appArmorProfile
- spec.containers[</em>].securityContext.seLinuxOptions
- spec.containers[<em>].securityContext.seccompProfile
- spec.containers[</em>].securityContext.capabilities
- spec.containers[<em>].securityContext.readOnlyRootFilesystem
- spec.containers[</em>].securityContext.privileged
- spec.containers[<em>].securityContext.allowPrivilegeEscalation
- spec.containers[</em>].securityContext.procMount
- spec.containers[<em>].securityContext.runAsUser
- spec.containers[</em>].securityContext.runAsGroup</p>
</td>
</tr>
<tr>
<td>
<code>hostUsers</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Use the host&rsquo;s user namespace.
Optional: Default to true.
If set to true or not present, the pod will be run in the host user namespace, useful
for when the pod needs a feature only available to the host user namespace, such as
loading a kernel module with CAP_SYS_MODULE.
When set to false, a new userns is created for the pod. Setting false is useful for
mitigating container breakout vulnerabilities even allowing users to run their
containers as root without actually having root privileges on the host.
This field is alpha-level and is only honored by servers that enable the UserNamespacesSupport feature.</p>
</td>
</tr>
<tr>
<td>
<code>schedulingGates</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#podschedulinggate-v1-core">
[]Kubernetes core/v1.PodSchedulingGate
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>SchedulingGates is an opaque list of values that if specified will block scheduling the pod.
If schedulingGates is not empty, the pod will stay in the SchedulingGated state and the
scheduler will not attempt to schedule the pod.</p>
<p>SchedulingGates can only be set at pod creation time, and be removed only afterwards.</p>
</td>
</tr>
<tr>
<td>
<code>resourceClaims</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#podresourceclaim-v1-core">
[]Kubernetes core/v1.PodResourceClaim
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>ResourceClaims defines which ResourceClaims must be allocated
and reserved before the Pod is allowed to start. The resources
will be made available to those containers which consume them
by name.</p>
<p>This is an alpha field and requires enabling the
DynamicResourceAllocation feature gate.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>resources</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#resourcerequirements-v1-core">
Kubernetes core/v1.ResourceRequirements
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Resources is the total amount of CPU and Memory resources required by all
containers in the pod. It supports specifying Requests and Limits for
&ldquo;cpu&rdquo; and &ldquo;memory&rdquo; resource names only. ResourceClaims are not supported.</p>
<p>This field enables fine-grained control over resource allocation for the
entire pod, allowing resource sharing among containers in a pod.
TODO: For beta graduation, expand this comment with a detailed explanation.</p>
<p>This is an alpha field and requires enabling the PodLevelResources feature
gate.</p>
</td>
</tr>
</table>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1.PodValidator">PodValidator
</h3>
<div>
<p>PodValidator is a callback to validate Pods.</p>
</div>
<h3 id="duck.knative.dev/v1.Source">Source
</h3>
<div>
<p>Source is the minimum resource shape to adhere to the Source Specification.
This duck type is intended to allow implementors of Sources and
Importers to verify their own resources meet the expectations.
This is not a real resource.
NOTE: The Source Specification is in progress and the shape and names could
be modified until it has been accepted.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#duck.knative.dev/v1.SourceSpec">
SourceSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>sink</code><br/>
<em>
<a href="#duck.knative.dev/v1.Destination">
Destination
</a>
</em>
</td>
<td>
<p>Sink is a reference to an object that will resolve to a uri to use as the sink.</p>
</td>
</tr>
<tr>
<td>
<code>ceOverrides</code><br/>
<em>
<a href="#duck.knative.dev/v1.CloudEventOverrides">
CloudEventOverrides
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>CloudEventOverrides defines overrides to control the output format and
modifications of the event sent to the sink.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#duck.knative.dev/v1.SourceStatus">
SourceStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1.SourceSpec">SourceSpec
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1.Source">Source</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>sink</code><br/>
<em>
<a href="#duck.knative.dev/v1.Destination">
Destination
</a>
</em>
</td>
<td>
<p>Sink is a reference to an object that will resolve to a uri to use as the sink.</p>
</td>
</tr>
<tr>
<td>
<code>ceOverrides</code><br/>
<em>
<a href="#duck.knative.dev/v1.CloudEventOverrides">
CloudEventOverrides
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>CloudEventOverrides defines overrides to control the output format and
modifications of the event sent to the sink.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1.SourceStatus">SourceStatus
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1.Source">Source</a>)
</p>
<div>
<p>SourceStatus shows how we expect folks to embed Addressable in
their Status field.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>Status</code><br/>
<em>
<a href="#duck.knative.dev/v1.Status">
Status
</a>
</em>
</td>
<td>
<p>
(Members of <code>Status</code> are embedded into this type.)
</p>
<p>inherits Status, which currently provides:
* ObservedGeneration - the &lsquo;Generation&rsquo; of the Service that was last
processed by the controller.
* Conditions - the latest available observations of a resource&rsquo;s current
state.</p>
</td>
</tr>
<tr>
<td>
<code>sinkUri</code><br/>
<em>
knative.dev/serving/pkg/apis.URL
</em>
</td>
<td>
<em>(Optional)</em>
<p>SinkURI is the current active sink URI that has been configured for the
Source.</p>
</td>
</tr>
<tr>
<td>
<code>ceAttributes</code><br/>
<em>
<a href="#duck.knative.dev/v1.CloudEventAttributes">
[]CloudEventAttributes
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>CloudEventAttributes are the specific attributes that the Source uses
as part of its CloudEvents.</p>
</td>
</tr>
<tr>
<td>
<code>sinkCACerts</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>SinkCACerts are Certification Authority (CA) certificates in PEM format
according to <a href="https://www.rfc-editor.org/rfc/rfc7468">https://www.rfc-editor.org/rfc/rfc7468</a>.</p>
</td>
</tr>
<tr>
<td>
<code>sinkAudience</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>SinkAudience is the OIDC audience of the sink.</p>
</td>
</tr>
<tr>
<td>
<code>auth</code><br/>
<em>
<a href="#duck.knative.dev/v1.AuthStatus">
AuthStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Auth defines the attributes that provide the generated service account
name in the resource status.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1.Status">Status
</h3>
<p>
(<em>Appears on:</em><a href="#autoscaling.internal.knative.dev/v1alpha1.MetricStatus">MetricStatus</a>, <a href="#autoscaling.internal.knative.dev/v1alpha1.PodAutoscalerStatus">PodAutoscalerStatus</a>, <a href="#duck.knative.dev/v1.KResource">KResource</a>, <a href="#duck.knative.dev/v1.SourceStatus">SourceStatus</a>, <a href="#serving.knative.dev/v1.ConfigurationStatus">ConfigurationStatus</a>, <a href="#serving.knative.dev/v1.RevisionStatus">RevisionStatus</a>, <a href="#serving.knative.dev/v1.RouteStatus">RouteStatus</a>, <a href="#serving.knative.dev/v1.ServiceStatus">ServiceStatus</a>, <a href="#serving.knative.dev/v1beta1.DomainMappingStatus">DomainMappingStatus</a>)
</p>
<div>
<p>Status shows how we expect folks to embed Conditions in
their Status field.
WARNING: Adding fields to this struct will add them to all Knative resources.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>observedGeneration</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>ObservedGeneration is the &lsquo;Generation&rsquo; of the Service that
was last processed by the controller.</p>
</td>
</tr>
<tr>
<td>
<code>conditions</code><br/>
<em>
<a href="#duck.knative.dev/v1.Conditions">
Conditions
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Conditions the latest available observations of a resource&rsquo;s current state.</p>
</td>
</tr>
<tr>
<td>
<code>annotations</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Annotations is additional Status fields for the Resource to save some
additional State as well as convey more information to the user. This is
roughly akin to Annotations on any k8s resource, just the reconciler conveying
richer information outwards.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1.WithPod">WithPod
</h3>
<div>
<p>WithPod is the shell that demonstrates how PodSpecable types wrap
a PodSpec.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#duck.knative.dev/v1.WithPodSpec">
WithPodSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>template</code><br/>
<em>
<a href="#duck.knative.dev/v1.PodSpecable">
PodSpecable
</a>
</em>
</td>
<td>
</td>
</tr>
</table>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1.WithPodSpec">WithPodSpec
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1.WithPod">WithPod</a>)
</p>
<div>
<p>WithPodSpec is the shell around the PodSpecable within WithPod.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>template</code><br/>
<em>
<a href="#duck.knative.dev/v1.PodSpecable">
PodSpecable
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<hr/>
<h2 id="duck.knative.dev/v1beta1">duck.knative.dev/v1beta1</h2>
<div>
</div>
Resource Types:
<ul></ul>
<h3 id="duck.knative.dev/v1beta1.AddressStatus">AddressStatus
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1beta1.AddressableType">AddressableType</a>)
</p>
<div>
<p>AddressStatus shows how we expect folks to embed Addressable in
their Status field.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>address</code><br/>
<em>
<a href="#duck.knative.dev/v1beta1.Addressable">
Addressable
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Address is a single Addressable address.
If Addresses is present, Address will be ignored by clients.</p>
</td>
</tr>
<tr>
<td>
<code>addresses</code><br/>
<em>
<a href="#duck.knative.dev/v1beta1.Addressable">
[]Addressable
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Addresses is a list of addresses for different protocols (HTTP and HTTPS)
If Addresses is present, Address must be ignored by clients.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1beta1.Addressable">Addressable
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1beta1.AddressStatus">AddressStatus</a>)
</p>
<div>
<p>Addressable provides a generic mechanism for a custom resource
definition to indicate a destination for message delivery.</p>
<p>Addressable is the schema for the destination information. This is
typically stored in the object&rsquo;s <code>status</code>, as this information may
be generated by the controller.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Name is the name of the address.</p>
</td>
</tr>
<tr>
<td>
<code>url</code><br/>
<em>
knative.dev/serving/pkg/apis.URL
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>CACerts</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>CACerts is the Certification Authority (CA) certificates in PEM format
according to <a href="https://www.rfc-editor.org/rfc/rfc7468">https://www.rfc-editor.org/rfc/rfc7468</a>.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1beta1.AddressableType">AddressableType
</h3>
<div>
<p>AddressableType is a skeleton type wrapping Addressable in the manner we expect
resource writers defining compatible resources to embed it.  We will
typically use this type to deserialize Addressable ObjectReferences and
access the Addressable data.  This is not a real resource.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#duck.knative.dev/v1beta1.AddressStatus">
AddressStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1beta1.Binding">Binding
</h3>
<div>
<p>Binding is a duck type that specifies the partial schema to which all
Binding implementations should adhere.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#duck.knative.dev/v1beta1.BindingSpec">
BindingSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>subject</code><br/>
<em>
knative.dev/serving/pkg/tracker.Reference
</em>
</td>
<td>
<p>Subject references the resource(s) whose &ldquo;runtime contract&rdquo; should be
augmented by Binding implementations.</p>
</td>
</tr>
</table>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1beta1.BindingSpec">BindingSpec
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1beta1.Binding">Binding</a>)
</p>
<div>
<p>BindingSpec specifies the spec portion of the Binding partial-schema.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>subject</code><br/>
<em>
knative.dev/serving/pkg/tracker.Reference
</em>
</td>
<td>
<p>Subject references the resource(s) whose &ldquo;runtime contract&rdquo; should be
augmented by Binding implementations.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1beta1.CloudEventOverrides">CloudEventOverrides
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1beta1.SourceSpec">SourceSpec</a>)
</p>
<div>
<p>CloudEventOverrides defines arguments for a Source that control the output
format of the CloudEvents produced by the Source.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>extensions</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Extensions specify what attribute are added or overridden on the
outbound event. Each <code>Extensions</code> key-value pair are set on the event as
an attribute extension independently.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1beta1.Conditions">Conditions
(<code>[]knative.dev/serving/pkg/apis.Condition</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1beta1.Status">Status</a>)
</p>
<div>
<p>Conditions is a simple wrapper around apis.Conditions to implement duck.Implementable.</p>
</div>
<h3 id="duck.knative.dev/v1beta1.Destination">Destination
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1beta1.SourceSpec">SourceSpec</a>)
</p>
<div>
<p>Destination represents a target of an invocation over HTTP.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ref</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#objectreference-v1-core">
Kubernetes core/v1.ObjectReference
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Ref points to an Addressable.</p>
</td>
</tr>
<tr>
<td>
<code>apiVersion</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>namespace</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>uri</code><br/>
<em>
knative.dev/serving/pkg/apis.URL
</em>
</td>
<td>
<em>(Optional)</em>
<p>URI can be an absolute URL(non-empty scheme and non-empty host) pointing to the target or a relative URI. Relative URIs will be resolved using the base URI retrieved from Ref.</p>
</td>
</tr>
<tr>
<td>
<code>CACerts</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>CACerts are Certification Authority (CA) certificates in PEM format
according to <a href="https://www.rfc-editor.org/rfc/rfc7468">https://www.rfc-editor.org/rfc/rfc7468</a>.
If set, these CAs are appended to the set of CAs provided
by the Addressable target, if any.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1beta1.KResource">KResource
</h3>
<div>
<p>KResource is a skeleton type wrapping Conditions in the manner we expect
resource writers defining compatible resources to embed it.  We will
typically use this type to deserialize Conditions ObjectReferences and
access the Conditions data.  This is not a real resource.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#duck.knative.dev/v1beta1.Status">
Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1beta1.Source">Source
</h3>
<div>
<p>Source is the minimum resource shape to adhere to the Source Specification.
This duck type is intended to allow implementors of Sources and
Importers to verify their own resources meet the expectations.
This is not a real resource.
NOTE: The Source Specification is in progress and the shape and names could
be modified until it has been accepted.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#duck.knative.dev/v1beta1.SourceSpec">
SourceSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>sink</code><br/>
<em>
<a href="#duck.knative.dev/v1beta1.Destination">
Destination
</a>
</em>
</td>
<td>
<p>Sink is a reference to an object that will resolve to a domain name or a
URI directly to use as the sink.</p>
</td>
</tr>
<tr>
<td>
<code>ceOverrides</code><br/>
<em>
<a href="#duck.knative.dev/v1beta1.CloudEventOverrides">
CloudEventOverrides
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>CloudEventOverrides defines overrides to control the output format and
modifications of the event sent to the sink.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#duck.knative.dev/v1beta1.SourceStatus">
SourceStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1beta1.SourceSpec">SourceSpec
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1beta1.Source">Source</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>sink</code><br/>
<em>
<a href="#duck.knative.dev/v1beta1.Destination">
Destination
</a>
</em>
</td>
<td>
<p>Sink is a reference to an object that will resolve to a domain name or a
URI directly to use as the sink.</p>
</td>
</tr>
<tr>
<td>
<code>ceOverrides</code><br/>
<em>
<a href="#duck.knative.dev/v1beta1.CloudEventOverrides">
CloudEventOverrides
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>CloudEventOverrides defines overrides to control the output format and
modifications of the event sent to the sink.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1beta1.SourceStatus">SourceStatus
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1beta1.Source">Source</a>)
</p>
<div>
<p>SourceStatus shows how we expect folks to embed Addressable in
their Status field.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>Status</code><br/>
<em>
<a href="#duck.knative.dev/v1beta1.Status">
Status
</a>
</em>
</td>
<td>
<p>
(Members of <code>Status</code> are embedded into this type.)
</p>
<p>inherits duck/v1beta1 Status, which currently provides:
* ObservedGeneration - the &lsquo;Generation&rsquo; of the Service that was last
processed by the controller.
* Conditions - the latest available observations of a resource&rsquo;s current
state.</p>
</td>
</tr>
<tr>
<td>
<code>sinkUri</code><br/>
<em>
knative.dev/serving/pkg/apis.URL
</em>
</td>
<td>
<em>(Optional)</em>
<p>SinkURI is the current active sink URI that has been configured for the
Source.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1beta1.Status">Status
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1beta1.KResource">KResource</a>, <a href="#duck.knative.dev/v1beta1.SourceStatus">SourceStatus</a>)
</p>
<div>
<p>Status shows how we expect folks to embed Conditions in
their Status field.
WARNING: Adding fields to this struct will add them to all Knative resources.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>observedGeneration</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>ObservedGeneration is the &lsquo;Generation&rsquo; of the Service that
was last processed by the controller.</p>
</td>
</tr>
<tr>
<td>
<code>conditions</code><br/>
<em>
<a href="#duck.knative.dev/v1beta1.Conditions">
Conditions
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Conditions the latest available observations of a resource&rsquo;s current state.</p>
</td>
</tr>
<tr>
<td>
<code>annotations</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Annotations is additional Status fields for the Resource to save some
additional State as well as convey more information to the user. This is
roughly akin to Annotations on any k8s resource, just the reconciler conveying
richer information outwards.</p>
</td>
</tr>
</tbody>
</table>
<hr/>
<h2 id="serving.knative.dev/v1">serving.knative.dev/v1</h2>
<div>
<p>Package v1 contains the Serving v1 API types.</p>
</div>
Resource Types:
<ul><li>
<a href="#serving.knative.dev/v1.Configuration">Configuration</a>
</li><li>
<a href="#serving.knative.dev/v1.Revision">Revision</a>
</li><li>
<a href="#serving.knative.dev/v1.Route">Route</a>
</li><li>
<a href="#serving.knative.dev/v1.Service">Service</a>
</li></ul>
<h3 id="serving.knative.dev/v1.Configuration">Configuration
</h3>
<div>
<p>Configuration represents the &ldquo;floating HEAD&rdquo; of a linear history of Revisions.
Users create new Revisions by updating the Configuration&rsquo;s spec.
The &ldquo;latest created&rdquo; revision&rsquo;s name is available under status, as is the
&ldquo;latest ready&rdquo; revision&rsquo;s name.
See also: <a href="https://github.com/knative/serving/blob/main/docs/spec/overview.md#configuration">https://github.com/knative/serving/blob/main/docs/spec/overview.md#configuration</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>
serving.knative.dev/v1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>Configuration</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
<em>(Optional)</em>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#serving.knative.dev/v1.ConfigurationSpec">
ConfigurationSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<br/>
<br/>
<table>
<tr>
<td>
<code>template</code><br/>
<em>
<a href="#serving.knative.dev/v1.RevisionTemplateSpec">
RevisionTemplateSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Template holds the latest specification for the Revision to be stamped out.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#serving.knative.dev/v1.ConfigurationStatus">
ConfigurationStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
</tbody>
</table>
<h3 id="serving.knative.dev/v1.Revision">Revision
</h3>
<div>
<p>Revision is an immutable snapshot of code and configuration.  A revision
references a container image. Revisions are created by updates to a
Configuration.</p>
<p>See also: <a href="https://github.com/knative/serving/blob/main/docs/spec/overview.md#revision">https://github.com/knative/serving/blob/main/docs/spec/overview.md#revision</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>
serving.knative.dev/v1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>Revision</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
<em>(Optional)</em>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#serving.knative.dev/v1.RevisionSpec">
RevisionSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<br/>
<br/>
<table>
<tr>
<td>
<code>PodSpec</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#podspec-v1-core">
Kubernetes core/v1.PodSpec
</a>
</em>
</td>
<td>
<p>
(Members of <code>PodSpec</code> are embedded into this type.)
</p>
</td>
</tr>
<tr>
<td>
<code>containerConcurrency</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>ContainerConcurrency specifies the maximum allowed in-flight (concurrent)
requests per container of the Revision.  Defaults to <code>0</code> which means
concurrency to the application is not limited, and the system decides the
target concurrency for the autoscaler.</p>
</td>
</tr>
<tr>
<td>
<code>timeoutSeconds</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>TimeoutSeconds is the maximum duration in seconds that the request instance
is allowed to respond to a request. If unspecified, a system default will
be provided.</p>
</td>
</tr>
<tr>
<td>
<code>responseStartTimeoutSeconds</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>ResponseStartTimeoutSeconds is the maximum duration in seconds that the request
routing layer will wait for a request delivered to a container to begin
sending any network traffic.</p>
</td>
</tr>
<tr>
<td>
<code>idleTimeoutSeconds</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>IdleTimeoutSeconds is the maximum duration in seconds a request will be allowed
to stay open while not receiving any bytes from the user&rsquo;s application. If
unspecified, a system default will be provided.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#serving.knative.dev/v1.RevisionStatus">
RevisionStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
</tbody>
</table>
<h3 id="serving.knative.dev/v1.Route">Route
</h3>
<div>
<p>Route is responsible for configuring ingress over a collection of Revisions.
Some of the Revisions a Route distributes traffic over may be specified by
referencing the Configuration responsible for creating them; in these cases
the Route is additionally responsible for monitoring the Configuration for
&ldquo;latest ready revision&rdquo; changes, and smoothly rolling out latest revisions.
See also: <a href="https://github.com/knative/serving/blob/main/docs/spec/overview.md#route">https://github.com/knative/serving/blob/main/docs/spec/overview.md#route</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>
serving.knative.dev/v1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>Route</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
<em>(Optional)</em>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#serving.knative.dev/v1.RouteSpec">
RouteSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Spec holds the desired state of the Route (from the client).</p>
<br/>
<br/>
<table>
<tr>
<td>
<code>traffic</code><br/>
<em>
<a href="#serving.knative.dev/v1.TrafficTarget">
[]TrafficTarget
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Traffic specifies how to distribute traffic over a collection of
revisions and configurations.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#serving.knative.dev/v1.RouteStatus">
RouteStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Status communicates the observed state of the Route (from the controller).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="serving.knative.dev/v1.Service">Service
</h3>
<div>
<p>Service acts as a top-level container that manages a Route and Configuration
which implement a network service. Service exists to provide a singular
abstraction which can be access controlled, reasoned about, and which
encapsulates software lifecycle decisions such as rollout policy and
team resource ownership. Service acts only as an orchestrator of the
underlying Routes and Configurations (much as a kubernetes Deployment
orchestrates ReplicaSets), and its usage is optional but recommended.</p>
<p>The Service&rsquo;s controller will track the statuses of its owned Configuration
and Route, reflecting their statuses and conditions as its own.</p>
<p>See also: <a href="https://github.com/knative/serving/blob/main/docs/spec/overview.md#service">https://github.com/knative/serving/blob/main/docs/spec/overview.md#service</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>
serving.knative.dev/v1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>Service</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
<em>(Optional)</em>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#serving.knative.dev/v1.ServiceSpec">
ServiceSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<br/>
<br/>
<table>
<tr>
<td>
<code>ConfigurationSpec</code><br/>
<em>
<a href="#serving.knative.dev/v1.ConfigurationSpec">
ConfigurationSpec
</a>
</em>
</td>
<td>
<p>
(Members of <code>ConfigurationSpec</code> are embedded into this type.)
</p>
<p>ServiceSpec inlines an unrestricted ConfigurationSpec.</p>
</td>
</tr>
<tr>
<td>
<code>RouteSpec</code><br/>
<em>
<a href="#serving.knative.dev/v1.RouteSpec">
RouteSpec
</a>
</em>
</td>
<td>
<p>
(Members of <code>RouteSpec</code> are embedded into this type.)
</p>
<p>ServiceSpec inlines RouteSpec and restricts/defaults its fields
via webhook.  In particular, this spec can only reference this
Service&rsquo;s configuration and revisions (which also influences
defaults).</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#serving.knative.dev/v1.ServiceStatus">
ServiceStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
</tbody>
</table>
<h3 id="serving.knative.dev/v1.ConfigurationSpec">ConfigurationSpec
</h3>
<p>
(<em>Appears on:</em><a href="#serving.knative.dev/v1.Configuration">Configuration</a>, <a href="#serving.knative.dev/v1.ServiceSpec">ServiceSpec</a>)
</p>
<div>
<p>ConfigurationSpec holds the desired state of the Configuration (from the client).</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>template</code><br/>
<em>
<a href="#serving.knative.dev/v1.RevisionTemplateSpec">
RevisionTemplateSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Template holds the latest specification for the Revision to be stamped out.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="serving.knative.dev/v1.ConfigurationStatus">ConfigurationStatus
</h3>
<p>
(<em>Appears on:</em><a href="#serving.knative.dev/v1.Configuration">Configuration</a>)
</p>
<div>
<p>ConfigurationStatus communicates the observed state of the Configuration (from the controller).</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>Status</code><br/>
<em>
<a href="#duck.knative.dev/v1.Status">
Status
</a>
</em>
</td>
<td>
<p>
(Members of <code>Status</code> are embedded into this type.)
</p>
</td>
</tr>
<tr>
<td>
<code>ConfigurationStatusFields</code><br/>
<em>
<a href="#serving.knative.dev/v1.ConfigurationStatusFields">
ConfigurationStatusFields
</a>
</em>
</td>
<td>
<p>
(Members of <code>ConfigurationStatusFields</code> are embedded into this type.)
</p>
</td>
</tr>
</tbody>
</table>
<h3 id="serving.knative.dev/v1.ConfigurationStatusFields">ConfigurationStatusFields
</h3>
<p>
(<em>Appears on:</em><a href="#serving.knative.dev/v1.ConfigurationStatus">ConfigurationStatus</a>, <a href="#serving.knative.dev/v1.ServiceStatus">ServiceStatus</a>)
</p>
<div>
<p>ConfigurationStatusFields holds the fields of Configuration&rsquo;s status that
are not generally shared.  This is defined separately and inlined so that
other types can readily consume these fields via duck typing.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>latestReadyRevisionName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>LatestReadyRevisionName holds the name of the latest Revision stamped out
from this Configuration that has had its &ldquo;Ready&rdquo; condition become &ldquo;True&rdquo;.</p>
</td>
</tr>
<tr>
<td>
<code>latestCreatedRevisionName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>LatestCreatedRevisionName is the last revision that was created from this
Configuration. It might not be ready yet, for that use LatestReadyRevisionName.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="serving.knative.dev/v1.ContainerStatus">ContainerStatus
</h3>
<p>
(<em>Appears on:</em><a href="#serving.knative.dev/v1.RevisionStatus">RevisionStatus</a>)
</p>
<div>
<p>ContainerStatus holds the information of container name and image digest value</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>imageDigest</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="serving.knative.dev/v1.RevisionSpec">RevisionSpec
</h3>
<p>
(<em>Appears on:</em><a href="#serving.knative.dev/v1.Revision">Revision</a>, <a href="#serving.knative.dev/v1.RevisionTemplateSpec">RevisionTemplateSpec</a>)
</p>
<div>
<p>RevisionSpec holds the desired state of the Revision (from the client).</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>PodSpec</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#podspec-v1-core">
Kubernetes core/v1.PodSpec
</a>
</em>
</td>
<td>
<p>
(Members of <code>PodSpec</code> are embedded into this type.)
</p>
</td>
</tr>
<tr>
<td>
<code>containerConcurrency</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>ContainerConcurrency specifies the maximum allowed in-flight (concurrent)
requests per container of the Revision.  Defaults to <code>0</code> which means
concurrency to the application is not limited, and the system decides the
target concurrency for the autoscaler.</p>
</td>
</tr>
<tr>
<td>
<code>timeoutSeconds</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>TimeoutSeconds is the maximum duration in seconds that the request instance
is allowed to respond to a request. If unspecified, a system default will
be provided.</p>
</td>
</tr>
<tr>
<td>
<code>responseStartTimeoutSeconds</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>ResponseStartTimeoutSeconds is the maximum duration in seconds that the request
routing layer will wait for a request delivered to a container to begin
sending any network traffic.</p>
</td>
</tr>
<tr>
<td>
<code>idleTimeoutSeconds</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>IdleTimeoutSeconds is the maximum duration in seconds a request will be allowed
to stay open while not receiving any bytes from the user&rsquo;s application. If
unspecified, a system default will be provided.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="serving.knative.dev/v1.RevisionStatus">RevisionStatus
</h3>
<p>
(<em>Appears on:</em><a href="#serving.knative.dev/v1.Revision">Revision</a>)
</p>
<div>
<p>RevisionStatus communicates the observed state of the Revision (from the controller).</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>Status</code><br/>
<em>
<a href="#duck.knative.dev/v1.Status">
Status
</a>
</em>
</td>
<td>
<p>
(Members of <code>Status</code> are embedded into this type.)
</p>
</td>
</tr>
<tr>
<td>
<code>logUrl</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>LogURL specifies the generated logging url for this particular revision
based on the revision url template specified in the controller&rsquo;s config.</p>
</td>
</tr>
<tr>
<td>
<code>containerStatuses</code><br/>
<em>
<a href="#serving.knative.dev/v1.ContainerStatus">
[]ContainerStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>ContainerStatuses is a slice of images present in .Spec.Container[*].Image
to their respective digests and their container name.
The digests are resolved during the creation of Revision.
ContainerStatuses holds the container name and image digests
for both serving and non serving containers.
ref: <a href="http://bit.ly/image-digests">http://bit.ly/image-digests</a></p>
</td>
</tr>
<tr>
<td>
<code>initContainerStatuses</code><br/>
<em>
<a href="#serving.knative.dev/v1.ContainerStatus">
[]ContainerStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>InitContainerStatuses is a slice of images present in .Spec.InitContainer[*].Image
to their respective digests and their container name.
The digests are resolved during the creation of Revision.
ContainerStatuses holds the container name and image digests
for both serving and non serving containers.
ref: <a href="http://bit.ly/image-digests">http://bit.ly/image-digests</a></p>
</td>
</tr>
<tr>
<td>
<code>actualReplicas</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>ActualReplicas reflects the amount of ready pods running this revision.</p>
</td>
</tr>
<tr>
<td>
<code>desiredReplicas</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>DesiredReplicas reflects the desired amount of pods running this revision.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="serving.knative.dev/v1.RevisionTemplateSpec">RevisionTemplateSpec
</h3>
<p>
(<em>Appears on:</em><a href="#serving.knative.dev/v1.ConfigurationSpec">ConfigurationSpec</a>)
</p>
<div>
<p>RevisionTemplateSpec describes the data a revision should have when created from a template.
Based on: <a href="https://github.com/kubernetes/api/blob/e771f807/core/v1/types.go#L3179-L3190">https://github.com/kubernetes/api/blob/e771f807/core/v1/types.go#L3179-L3190</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
<em>(Optional)</em>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#serving.knative.dev/v1.RevisionSpec">
RevisionSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<br/>
<br/>
<table>
<tr>
<td>
<code>PodSpec</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#podspec-v1-core">
Kubernetes core/v1.PodSpec
</a>
</em>
</td>
<td>
<p>
(Members of <code>PodSpec</code> are embedded into this type.)
</p>
</td>
</tr>
<tr>
<td>
<code>containerConcurrency</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>ContainerConcurrency specifies the maximum allowed in-flight (concurrent)
requests per container of the Revision.  Defaults to <code>0</code> which means
concurrency to the application is not limited, and the system decides the
target concurrency for the autoscaler.</p>
</td>
</tr>
<tr>
<td>
<code>timeoutSeconds</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>TimeoutSeconds is the maximum duration in seconds that the request instance
is allowed to respond to a request. If unspecified, a system default will
be provided.</p>
</td>
</tr>
<tr>
<td>
<code>responseStartTimeoutSeconds</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>ResponseStartTimeoutSeconds is the maximum duration in seconds that the request
routing layer will wait for a request delivered to a container to begin
sending any network traffic.</p>
</td>
</tr>
<tr>
<td>
<code>idleTimeoutSeconds</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>IdleTimeoutSeconds is the maximum duration in seconds a request will be allowed
to stay open while not receiving any bytes from the user&rsquo;s application. If
unspecified, a system default will be provided.</p>
</td>
</tr>
</table>
</td>
</tr>
</tbody>
</table>
<h3 id="serving.knative.dev/v1.RouteSpec">RouteSpec
</h3>
<p>
(<em>Appears on:</em><a href="#serving.knative.dev/v1.Route">Route</a>, <a href="#serving.knative.dev/v1.ServiceSpec">ServiceSpec</a>)
</p>
<div>
<p>RouteSpec holds the desired state of the Route (from the client).</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>traffic</code><br/>
<em>
<a href="#serving.knative.dev/v1.TrafficTarget">
[]TrafficTarget
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Traffic specifies how to distribute traffic over a collection of
revisions and configurations.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="serving.knative.dev/v1.RouteStatus">RouteStatus
</h3>
<p>
(<em>Appears on:</em><a href="#serving.knative.dev/v1.Route">Route</a>)
</p>
<div>
<p>RouteStatus communicates the observed state of the Route (from the controller).</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>Status</code><br/>
<em>
<a href="#duck.knative.dev/v1.Status">
Status
</a>
</em>
</td>
<td>
<p>
(Members of <code>Status</code> are embedded into this type.)
</p>
</td>
</tr>
<tr>
<td>
<code>RouteStatusFields</code><br/>
<em>
<a href="#serving.knative.dev/v1.RouteStatusFields">
RouteStatusFields
</a>
</em>
</td>
<td>
<p>
(Members of <code>RouteStatusFields</code> are embedded into this type.)
</p>
</td>
</tr>
</tbody>
</table>
<h3 id="serving.knative.dev/v1.RouteStatusFields">RouteStatusFields
</h3>
<p>
(<em>Appears on:</em><a href="#serving.knative.dev/v1.RouteStatus">RouteStatus</a>, <a href="#serving.knative.dev/v1.ServiceStatus">ServiceStatus</a>)
</p>
<div>
<p>RouteStatusFields holds the fields of Route&rsquo;s status that
are not generally shared.  This is defined separately and inlined so that
other types can readily consume these fields via duck typing.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>url</code><br/>
<em>
knative.dev/serving/pkg/apis.URL
</em>
</td>
<td>
<em>(Optional)</em>
<p>URL holds the url that will distribute traffic over the provided traffic targets.
It generally has the form http[s]://{route-name}.{route-namespace}.{cluster-level-suffix}</p>
</td>
</tr>
<tr>
<td>
<code>address</code><br/>
<em>
<a href="#duck.knative.dev/v1.Addressable">
Addressable
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Address holds the information needed for a Route to be the target of an event.</p>
</td>
</tr>
<tr>
<td>
<code>traffic</code><br/>
<em>
<a href="#serving.knative.dev/v1.TrafficTarget">
[]TrafficTarget
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Traffic holds the configured traffic distribution.
These entries will always contain RevisionName references.
When ConfigurationName appears in the spec, this will hold the
LatestReadyRevisionName that we last observed.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="serving.knative.dev/v1.RoutingState">RoutingState
(<code>string</code> alias)</h3>
<div>
<p>RoutingState represents states of a revision with regards to serving a route.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;active&#34;</p></td>
<td><p>RoutingStateActive is a state for a revision which is actively referenced by a Route.</p>
</td>
</tr><tr><td><p>&#34;pending&#34;</p></td>
<td><p>RoutingStatePending is a state after a revision is created, but before
its routing state has been determined. It is treated like active for the purposes
of revision garbage collection.</p>
</td>
</tr><tr><td><p>&#34;reserve&#34;</p></td>
<td><p>RoutingStateReserve is a state for a revision which is no longer referenced by a Route,
and is scaled down, but may be rapidly pinned to a route to be made active again.</p>
</td>
</tr></tbody>
</table>
<h3 id="serving.knative.dev/v1.ServiceSpec">ServiceSpec
</h3>
<p>
(<em>Appears on:</em><a href="#serving.knative.dev/v1.Service">Service</a>)
</p>
<div>
<p>ServiceSpec represents the configuration for the Service object.
A Service&rsquo;s specification is the union of the specifications for a Route
and Configuration.  The Service restricts what can be expressed in these
fields, e.g. the Route must reference the provided Configuration;
however, these limitations also enable friendlier defaulting,
e.g. Route never needs a Configuration name, and may be defaulted to
the appropriate &ldquo;run latest&rdquo; spec.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ConfigurationSpec</code><br/>
<em>
<a href="#serving.knative.dev/v1.ConfigurationSpec">
ConfigurationSpec
</a>
</em>
</td>
<td>
<p>
(Members of <code>ConfigurationSpec</code> are embedded into this type.)
</p>
<p>ServiceSpec inlines an unrestricted ConfigurationSpec.</p>
</td>
</tr>
<tr>
<td>
<code>RouteSpec</code><br/>
<em>
<a href="#serving.knative.dev/v1.RouteSpec">
RouteSpec
</a>
</em>
</td>
<td>
<p>
(Members of <code>RouteSpec</code> are embedded into this type.)
</p>
<p>ServiceSpec inlines RouteSpec and restricts/defaults its fields
via webhook.  In particular, this spec can only reference this
Service&rsquo;s configuration and revisions (which also influences
defaults).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="serving.knative.dev/v1.ServiceStatus">ServiceStatus
</h3>
<p>
(<em>Appears on:</em><a href="#serving.knative.dev/v1.Service">Service</a>)
</p>
<div>
<p>ServiceStatus represents the Status stanza of the Service resource.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>Status</code><br/>
<em>
<a href="#duck.knative.dev/v1.Status">
Status
</a>
</em>
</td>
<td>
<p>
(Members of <code>Status</code> are embedded into this type.)
</p>
</td>
</tr>
<tr>
<td>
<code>ConfigurationStatusFields</code><br/>
<em>
<a href="#serving.knative.dev/v1.ConfigurationStatusFields">
ConfigurationStatusFields
</a>
</em>
</td>
<td>
<p>
(Members of <code>ConfigurationStatusFields</code> are embedded into this type.)
</p>
<p>In addition to inlining ConfigurationSpec, we also inline the fields
specific to ConfigurationStatus.</p>
</td>
</tr>
<tr>
<td>
<code>RouteStatusFields</code><br/>
<em>
<a href="#serving.knative.dev/v1.RouteStatusFields">
RouteStatusFields
</a>
</em>
</td>
<td>
<p>
(Members of <code>RouteStatusFields</code> are embedded into this type.)
</p>
<p>In addition to inlining RouteSpec, we also inline the fields
specific to RouteStatus.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="serving.knative.dev/v1.TrafficTarget">TrafficTarget
</h3>
<p>
(<em>Appears on:</em><a href="#serving.knative.dev/v1.RouteSpec">RouteSpec</a>, <a href="#serving.knative.dev/v1.RouteStatusFields">RouteStatusFields</a>)
</p>
<div>
<p>TrafficTarget holds a single entry of the routing table for a Route.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>tag</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Tag is optionally used to expose a dedicated url for referencing
this target exclusively.</p>
</td>
</tr>
<tr>
<td>
<code>revisionName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>RevisionName of a specific revision to which to send this portion of
traffic.  This is mutually exclusive with ConfigurationName.</p>
</td>
</tr>
<tr>
<td>
<code>configurationName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ConfigurationName of a configuration to whose latest revision we will send
this portion of traffic. When the &ldquo;status.latestReadyRevisionName&rdquo; of the
referenced configuration changes, we will automatically migrate traffic
from the prior &ldquo;latest ready&rdquo; revision to the new one.  This field is never
set in Route&rsquo;s status, only its spec.  This is mutually exclusive with
RevisionName.</p>
</td>
</tr>
<tr>
<td>
<code>latestRevision</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>LatestRevision may be optionally provided to indicate that the latest
ready Revision of the Configuration should be used for this traffic
target.  When provided LatestRevision must be true if RevisionName is
empty; it must be false when RevisionName is non-empty.</p>
</td>
</tr>
<tr>
<td>
<code>percent</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>Percent indicates that percentage based routing should be used and
the value indicates the percent of traffic that is be routed to this
Revision or Configuration. <code>0</code> (zero) mean no traffic, <code>100</code> means all
traffic.
When percentage based routing is being used the follow rules apply:
- the sum of all percent values must equal 100
- when not specified, the implied value for <code>percent</code> is zero for
that particular Revision or Configuration</p>
</td>
</tr>
<tr>
<td>
<code>url</code><br/>
<em>
knative.dev/serving/pkg/apis.URL
</em>
</td>
<td>
<em>(Optional)</em>
<p>URL displays the URL for accessing named traffic targets. URL is displayed in
status, and is disallowed on spec. URL must contain a scheme (e.g. http://) and
a hostname, but may not contain anything else (e.g. basic auth, url path, etc.)</p>
</td>
</tr>
</tbody>
</table>
<hr/>
<h2 id="serving.knative.dev/v1alpha1">serving.knative.dev/v1alpha1</h2>
<div>
<p>Package v1alpha1 contains the v1alpha1 versions of the serving apis.
Api versions allow the api contract for a resource to be changed while keeping
backward compatibility by support multiple concurrent versions
of the same resource</p>
</div>
Resource Types:
<ul></ul>
<h3 id="serving.knative.dev/v1alpha1.CannotConvertError">CannotConvertError
</h3>
<div>
<p>CannotConvertError is returned when a field cannot be converted.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>Message</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>Field</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<hr/>
<h2 id="serving.knative.dev/v1beta1">serving.knative.dev/v1beta1</h2>
<div>
<p>Package v1beta1 contains the v1beta1 versions of the serving apis.
Api versions allow the api contract for a resource to be changed while keeping
backward compatibility by support multiple concurrent versions
of the same resource</p>
</div>
Resource Types:
<ul><li>
<a href="#serving.knative.dev/v1beta1.DomainMapping">DomainMapping</a>
</li></ul>
<h3 id="serving.knative.dev/v1beta1.DomainMapping">DomainMapping
</h3>
<div>
<p>DomainMapping is a mapping from a custom hostname to an Addressable.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>
serving.knative.dev/v1beta1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>DomainMapping</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Standard object&rsquo;s metadata.
More info: <a href="https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata">https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata</a></p>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#serving.knative.dev/v1beta1.DomainMappingSpec">
DomainMappingSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Spec is the desired state of the DomainMapping.
More info: <a href="https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status">https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status</a></p>
<br/>
<br/>
<table>
<tr>
<td>
<code>ref</code><br/>
<em>
<a href="#duck.knative.dev/v1.KReference">
KReference
</a>
</em>
</td>
<td>
<p>Ref specifies the target of the Domain Mapping.</p>
<p>The object identified by the Ref must be an Addressable with a URL of the
form <code>{name}.{namespace}.{domain}</code> where <code>{domain}</code> is the cluster domain,
and <code>{name}</code> and <code>{namespace}</code> are the name and namespace of a Kubernetes
Service.</p>
<p>This contract is satisfied by Knative types such as Knative Services and
Knative Routes, and by Kubernetes Services.</p>
</td>
</tr>
<tr>
<td>
<code>tls</code><br/>
<em>
<a href="#serving.knative.dev/v1beta1.SecretTLS">
SecretTLS
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>TLS allows the DomainMapping to terminate TLS traffic with an existing secret.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#serving.knative.dev/v1beta1.DomainMappingStatus">
DomainMappingStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Status is the current state of the DomainMapping.
More info: <a href="https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status">https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status</a></p>
</td>
</tr>
</tbody>
</table>
<h3 id="serving.knative.dev/v1beta1.CannotConvertError">CannotConvertError
</h3>
<div>
<p>CannotConvertError is returned when a field cannot be converted.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>Message</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>Field</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="serving.knative.dev/v1beta1.DomainMappingSpec">DomainMappingSpec
</h3>
<p>
(<em>Appears on:</em><a href="#serving.knative.dev/v1beta1.DomainMapping">DomainMapping</a>)
</p>
<div>
<p>DomainMappingSpec describes the DomainMapping the user wishes to exist.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ref</code><br/>
<em>
<a href="#duck.knative.dev/v1.KReference">
KReference
</a>
</em>
</td>
<td>
<p>Ref specifies the target of the Domain Mapping.</p>
<p>The object identified by the Ref must be an Addressable with a URL of the
form <code>{name}.{namespace}.{domain}</code> where <code>{domain}</code> is the cluster domain,
and <code>{name}</code> and <code>{namespace}</code> are the name and namespace of a Kubernetes
Service.</p>
<p>This contract is satisfied by Knative types such as Knative Services and
Knative Routes, and by Kubernetes Services.</p>
</td>
</tr>
<tr>
<td>
<code>tls</code><br/>
<em>
<a href="#serving.knative.dev/v1beta1.SecretTLS">
SecretTLS
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>TLS allows the DomainMapping to terminate TLS traffic with an existing secret.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="serving.knative.dev/v1beta1.DomainMappingStatus">DomainMappingStatus
</h3>
<p>
(<em>Appears on:</em><a href="#serving.knative.dev/v1beta1.DomainMapping">DomainMapping</a>)
</p>
<div>
<p>DomainMappingStatus describes the current state of the DomainMapping.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>Status</code><br/>
<em>
<a href="#duck.knative.dev/v1.Status">
Status
</a>
</em>
</td>
<td>
<p>
(Members of <code>Status</code> are embedded into this type.)
</p>
</td>
</tr>
<tr>
<td>
<code>url</code><br/>
<em>
knative.dev/serving/pkg/apis.URL
</em>
</td>
<td>
<em>(Optional)</em>
<p>URL is the URL of this DomainMapping.</p>
</td>
</tr>
<tr>
<td>
<code>address</code><br/>
<em>
<a href="#duck.knative.dev/v1.Addressable">
Addressable
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Address holds the information needed for a DomainMapping to be the target of an event.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="serving.knative.dev/v1beta1.SecretTLS">SecretTLS
</h3>
<p>
(<em>Appears on:</em><a href="#serving.knative.dev/v1beta1.DomainMappingSpec">DomainMappingSpec</a>)
</p>
<div>
<p>SecretTLS wrapper for TLS SecretName.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>secretName</code><br/>
<em>
string
</em>
</td>
<td>
<p>SecretName is the name of the existing secret used to terminate TLS traffic.</p>
</td>
</tr>
</tbody>
</table>
<hr/>
<p><em>
Generated with <code>gen-crd-api-reference-docs</code>
.
</em></p>
