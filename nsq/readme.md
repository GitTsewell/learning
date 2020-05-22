nsq有三部分组成,nsqlookupd+nsqadmin+nsqd  

## nsqlookupd
nsqlookupd是守护进程负责管理拓扑信息.客户端通过查询nsqlookupd来发现指定话题(topic)的生产者,然后nsqd节点广播话题(topic)和通道(channels)信息  
简单的说nsqlokupd就是中心管理服务,它默认使用tcp:4160端口管理nsqd服务,默认使用http:4161端口管理nsqadmin服务,同时为客户端提供查询功能
nsqlookupd的功能和特点:
+ 唯一性,一个nsq服务中只有一个nsqlookupd
+ 去中心和,及时nsqlookupd崩溃,也不会影响正在运行的nsqd服务
+ 充当nsqd和nsqadmin服务信息交互的中间件
+ 提供一个http查询服务,给客户端定时更新nsqd的地址服务

## nsqadmin
nsqadmin就是一套web ui,用来实时统计,并且执行不同的管理任务,nsqadmin默认访问地址是http://127.0.0.1:4171/
+ 查看所有topic和channel的实时监控
+ 展示所有message数量
+ 能够在后台创建channel和topic

## nsqd
nsqd是一个守护进程,负责接受,排队,投递消息给客户端.简单的说,真正干活的就是nsqd,默认监听tcp:4150端口和http:4151端口和一个可选的https端口
+ 对于订阅了同一个topic,同一个channel的消费者使用负载均衡策略
+ 只要channel存在,即使没有channel的消费者,也会将生产者的message缓存到队列中
+ 保证队列中的message至少会被消费一次,如果nsqd退出,会将队列中的消息暂存在磁盘上
+ 限定内存使用,能够配置nsqd中channel的message数量,如果过多超出,将message缓存到磁盘中
+ topic和channel一旦建立,将会一直存在,要及时在管理台或者用代理清楚无效的topic和channel

## 消费者
消费者有两种模式与nsqd建立连接:
+ 消费者直连nsqd,这是最简单的,缺点是nsqd服务无法实现动态伸缩
+ 消费者通过http查询nsqlookupd获取上面所有的nsqd连接地址,然后再分别和这些nsqd建立连接,这也是官方推荐的做法,但是客户端会不停的向nsqlookupd查询最新的nsqd地址目录

