### ADD 复制文件
从src目录复制文件到容器的dest。其中src可以是Dockerfile所在目录的相对路径，也可以是一个URL，还可以是一个压缩包 
```
ADD <src>... <dest>
ADD microservice-discovery-eureka-0.0.1-SNAPSHOT.jar app.jar
``` 
注意事项:
+ src必须在构建的上下文内，不能使用例如：ADD ../somethine /something 这样的命令，因为docker build 命令首先会将上下文路径和其子目录发送到docker daemon
+ 如果src是一个URL，同时dest不以斜杠结尾，dest将会被视为文件，src对应内容文件将会被下载到dest
+ 如果src是一个URL，同时dest以斜杠结尾，dest将被视为目录，src对应内容将会被下载到dest目录
+ 如果src是一个目录，那么整个目录下的内容将会被拷贝，包括文件系统元数据
+ 如果文件是可识别的压缩包格式，则docker会自动解压

### COPY 复制文件
复制本地端的src到容器的dest。COPY指令和ADD指令类似，COPY不支持URL和压缩包
```
COPY <src>... <dest>
```

### ARG 设置构建参数
ARG指令用于设置构建参数，类似于ENV。和ARG不同的是，ARG设置的是构建时的环境变量，在容器运行时是不会存在这些变量的
```
ARG <name>[=<default value>]
ARG user1=someuser
```

### ENV 设置环境变量
ENV指令用于设置环境变量
```
ENV <key> <value>
比如:
ENV GOPROXY https://goproxy.cn
ENV GO111MODULE on
```

### CMD 容器启动命令
CMD指令用于为执行容器提供默认值。每个Dockerfile只有一个CMD命令，如果指定了多个CMD命令，那么只有最后一条会被执行，如果启动容器的时候指定了运行的命令，则会覆盖掉CMD指定的命令
```
CMD ["executable","param1","param2"] (推荐使用)
CMD ["param1","param2"] (为ENTRYPOINT指令提供预设参数)
CMD command param1 param2 (在shell中执行)
比如:
CMD echo "This is a test." | wc -
```

### ENTRYPOINT 入口点
```
ENTRYPOINT ["executable", "param1", "param2"]
ENTRYPOINT command param1 param2
比如:
ENTRYPOINT  ["./main","-c","dev"]
```
ENTRYPOINT和CMD指令的目的一样，都是指定Docker容器启动时执行的命令，可多次设置，但只有最后一个有效。ENTRYPOINT指令给出容器启动后的默认行为，一般难以在启动容器时覆盖，但是可以追加命令参数  
列如:
```
docker run --entrypoint /bin/bash mysql:latest ...
```

### EXPOSE 声明暴露的端口
EXPOSE指令用于声明在运行时容器提供服务的端口
```
EXPOSE <port> [<port>...]
# 声明暴露一个端口示例
EXPOSE port1
# 声明暴露多个端口示例
EXPOSE port1 port2 port3
```
需要注意的是，这只是一个声明，运行时并不会因为该声明就打开相应端口。该指令的作用主要是帮助镜像使用者理解该镜像服务的守护端口；其次是当运行时使用随机映射时，会自动映射EXPOSE的端口


### FROM 指定基础镜像
使用FROM指令指定基础镜像，FROM指令有点像Java里面的“extends”关键字。需要注意的是，FROM指令必须指定且需要写在其他指令之前。FROM指令后的所有指令都依赖于该指令所指定的镜像。
```
FROM <image>
FROM <image>:<tag>
FROM <image>@<digest>  digest是对manifest文件的sha256
```

### LABEL 为镜像添加元数据
```
LABEL <key>=<value> <key>=<value> <key>=<value> ...
比如:
LABEL "com.example.vendor"="ACME Incorporated"
LABEL com.example.label-with-value="foo"
LABEL version="1.0"
```





