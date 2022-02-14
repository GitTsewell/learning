* [goroutine 调度器详解](https://juejin.cn/post/6844903589400281096)
* [调度机制 抢占式调度](https://blog.51cto.com/u_15107299/3935086)


### 基于信号的抢占式调度
1. 程序启动时,注册一个SIGURG信号处理函数(doSigPreempt)
2. sysmon监控线程检测到执行时间过长的goroutine(超过10ms)会向相应的M发送SIGURG信号
3. GC回收时,会把正在运行的G标记为可抢占(preemptStop=true),然后调用preemptM函数触发抢占,runtime.preemptM 会调用 runtime.signalM 向线程发送信号 SIGURG；
4. 操作系统中断线程,执行预先注册的信号处理函数(doSigPreempt)
5. doSigPreempt 函数会修改当前M的寄存器,程序回到用户态时,修改当前G状态为 被抢占,让当前函数陷入休眠并让出线程,当前G回归全局可运行队列,调度器会选择其它的 Goroutine 继续执行

### 调度器启动
初始化函数执行过程中会将maxcount设置成10000,是golang能够创建的最大线程数,虽然可以创建10000,但是同时可运行线程由GOMAXPROCS变量控制

