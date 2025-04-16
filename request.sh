kubectl -n default exec -it title -c title -- curl -H 'Host: helloworld-go.default.127.0.0.1.sslip.io' kourier-ingress.kourier-system

# kn quickstart kind --kubernetes-version=ccr.ccs.tencentyun.com/acejilam/node:v1.30.3 --install-serving=true
