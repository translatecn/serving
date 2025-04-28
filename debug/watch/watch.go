package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	certmanagerv1versioned "github.com/cert-manager/cert-manager/pkg/client/clientset/versioned"
	"github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	informerappv1 "k8s.io/client-go/informers/apps/v1"
	informerv1 "k8s.io/client-go/informers/core/v1"
	informernetv1 "k8s.io/client-go/informers/networking/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	networkingv1versioned "knative.dev/serving/networking/pkg/client/clientset/versioned"
	servingv1versioned "knative.dev/serving/pkg/client/clientset/versioned"

	"net/http"

	"k8s.io/client-go/rest"

	certmanagerv1informers "github.com/cert-manager/cert-manager/pkg/client/informers/externalversions/certmanager/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
	networkv1alpha1informer "knative.dev/serving/networking/pkg/client/informers/externalversions/networking/v1alpha1"
	autov1alpha1informer "knative.dev/serving/pkg/client/informers/externalversions/autoscaling/v1alpha1"
	servingv1informer "knative.dev/serving/pkg/client/informers/externalversions/serving/v1"
	servingv1beta1informer "knative.dev/serving/pkg/client/informers/externalversions/serving/v1beta1"
)

func main() {
	os.Remove("data")
	os.MkdirAll("data", 0777)
	cfg, err := clientcmd.BuildConfigFromFlags("", "/Users/acejilam/.kube/koord")
	if err != nil {
		panic(err)
	}
	cfg.WrapTransport = func(rt http.RoundTripper) http.RoundTripper {
		return &LoggingTransport{rt: rt}
	}

	WatchV1(cfg)
	WatchServingV1(cfg)
	WatchServingV1beta1(cfg)
	WatchAuto(cfg)
	WatchNetworkV1beta1(cfg)
	WatchCert(cfg)
	<-wait.NeverStop
}

func WatchV1(cfg *rest.Config) {
	client, err := kubernetes.NewForConfig(cfg)

	if err != nil {
		panic(err)
	}
	pi := informerv1.NewPodInformer(client, metav1.NamespaceAll, 0, cache.Indexers{})
	pi.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			write(nil, obj, "Pod-Create")
		},
		UpdateFunc: func(old, new interface{}) {
			write(old, new, "Pod")
		},
		DeleteFunc: func(obj interface{}) {
			write(obj, nil, "Pod-Delete")
		},
	})
	go pi.Run(wait.NeverStop)

	ep := informerv1.NewEndpointsInformer(client, metav1.NamespaceAll, 0, cache.Indexers{})
	ep.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			write(nil, obj, "Endpoints-Create")
		},
		UpdateFunc: func(old, new interface{}) {
			write(old, new, "Endpoints")
		},
		DeleteFunc: func(obj interface{}) {
			write(obj, nil, "Endpoints-Delete")
		},
	})
	go ep.Run(wait.NeverStop)

	ss := informerv1.NewServiceInformer(client, metav1.NamespaceAll, 0, cache.Indexers{})
	ss.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			write(nil, obj, "Service-Create")
		},
		UpdateFunc: func(old, new interface{}) {
			write(old, new, "Service")
		},
		DeleteFunc: func(obj interface{}) {
			write(obj, nil, "Service-Delete")
		},
	})
	go ss.Run(wait.NeverStop)

	ss2 := informerv1.NewSecretInformer(client, metav1.NamespaceAll, 0, cache.Indexers{})
	ss2.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			write(nil, obj, "Secret-Create")
		},
		UpdateFunc: func(old, new interface{}) {
			write(old, new, "Secret")
		},
		DeleteFunc: func(obj interface{}) {
			write(obj, nil, "Secret-Delete")
		},
	})
	go ss2.Run(wait.NeverStop)

	di := informerappv1.NewDeploymentInformer(client, metav1.NamespaceAll, 0, cache.Indexers{})
	di.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			write(nil, obj, "Deployment-Create")
		},
		UpdateFunc: func(old, new interface{}) {
			write(old, new, "Deployment")
		},
		DeleteFunc: func(obj interface{}) {
			write(obj, nil, "Deployment-Delete")
		},
	})
	go di.Run(wait.NeverStop)

	ki := informernetv1.NewIngressInformer(client, metav1.NamespaceAll, 0, cache.Indexers{})
	ki.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			write(nil, obj, "Ingress-Create")
		},
		UpdateFunc: func(old, new interface{}) {
			write(old, new, "Ingress")
		},
		DeleteFunc: func(obj interface{}) {
			write(obj, nil, "Ingress-Delete")
		},
	})
	go ki.Run(wait.NeverStop)
}

