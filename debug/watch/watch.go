package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	cachingv1alpha1 "knative.dev/serving/caching/pkg/apis/caching/v1alpha1"
	networkingv1beta1 "knative.dev/serving/networking/pkg/apis/networking/v1alpha1"
	autoscalingv1alpha1 "knative.dev/serving/pkg/apis/autoscaling/v1alpha1"
	servingv1 "knative.dev/serving/pkg/apis/serving/v1"
	servingv1beta1 "knative.dev/serving/pkg/apis/serving/v1beta1"

	"net/http"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
	cachingv1beta1client "knative.dev/serving/caching/pkg/client/clientset/versioned/typed/caching/v1alpha1"
	networkingv1alpha1client "knative.dev/serving/networking/pkg/client/clientset/versioned/typed/networking/v1alpha1"
	autoscalingv1alpha1client "knative.dev/serving/pkg/client/clientset/versioned/typed/autoscaling/v1alpha1"
	servingv1client "knative.dev/serving/pkg/client/clientset/versioned/typed/serving/v1"
	servingv1beta1client "knative.dev/serving/pkg/client/clientset/versioned/typed/serving/v1beta1"
)

func main() {

	cfg, err := clientcmd.BuildConfigFromFlags("", "/Users/acejilam/.kube/koord")
	if err != nil {
		panic(err)
	}
	cfg.WrapTransport = func(rt http.RoundTripper) http.RoundTripper {
		return &LoggingTransport{rt: rt}
	}

	WatchServingV1(cfg)
	WatchServingV1beta1(cfg)
	WatchAuto(cfg)
	WatchNetworkV1beta1(cfg)
	c, _ := cachingv1beta1client.NewForConfig(cfg)
	watch, _ := c.Images("").Watch(context.Background(), metav1.ListOptions{})

	for {
		info := <-watch.ResultChan()
		autoscaler := info.Object.DeepCopyObject().(*cachingv1alpha1.Image)
		autoscaler.ObjectMeta.ManagedFields = nil
		info.Object = autoscaler
		out, err := yaml.Marshal(info)
		if err != nil {
			panic(err)
		}
		ioutil.WriteFile(fmt.Sprintf("data/%d-images", time.Now().UnixNano()), out, 0644)
	}
}

