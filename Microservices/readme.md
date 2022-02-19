## 微服务架构
### 服务注册和服务发现
使用 K8S 的 Service 和 DNS:

每个微服务 都在 K8S 中创建一个 Service ,名起名比如: user.xingren.host ,
然后,其他微服务只需要 配置好这个 K8s 中的 Service name 即可,
最后,只要这些微服务服务都在一个 k8S 集群中运行,便可省去注册中心与服务发现的这些微服务组件

! [K8S 在微服务架构下做服务注册中心的一种思路](https://blog.csdn.net/itguangit/article/details/109731971)

