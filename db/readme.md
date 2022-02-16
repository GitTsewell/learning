* [如何保障mysql和redis之间的数据一致性](https://juejin.cn/post/6844904073783689224)
* [再有人问你分布式事务，把这篇扔给他](https://juejin.cn/post/6844903647197806605#heading-15)
* [CAP理论和Base理论](https://zhuanlan.zhihu.com/p/335617791)
* [MySQL对分布式事务（XA Transactions）的支持](http://www.asktheway.org/2020/04/26/266/)

### 2pac XA 一句话总结
+ 优点: 强一致性,实现简单
+ 缺点: 1.协调者单点故障  2.阻塞

### TCC (try,confirm,cancel) 一句话总结
+ 优点: 每一个阶段都会提交事务,不会阻塞,性能好
+ 缺点: 1.大量侵入式代码 2.业务代码分为三部分,增加复杂度  3.如果confirm失败,需要retry,还需要考虑幂等性

* [用Go轻松完成一个TCC分布式事务，保姆级教程](https://segmentfault.com/a/1190000040331793)
* [golang 分布式框架DTM](https://github.com/dtm-labs/dtm)
