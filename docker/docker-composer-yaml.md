### build
配置构建时的选项，Compose会利用它自动构建镜像。build的值可以是一个路径
```
build: ./dir
```
也可以是一个对象,用于指定dockerfile和参数
```
build:
  context: ./dir
  dockerfile: Dockerfile-alternate
  args:
    buildno: 1
```

### command
覆盖容器启动后默认执行的命令
```
command: bundle exec thin -p 3000
```
也可以是一个list,类似于dockerfile中的cmd指令
```
command: [bundle, exec, thin, -p, 3000]
```

### environment
环境变量设置,可使用数组或字典两种方式
```
environment:
  RACK_ENV: development
  SHOW: 'true'
  SESSION_SECRET:

environment:
  - RACK_ENV=development
  - SHOW=true
  - SESSION_SECRET
```

### expose
暴露端口,只将端口暴露给连接的服务,而不暴露给宿主机
```
expose:
 - "3000"
 - "8000"
```

### image
指定镜像名称或镜像id,如果本地不存在该镜像,compose会尝试下载该镜像
```
image:java
```

### link
连接到其他服务的容器,可以指定服务名称和服务别名,也可以指定服务名称
```
web:
  links:
   - db
   - db:database
   - redis
```

### port
暴露端口信息，可使用HOST:CONTAINER 的格式，也可只指定容器端口（此时宿主机将会随机选择端口），类似于docker run -p 。  
需要注意的是，当使用HOST:CONTAINER 格式映射端口时，容器端口小于60将会得到错误的接口，因为yaml会把xx:yy 的数字解析为60进制。因此，建议使用字符串的形式
```
ports:
 - "3000"
 - "3000-3005"
 - "8000:8000"
 - "9090-9091:8080-8081"
 - "49100:22"
 - "127.0.0.1:8001:8001"
 - "127.0.0.1:5000-5010:5000-5010"
```

### volumes
卷挂载路径设置。可以设置宿主机路径 （HOST:CONTAINER） ，也可指定访问模式
``` 
volumes:
  # Just specify a path and let the Engine create a volume
  - /var/lib/mysql

  # Specify an absolute path mapping
  - /opt/data:/var/lib/mysql

  # Path on the host, relative to the Compose file
  - ./cache:/tmp/cache

  # User-relative path
  - ~/configs:/etc/configs/:ro

  # Named volume
  - datavolume:/var/lib/mysql
```