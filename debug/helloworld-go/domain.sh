kubectl delete certificaterequests.cert-manager.io -A --all
kubectl delete certificates.cert-manager.io -A --all
kubectl delete certificates.networking.internal.knative.dev -A --all
kubectl delete certificatesigningrequests.certificates.k8s.io -A --all
kubectl delete domainmappings.serving.knative.dev -A --all
kubectl delete clusterdomainclaims.networking.internal.knative.dev -A --all
kubectl delete secret --all

openssl genrsa -out /tmp/knativetop-key.pem 4096
openssl req -subj "/CN=helloworld.knative.top" -sha256  -new -key /tmp/knativetop-key.pem -out /tmp/knativetop.csr

echo subjectAltName = DNS:helloworld.knative.top > /tmp/extfile.cnf

openssl x509 -req -days 3650 -sha256 -in /tmp/knativetop.csr -signkey /tmp/knativetop-key.pem -out /tmp/knativetop-cert.pem -extfile extfile.cnf
kubectl create secret tls secret-tls --key /tmp/knativetop-key.pem --cert /tmp/knativetop-cert.pem

kubectl apply -f - <<EOF
apiVersion: serving.knative.dev/v1beta1
kind: DomainMapping
metadata:
  name: helloworld.knative.top
  namespace: default
spec:
  ref:
    name: stock-service-example-v1
    kind: Service
    apiVersion: v1
  tls:
    secretName: secret-tls
EOF