func WatchServingV1(cfg *rest.Config) {
	c, err := servingv1client.NewForConfig(cfg)
	if err != nil {
		panic(err)
	}
	go func() {
		watch, err := c.Configurations("").Watch(context.Background(), metav1.ListOptions{})
		if err != nil {
			panic(err)
		}
		for {
			info := <-watch.ResultChan()
			autoscaler := info.Object.DeepCopyObject().(*servingv1.Configuration)
			autoscaler.ObjectMeta.ManagedFields = nil
			info.Object = autoscaler
			out, err := yaml.Marshal(info)
			if err != nil {
				panic(err)
			}
			ioutil.WriteFile(fmt.Sprintf("data/%d-configurations", time.Now().UnixNano()), out, 0644)
		}
	}()
	go func() {
		watch, err := c.Revisions("").Watch(context.Background(), metav1.ListOptions{})
		if err != nil {
			panic(err)
		}
		for {
			info := <-watch.ResultChan()
			autoscaler := info.Object.DeepCopyObject().(*servingv1.Revision)
			autoscaler.ObjectMeta.ManagedFields = nil
			info.Object = autoscaler
			out, err := yaml.Marshal(info)
			if err != nil {
				panic(err)
			}
			ioutil.WriteFile(fmt.Sprintf("data/%d-revisions", time.Now().UnixNano()), out, 0644)
		}
	}()
	go func() {
		watch, err := c.Services("").Watch(context.Background(), metav1.ListOptions{})
		if err != nil {
			panic(err)
		}
		for {
			info := <-watch.ResultChan()
			autoscaler := info.Object.DeepCopyObject().(*servingv1.Service)
			autoscaler.ObjectMeta.ManagedFields = nil
			info.Object = autoscaler
			out, err := yaml.Marshal(info)
			if err != nil {
				panic(err)
			}
			ioutil.WriteFile(fmt.Sprintf("data/%d-services", time.Now().UnixNano()), out, 0644)
		}
	}()
	go func() {
		watch, err := c.Routes("").Watch(context.Background(), metav1.ListOptions{})
		if err != nil {
			panic(err)
		}
		for {
			info := <-watch.ResultChan()
			autoscaler := info.Object.DeepCopyObject().(*servingv1.Route)
			autoscaler.ObjectMeta.ManagedFields = nil
			info.Object = autoscaler
			out, err := yaml.Marshal(info)
			if err != nil {
				panic(err)
			}
			ioutil.WriteFile(fmt.Sprintf("data/%d-route", time.Now().UnixNano()), out, 0644)
		}
	}()
}
func WatchServingV1beta1(cfg *rest.Config) {
	c, err := servingv1beta1client.NewForConfig(cfg)
	if err != nil {
		panic(err)
	}
	go func() {
		watch, err := c.DomainMappings("").Watch(context.Background(), metav1.ListOptions{})
		if err != nil {
			panic(err)
		}
		for {
			info := <-watch.ResultChan()
			autoscaler := info.Object.DeepCopyObject().(*servingv1beta1.DomainMapping)
			autoscaler.ObjectMeta.ManagedFields = nil
			info.Object = autoscaler
			out, err := yaml.Marshal(info)
			if err != nil {
				panic(err)
			}
			ioutil.WriteFile(fmt.Sprintf("data/%d-domainmappings", time.Now().UnixNano()), out, 0644)
		}
	}()
}
func WatchNetworkV1beta1(cfg *rest.Config) {
	c, err := networkingv1alpha1client.NewForConfig(cfg)
	if err != nil {
		panic(err)
	}
	go func() {
		watch, err := c.Certificates("").Watch(context.Background(), metav1.ListOptions{})
		if err != nil {
			panic(err)
		}
		for {
			info := <-watch.ResultChan()
			autoscaler := info.Object.DeepCopyObject().(*networkingv1beta1.Certificate)
			autoscaler.ObjectMeta.ManagedFields = nil
			info.Object = autoscaler
			out, err := yaml.Marshal(info)
			if err != nil {
				panic(err)
			}
			ioutil.WriteFile(fmt.Sprintf("data/%d-certificates", time.Now().UnixNano()), out, 0644)
		}
	}()
	go func() {
		watch, err := c.Ingresses("").Watch(context.Background(), metav1.ListOptions{})
		if err != nil {
			panic(err)
		}
		for {
			info := <-watch.ResultChan()
			autoscaler := info.Object.DeepCopyObject().(*networkingv1beta1.Ingress)
			autoscaler.ObjectMeta.ManagedFields = nil
			info.Object = autoscaler
			out, err := yaml.Marshal(info)
			if err != nil {
				panic(err)
			}
			ioutil.WriteFile(fmt.Sprintf("data/%d-ingresses", time.Now().UnixNano()), out, 0644)
		}
	}()
	go func() {
		watch, err := c.ServerlessServices("").Watch(context.Background(), metav1.ListOptions{})
		if err != nil {
			panic(err)
		}
		for {
			info := <-watch.ResultChan()
			autoscaler := info.Object.DeepCopyObject().(*networkingv1beta1.ServerlessService)
			autoscaler.ObjectMeta.ManagedFields = nil
			info.Object = autoscaler
			out, err := yaml.Marshal(info)
			if err != nil {
				panic(err)
			}
			ioutil.WriteFile(fmt.Sprintf("data/%d-serverlessservices", time.Now().UnixNano()), out, 0644)
		}
	}()
	go func() {
		watch, err := c.ClusterDomainClaims().Watch(context.Background(), metav1.ListOptions{})
		if err != nil {
			panic(err)
		}
		for {
			info := <-watch.ResultChan()
			autoscaler := info.Object.DeepCopyObject().(*networkingv1beta1.ClusterDomainClaim)
			autoscaler.ObjectMeta.ManagedFields = nil
			info.Object = autoscaler
			out, err := yaml.Marshal(info)
			if err != nil {
				panic(err)
			}
			ioutil.WriteFile(fmt.Sprintf("data/%d-clusterdomainclaims", time.Now().UnixNano()), out, 0644)
		}
	}()
}

func WatchAuto(cfg *rest.Config) {
	ac, err := autoscalingv1alpha1client.NewForConfig(cfg)
	if err != nil {
		panic(err)
	}
	go func() {
		watch, err := ac.PodAutoscalers("").Watch(context.Background(), metav1.ListOptions{})
		if err != nil {
			panic(err)
		}
		for {
			info := <-watch.ResultChan()
			autoscaler := info.Object.DeepCopyObject().(*autoscalingv1alpha1.PodAutoscaler)
			autoscaler.ObjectMeta.ManagedFields = nil
			info.Object = autoscaler
			out, err := yaml.Marshal(info)
			if err != nil {
				panic(err)
			}
			ioutil.WriteFile(fmt.Sprintf("data/%d-podautoscalers", time.Now().UnixNano()), out, 0644)
		}
	}()

	go func() {
		watch, err := ac.Metrics("").Watch(context.Background(), metav1.ListOptions{})
		if err != nil {
			panic(err)
		}
		for {
			info := <-watch.ResultChan()
			autoscaler := info.Object.DeepCopyObject().(*autoscalingv1alpha1.Metric)
			autoscaler.ObjectMeta.ManagedFields = nil
			info.Object = autoscaler
			out, err := yaml.Marshal(info)
			if err != nil {
				panic(err)
			}
			ioutil.WriteFile(fmt.Sprintf("data/%d-metric", time.Now().UnixNano()), out, 0644)
		}
	}()
}

type LoggingTransport struct {
	rt http.RoundTripper
}

func (l *LoggingTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	klog.Infoln(request.URL.String(), request.Method)
	return l.rt.RoundTrip(request)
}
