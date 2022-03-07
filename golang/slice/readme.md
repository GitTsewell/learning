* [Go面试必考题目之slice篇](https://juejin.cn/post/6844903859257606158)


### copy
slice 的深浅拷贝,copy拷贝slice底层数组

### 切片截取
s = s[low : high : max] 切片的三个参数的切片截取的意义为 low 为截取的起始下标（含）， high 为窃取的结束下标（不含 high），
max 为切片保留的原切片的最大下标（不含 max）；即新切片从老切片的 low 下标元素开始，len = high - low, cap = max - low；high 和 max 一旦超出在老切片中越界，就会发生 runtime err，slice out of range。另外如果省略第三个参数的时候，第三个参数默认和第二个参数相同，即 len = cap。