func WatchServingV1(cfg *rest.Config) { // 4
	client, err := servingv1versioned.NewForConfig(cfg)
	if err != nil {
		panic(err)
	}

	ri := servingv1informer.NewRevisionInformer(client, "", 0, nil)
	ri.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			write(nil, obj, "Revision-Create")
		},
		UpdateFunc: func(old, new interface{}) {
			write(old, new, "Revision")
		},
		DeleteFunc: func(obj interface{}) {
			write(obj, nil, "Revision-Delete")
		},
	})
	go ri.Run(wait.NeverStop)

	c := servingv1informer.NewConfigurationInformer(client, "", 0, nil)
	c.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			write(nil, obj, "Configuration-Create")
		},
		UpdateFunc: func(old, new interface{}) {
			write(old, new, "Configuration")
		},
		DeleteFunc: func(obj interface{}) {
			write(obj, nil, "Configuration-Delete")
		},
	})
	go c.Run(wait.NeverStop)

	r := servingv1informer.NewRouteInformer(client, "", 0, nil)
	r.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			write(nil, obj, "Route-Create")
		},
		UpdateFunc: func(old, new interface{}) {
			write(old, new, "Route")
		},
		DeleteFunc: func(obj interface{}) {
			write(obj, nil, "Route-Delete")
		},
	})
	go r.Run(wait.NeverStop)

	ks := servingv1informer.NewServiceInformer(client, "", 0, nil)
	ks.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			write(nil, obj, "KService-Create")
		},
		UpdateFunc: func(old, new interface{}) {
			write(old, new, "KService")
		},
		DeleteFunc: func(obj interface{}) {
			write(obj, nil, "KService-Delete")
		},
	})
	go ks.Run(wait.NeverStop)

}

func WatchServingV1beta1(cfg *rest.Config) { //1
	client, err := servingv1versioned.NewForConfig(cfg)
	if err != nil {
		panic(err)
	}

	ri := servingv1beta1informer.NewDomainMappingInformer(client, "", 0, nil)
	ri.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			write(nil, obj, "DomainMapping-Create")
		},
		UpdateFunc: func(old, new interface{}) {
			write(old, new, "DomainMapping")
		},
		DeleteFunc: func(obj interface{}) {
			write(obj, nil, "DomainMapping-Delete")
		},
	})
	go ri.Run(wait.NeverStop)

}

func WatchNetworkV1beta1(cfg *rest.Config) { //4
	client, err := networkingv1versioned.NewForConfig(cfg)
	if err != nil {
		panic(err)
	}

	ri := networkv1alpha1informer.NewCertificateInformer(client, "", 0, nil)
	ri.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			write(nil, obj, "Certificate-Create")
		},
		UpdateFunc: func(old, new interface{}) {
			write(old, new, "Certificate")
		},
		DeleteFunc: func(obj interface{}) {
			write(obj, nil, "Certificate-Delete")
		},
	})
	go ri.Run(wait.NeverStop)

	cd := networkv1alpha1informer.NewClusterDomainClaimInformer(client, 0, nil)
	cd.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			write(nil, obj, "ClusterDomainClaim-Create")
		},
		UpdateFunc: func(old, new interface{}) {
			write(old, new, "ClusterDomainClaim")
		},
		DeleteFunc: func(obj interface{}) {
			write(obj, nil, "ClusterDomainClaim-Delete")
		},
	})
	go cd.Run(wait.NeverStop)

	ki := networkv1alpha1informer.NewIngressInformer(client, "", 0, nil)
	ki.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			write(nil, obj, "Kingress-Create")
		},
		UpdateFunc: func(old, new interface{}) {
			write(old, new, "Kingress")
		},
		DeleteFunc: func(obj interface{}) {
			write(obj, nil, "Kingress-Delete")
		},
	})
	go ki.Run(wait.NeverStop)

	ss := networkv1alpha1informer.NewServerlessServiceInformer(client, "", 0, nil)
	ss.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			write(nil, obj, "ServerlessService-Create")
		},
		UpdateFunc: func(old, new interface{}) {
			write(old, new, "ServerlessService")
		},
		DeleteFunc: func(obj interface{}) {
			write(obj, nil, "ServerlessService-Delete")
		},
	})
	go ss.Run(wait.NeverStop)

}

