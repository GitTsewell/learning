* [Golang三色标记+混合写屏障GC模式全分析](https://github.com/aceld/golang/blob/main/5%E3%80%81Golang%E4%B8%89%E8%89%B2%E6%A0%87%E8%AE%B0%2B%E6%B7%B7%E5%90%88%E5%86%99%E5%B1%8F%E9%9A%9CGC%E6%A8%A1%E5%BC%8F%E5%85%A8%E5%88%86%E6%9E%90.md)

## 垃圾收集的多个阶段
1. 清理终止阶段；
  + 暂停程序，所有的处理器在这时会进入安全点（Safe point）；
  +  如果当前垃圾收集循环是强制触发的，我们还需要处理还未被清理的内存管理单元；
2. 标记阶段；
  + 将状态切换至 _GCmark、开启写屏障、用户程序协助（Mutator Assists）并将根对象入队；
  + 恢复执行程序，标记进程和用于协助的用户程序会开始并发标记内存中的对象，写屏障会将被覆盖的指针和新指针都标记成灰色，而所有新创建的对象都会被直接标记成黑色；
  + 开始扫描根对象，包括所有 Goroutine 的栈、全局对象以及不在堆中的运行时数据结构，扫描 Goroutine 栈期间会暂停当前处理器；
  + 依次处理灰色队列中的对象，将对象标记成黑色并将它们指向的对象标记成灰色；
  + 使用分布式的终止算法检查剩余的工作，发现标记阶段完成后进入标记终止阶段；
3. 标记终止阶段；
  + 暂停程序、将状态切换至 _GCmarktermination 并关闭辅助标记的用户程序；
  + 清理处理器上的线程缓存；
4. 清理阶段；
  + 将状态切换至 _GCoff 开始清理阶段，初始化清理状态并关闭写屏障；
  + 恢复用户程序，所有新创建的对象会标记成白色；
  + 后台并发清理所有的内存管理单元，当 Goroutine 申请新的内存管理单元时就会触发清理；

## 触发时间
1. runtime.sysmon 和 runtime.forcegchelper — 后台运行定时检查和垃圾收集；
2. runtime.GC — 用户程序手动触发垃圾收集；
3. runtime.mallocgc — 申请内存时根据堆大小触发垃圾收集；
