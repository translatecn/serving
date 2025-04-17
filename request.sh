#docker run --rm -it -v `pwd`:/data -w /data registry.cn-hangzhou.aliyuncs.com/ls-2018/mygo:v1.24.0 ./hack/update-codegen.sh

kubectl -n default exec -it title -c title -- curl -H 'Host: stock-service-example.default.127.0.0.1.sslip.io' kourier-ingress.kourier-system
#curl -H 'Host: stock-service-example.default.127.0.0.1.sslip.io' kourier-ingress.kourier-system