func WatchAuto(cfg *rest.Config) { // 2
	client, err := servingv1versioned.NewForConfig(cfg)
	if err != nil {
		panic(err)
	}

	ri := autov1alpha1informer.NewMetricInformer(client, "", 0, nil)
	ri.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			write(nil, obj, "Metric-Create")
		},
		UpdateFunc: func(old, new interface{}) {
			write(old, new, "Metric")
		},
		DeleteFunc: func(obj interface{}) {
			write(obj, nil, "Metric-Delete")
		},
	})
	go ri.Run(wait.NeverStop)

	ri2 := autov1alpha1informer.NewPodAutoscalerInformer(client, "", 0, nil)
	ri2.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			write(nil, obj, "PodAutoscaler-Create")
		},
		UpdateFunc: func(old, new interface{}) {
			write(old, new, "PodAutoscaler")
		},
		DeleteFunc: func(obj interface{}) {
			write(obj, nil, "PodAutoscaler-Delete")
		},
	})
	go ri2.Run(wait.NeverStop)

}

func write(_old, _new interface{}, name string) {
	var old, new map[string]interface{}
	marshal, _ := json.Marshal(_old)
	json.Unmarshal(marshal, &old)

	marshal, _ = json.Marshal(_new)
	json.Unmarshal(marshal, &new)

	if v, ok := old["metadata"]; ok {
		if _v, ok2 := v.(map[string]interface{}); ok2 {
			delete(_v, "managedFields")
		}
	}
	if v, ok := new["metadata"]; ok {
		if _v, ok2 := v.(map[string]interface{}); ok2 {
			delete(_v, "managedFields")
		}
	}

	if v, ok := old["metadata"]; ok {
		if _v, ok2 := v.(map[string]interface{})["annotations"]; ok2 {
			if _v2, ok3 := _v.(map[string]interface{}); ok3 {
				delete(_v2, "kubectl.kubernetes.io/last-applied-configuration")
			}
		}
	}
	if v, ok := new["metadata"]; ok {
		if _v, ok2 := v.(map[string]interface{})["annotations"]; ok2 {
			if _v2, ok3 := _v.(map[string]interface{}); ok3 {
				delete(_v2, "kubectl.kubernetes.io/last-applied-configuration")
			}
		}
	}

	a, _ := json.MarshalIndent(old, "  ", "  ")
	b, _ := json.MarshalIndent(new, "  ", "  ")
	differ := gojsondiff.New()
	d, _ := differ.Compare(a, b)

	if d.Modified() {
		config := formatter.AsciiFormatterConfig{
			ShowArrayIndex: true,
			Coloring:       false,
		}

		formatter := formatter.NewAsciiFormatter(old, config)
		diffString, err := formatter.Format(d)
		if err != nil {
			klog.Errorf("error formatting %s: %v", name, err)
		}
		ioutil.WriteFile(fmt.Sprintf("data/%d-%s", time.Now().UnixNano(), name), []byte(diffString), 0644)
	}
	//_diff := cmp.Diff(old, new)
	//_diff = strings.Replace(_diff, strconv.Itoa(int(' ')), "", -1)
}

func WatchCert(cfg *rest.Config) {
	client, err := certmanagerv1versioned.NewForConfig(cfg)
	if err != nil {
		panic(err)
	}

	informer := certmanagerv1informers.NewCertificateInformer(client, "", 0, nil)
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			write(nil, obj, "CertManager-Certificate-Create")
		},
		UpdateFunc: func(old, new interface{}) {
			write(old, new, "CertManager-Certificate")
		},
		DeleteFunc: func(obj interface{}) {
			write(obj, nil, "CertManager-Certificate-Delete")
		},
	})
	go informer.Run(wait.NeverStop)

	informer2 := certmanagerv1informers.NewCertificateRequestInformer(client, "", 0, nil)
	informer2.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			write(nil, obj, "CertManager-CertificateRequest-Create")
		},
		UpdateFunc: func(old, new interface{}) {
			write(old, new, "CertManager-CertificateRequest")
		},
		DeleteFunc: func(obj interface{}) {
			write(obj, nil, "CertManager-CertificateRequest-Delete")
		},
	})
	go informer2.Run(wait.NeverStop)
}

type LoggingTransport struct {
	rt http.RoundTripper
}

func (l *LoggingTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	klog.Infoln(request.URL.String(), request.Method)
	return l.rt.RoundTrip(request)
}
