* [【Kubernetes系列】第1篇 架构及组件介绍](https://juejin.cn/post/6844903934964809736)
## 组件
#### 1. ETCD
服务发现,一致性(选主)

#### 2. kube-apiserver
提供集群管理的REST API接口，包括认证授权、数据校验以及集群状态变更等

提供同其他模块之间的数据交互(其他模块通过API Server查询或修改数据，只有API Server才直接操作etcd)

#### 3. kube-scheduler
kube-scheduler负责分配调度Pod到集群内的节点上，它监听kube-apiserver，查询还未分配Node的Pod，然后根据调度策略为这些Pod分配节点

#### 4. kube-controller-manager
kube-controller-manager是Kubernetes的大脑，通过kube-apiserver监控整个集群的状态，并确保集群处于预期的工作状态

#### 5. Kubelet
个Node节点上都运行一个kubelet守护进程，默认监听10250端口，接收并执行master发来的指令，管理Pod及Pod中的容器。每个kubelet进程会在API Server上注册节点自身信息，定期向master节点汇报节点的资源使用情况

#### 6. kube-proxy
每台机器上都运行一个kube-proxy服务，它监听API Server中service和Pod的变化情况，并通过userspace、iptables、ipvs等proxier来为服务配置负载均衡

## 简述Kubernetes如何实现集群管理

在集群管理方面，Kubernetes将集群中的机器划分为一个Master节点和一群工作节点Node。其中，在Master节点运行着集群管理相关的一组进程
kube-apiserver、kube-controller-manager和kube-scheduler，这些进程实现了整个集群的资源管理、Pod调度、弹性伸缩、安全控制、系统监控和纠错等管理能力，并且都是全自动完成的

### 资源管理
通过kube-apiserver中的metric-server获取集群的监控信息,metric-server实际是调用每个node节点的Kubelet获取资源使用情况

### pod调度
! [这应该是最全的K8s-Pod调度策略了](https://cloud.tencent.com/developer/article/1644857)

### 弹性伸缩
! [kubernetes 指南 -- 弹性伸缩](https://segmentfault.com/a/1190000021545907)
由service组件负责


! [60道常见的Kubernetes面试题总结](http://dockone.io/article/2434304)
