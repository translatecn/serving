PodAutoscaler、ServerlessService、Service、Pod -> PodAutoscaler
    创建更新  ServerlessService、Metrics、Deployment

ServerlessServices、Private Endpoint、Public Service、 activator-service -> ServerlessService
    创建更新  Public、Private Service , 更新 PublicEndpoint

KServices、Configurations、Route -> KServices
    创建更新 Configurations、KServices、Route

Revision、PodAutoscaler Owner、Deployment Owner、Certificates -> Revision
    创建更新 Deployment、PodAutoscaler、创建Image



dynamicClient查询所有deployment : Unstructured, 并转换成  PodScalable
