## join
创建两个表t1和t2:比如
```
CREATE TABLE `t2` (
`id` int(11) NOTNULL,
`a` int(11) DEFAULTNULL,
`b` int(11) DEFAULTNULL,
PRIMARY KEY (`id`),
KEY `a` (`a`)
) ENGINE=InnoDB;
```

执行一个join语句,比如:
```
select * from t1 straight_join t2 on (t1.a=t2.a)
```
t1表示驱动表,t2表示被驱动表

### 举个例子
假设t1有N行数据,t2有M行数据
N=100,M=1000

### 被驱动表走索引
假设被驱动表走索引,那么每次被驱动表查找的复杂度是log2M,上面的sql语句总的时间复杂度是N+log2M*N


显然N对于扫描行数的影响更大,所以我们选择驱动表要选用小表

### 被驱动表没走索引
假设被驱动表没走索引,那么每次被驱动表查找的复杂度是M,上面的sql语句总的时间复杂度是N+N*M




