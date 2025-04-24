PodAutoscaler、ServerlessService、Service、Pod -> PodAutoscaler
    创建更新  ServerlessService、Metrics、Deployment

ServerlessServices、Private Endpoint、Public Service、 activator-service -> ServerlessService
    创建更新  Public、Private Service , 更新 PublicEndpoint

KServices、Configurations、Route -> KServices
    创建更新 Configurations、KServices、Route

Revision、PodAutoscaler Owner、Deployment Owner、Certificates -> Revision
    创建更新 Deployment、PodAutoscaler、创建Image
Configurations、Revision Owner -> Revision
    删除一些不活跃的 Revision

Namespace、Certificates Owner -> Namespace
    删除、创建 Certificates

Certificates、 CertManger Certificates Owner、Services -> Certificate
    创建、更新 CertManger Certificates、Certificates

Metric -> Metric
    // 每个revision 都会定时去统计范围内的所有pod 的请求数


Route、Configurations、Revision -> Route

DomainMapping、Certificate Owner、Ingress Owner -> DomainMapping
    创建 Certificate


dynamicClient查询所有deployment : Unstructured, 并转换成  PodScalable
