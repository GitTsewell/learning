## order by
### 全字段排序
mysql把结果集读取放在内存或者临时文件中,根据需要排序字段做快速排序.

在explain的extra中 可以看到using:filesort,表示需要排序

### 优化方案
思路就是在索引层解决排序的问题,不需要回表取得结果集之后再排序,比如:
```
select * from xx where `city` = "xx" order by `name` desc
```
这个sql语句可以把 city和name做成一个联合索引