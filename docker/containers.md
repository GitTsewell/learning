### 新建并启动容器
```
docker run [OPTIONS] IMAGE [COMMAND] [ARG...]
```
options 有很多参数
+ -d 后台运行
+ -p 端口映射 有以下四种格式
    + ip:hostPort:containerPort 
    + ip::containerPort
    + hostPort:containerPort
    + containerPort
+ -network 指定网络模式，该选项有以下可选参数
    + –network=bridge： 默认选项，表示连接到默认的网桥。
    + –network=host：容器使用宿主机的网络。
    + –network=container:NAME_or_ID：告诉Docker让新建的容器使用已有容器的网络配置。
    + –network=none：不配置该容器的网络，用户可自定义网络配置。
+ -i: 以交互模式运行容器，通常与 -t 同时使用
+ -t: 为容器重新分配一个伪输入终端，通常与 -i 同时使用；列如
```
docker run -it nginx:latest /bin/bash
```
+ --name="nginx-lb": 为容器指定一个名称
+ -h "mars": 指定容器的hostname
+ --link=[]: 添加链接到另一个容器
+ --expose=[]: 开放一个端口或一组端口
+ --volume , -v: 绑定一个卷

### 列出容器
```
docker ps
```
+ CONTAINER_ID：表示容器ID。
+ IMAGE：表示镜像名称。
+ COMMAND：表示启动容器时运行的命令。
+ CREATED：表示容器的创建时间。
+ STATUS：表示容器运行的状态。Up表示运行中，Exited表示已停止。
+ PORTS：表示容器对外的端口号。
+ NAMES：表示容器名称。该名称默认由Docker自动生成，也可使用docker run命令的–name选项自行指定

### 停止容器
```
docker stop [OPTIONS] CONTAINER [CONTAINER...]
```
参数 --time -t 默认10秒 强制杀死容器前等待的时间，单位是秒  
在Docker大行其道的今天，我们能够非常方便的使用容器打包我们的应用程序，并且将它在我们的服务器上部署并运行起来。但是，谈论到如何停掉运行中的docker容器并正确的终止其中的程序，这就成为一个非常值得讨论的话题了。  
事实上，在我们日常的项目当中，这是我们经常需要面对和处理的问题：  
+ 场景A：假如我们打包在容器中的程序，提供HTTP方式的服务，负责处理各种HTTP requests并返回结果，我们必然希望在容器被停掉的时候，能够让程序有时间把已经在处理中的请求继续处理完毕，并返回结果给客户端。  
+ 场景B：又比如我们打包在容器中的程序，负责写入数据到某个数据文件中，我们希望程序能够在容器被停掉的时候，有时间把内存中缓存的数据持久化到存储设备中，以防数据丢失。  
+ 场景C：再比如现在流行的微服务架构中，一般会有服务发现的机制，也即每一个微服务在启动之后，都会主动把自己的地址信息注册到服务发现模块当中，让其他的服务可以知道自己的存在。而在容器被停掉的时候，微服务需要即时从服务发现模块中注销自己，以防止从API Gateway而来的请求被错误的路由到了已经被停止掉的微服务。  
当我们用docker stop命令来停掉容器的时候，docker默认会允许容器中的应用程序有10秒的时间用以终止运行  
在docker stop命令执行的时候，会先向容器中PID为1的进程发送系统信号SIGTERM，然后等待容器中的应用程序终止执行，如果等待时间达到设定的超时时间，或者默认的10秒，会继续发送SIGKILL的系统信号强行kill掉进程。在容器中的应用程序，可以选择忽略和不处理SIGTERM信号，不过一旦达到超时时间，程序就会被系统强行kill掉，因为SIGKILL信号是直接发往系统内核的，应用程序没有机会去处理它。在使用docker stop命令的时候，我们唯一能控制的是超时时间，比如设置为20秒超时  

### 强制停止容器
```
docker kill [OPTIONS] CONTAINER [CONTAINER...]
```
参数 --signal, -s	KILL	向容器发送一个信号  
与docker stop命令不一样的地方在于，docker kill没有任何的超时时间设置，它会直接发送SIGKILL信号，以及用户通过signal参数指定的其他信号。  
其实不难看出，docker stop命令，更类似于Linux系统中的kill命令，二者都是发送系统信号SIGTERM。而docker kill命令，更像是Linux系统中的kill -9或者是kill -SIGKILL命令，用来发送SIGKILL信号，强行终止进程

## 启动已停止的容器
```
docker start [OPTIONS] CONTAINER [CONTAINER...]
```

## 重启容器
可使用docker restart 命令来重启容器。该命令实际上是先执行了docker stop 命令，然后执行了docker start 命令
```
docker restart [OPTIONS] CONTAINER [CONTAINER...]
```
参数 --time, -t	10	关闭容器前等待的时间，单位是秒