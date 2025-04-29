set -ex

#docker run --rm -it -v `pwd`:/data -w /data registry.cn-hangzhou.aliyuncs.com/ls-2018/mygo:v1.24.1 ./hack/update-codegen.sh

kubectl -n default exec -it title -c title -- curl -H 'Host: stock-service-example.default.127.0.0.1.sslip.io' kourier-internal.kourier-system

kubectl -n default exec -it title -c title -- curl -H 'Host: helloworld.knative.top' kourier-internal.kourier-system

#apt install apache2-utils -y
# ab -n 3000 -c 20 -H 'Host: stock-service-example.default.127.0.0.1.sslip.io' http://kourier-internal.kourier-system:80/

# hey -z 30s -c 50 -m GET -H 'Host: stock-service-example.default.127.0.0.1.sslip.io' http://kourier-internal.kourier-system:80/

#pod_name=`kubectl get pods -n kourier-system |grep -v NAME|awk -F ' ' '{print $1}'`
#istioctl proxy-config --proxy-admin-port=9901 all ${pod_name}.kourier-system
