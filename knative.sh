#mkdir -p out
#cd out
#curl -L https://istio.io/downloadIstio | ISTIO_VERSION=1.25.0 TARGET_ARCH=arm64 sh -
#cd istio-1.25.0
#export PATH=$PWD/bin:$PATH
#istioctl install --set profile=default -y --set hub=registry.cn-hangzhou.aliyuncs.com/acejilam
install-k8s-by-kind.sh koord v1.30.3

function t() {
	docker pull $1
	kind load docker-image -n koord $1
}

t registry.cn-hangzhou.aliyuncs.com/acejilam/knative-serving-activator:v1.17.0
t registry.cn-hangzhou.aliyuncs.com/acejilam/knative-serving-autoscaler:v1.17.0
t registry.cn-hangzhou.aliyuncs.com/acejilam/knative-serving-controller:v1.17.0
t registry.cn-hangzhou.aliyuncs.com/acejilam/knative-serving-webhook:v1.17.0
t registry.cn-hangzhou.aliyuncs.com/acejilam/knative-serving-queue:v1.17.0
t registry.cn-hangzhou.aliyuncs.com/acejilam/net-kourier-kourier:v1.17.0
t registry.cn-hangzhou.aliyuncs.com/acejilam/centos:7
t registry.cn-hangzhou.aliyuncs.com/acejilam/envoyproxy-envoy:v1.31-latest
t registry.cn-hangzhou.aliyuncs.com/ls-2018/mygo:v1.24.0

#cd ../..
kubectl create ns knative-serving
kubectl apply -f ./debug/yaml/serving-crds.yaml
kubectl wait --for=condition=Established --all crd

kubectl apply -f ./debug/yaml/serving-core.yaml

kubectl wait -A --for=condition=Ready --all pod --timeout=3000s
kubectl apply -f ./debug/yaml/kourier.yaml
kubectl wait -A --for=condition=Ready --all pod --timeout=3000s

kubectl patch configmap/config-network \
	--namespace knative-serving \
	--type merge \
	--patch '{"data":{"ingress-class":"kourier.ingress.networking.knative.dev"}}'

kubectl apply -f ./debug/yaml/kourier-ingress.yaml

kubectl patch configmap/config-domain \
	--namespace knative-serving \
	--type merge \
	--patch '{"data":{"127.0.0.1.sslip.io":""}}'

#kubectl apply -f ./yaml
#export https_proxy=http://127.0.0.1:7890 http_proxy=http://127.0.0.1:7890 all_proxy=socks5://127.0.0.1:7890
#kubectl apply --filename https://github.com/knative/serving/releases/download/knative-v1.17.0/serving-core.yaml
#kubectl apply --filename https://github.com/knative/serving/releases/download/knative-v1.17.0/serving-crds.yaml
#kubectl apply --filename https://github.com/knative-extensions/net-istio/releases/download/knative-v1.17.0/net-istio.yaml

#kubectl label namespace default istio-injection=enabled
kubectl get pods --namespace knative-serving

#git clone https://github.com/knative/docs.git
#cd docs/code-samples/serving/hello-world/helloworld-go
cd helloworld-go
#docker build -t registry.cn-hangzhou.aliyuncs.com/ls-2018/knative:helloworld-go .
#docker push registry.cn-hangzhou.aliyuncs.com/ls-2018/knative:helloworld-go
#sed -i 's@docker.io/{username}/helloworld-go@registry.cn-hangzhou.aliyuncs.com/ls-2018/knative:helloworld-go@g' service.yaml
kubectl apply -f pod.yaml
#cd -
# k exec -it title -c title -- curl -H 'Host: helloworld-go.default.127.0.0.1.sslip.io' kourier.kourier-system
