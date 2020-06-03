### 搜索镜像
```
docker search [option] xxx
```

### 下载镜像
```
docker pull [OPTIONS] NAME[:TAG|@DIGEST]
```
参数
+ --all-tags, -a 下载所有标签的镜像
+ --disable-content-trus 忽略镜像校验  
还可以指定仓库和版本,列如
``` 
docker pull reg.itmuch.com/java:7
```

### 列出镜像
docker images 列出已下载的镜像,大致格式如下  
REPOSITORY|TAG|IMAGE ID|CREATED|SIZE  
----------|---|--------|-------|---
java      |     latest      |        861e95c114d6  |      4 weeks ago  |       643.1 MB  
hello-world    |     latest      |       c54a2cc56cbb   |     5 months ago  |      1.848  

含义如下  
+ REPOSITORY：镜像所属仓库名称。
+ TAG：镜像标签。默认是latest，表示最新。
+ IMAGE ID：镜像ID，表示镜像唯一标识。
+ CREATED：镜像创建时间。
+ SIZE：镜像大小。
```
docker images [OPTIONS] [REPOSITORY[:TAG]]
```

### 删除本地镜像
```
docker rmi [OPTIONS] IMAGE [IMAGE...]
```
+ -f 强制删除 (比如有容器在使用是不能直接删除的 走强制删除)

### 保存镜像 
把镜像保存为本地文件
```
docker save [OPTIONS] IMAGE [IMAGE...]
docker save busybox > busybox.tar
docker save --output busybox.tar busybox
```

### 加载镜像 
从本地文件中加载
```
docker load [OPTIONS]
docker load < busybox.tar.gz
docker load --input fedora.tar
```

### 构建镜像

```
docker build [OPTIONS] PATH | URL | -
```

常用的参数
+ --add-host  添加自定义从host到IP的映射，格式为（host:ip）
+ --file, -f  指定Dockerfile的名称，默认是‘PATH/Dockerfile’
+ --tag, -t   设置标签，格式：name:tag，tag可选

[更多参数](http://www.itmuch.com/docker/04-docker-command-images/)