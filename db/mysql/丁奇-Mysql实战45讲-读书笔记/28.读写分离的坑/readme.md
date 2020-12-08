### 1. 过期读(也就是从库延迟)
下面是一些处理过期读的方案:
1. 强制读主库方案
2. sleep方案
3. 判断主从是否有延迟方案
4. 配合semi-sync方案
5. 等主库位点方案
6. 等GTID方案

#### 强制读主库方案
这个方案是用的很多的,但是问题是如果遇到了一些业务,所有的查询都不能是过期读的需求,比如金融业务,那么如果都是强制读主库的话,相当于放弃了主从的扩展性

#### sleep方案
主库更新或,在读取从库之前sleep(1),大多数情况下主备的延迟都在1秒以内,所以sleep一秒很大概率拿到最新的数据,不够这个方案看上去是有点滑稽,不是那个靠谱

#### 判断主备无延迟
1. 通过show_slave_status里面的second_behind_master参数 == 0 
2. Master_Log_File和Read_Master_Log_Pos,表示的是读到的主库的最新位点; Relay_Master_Log_File和Exec_Master_Log_Pos,表示的是备库执行的最新位点。
3. 对比GTID集合确保无延迟

#### semi-sync方案
也就是半同步方案:
1. 事务提交的时候,主库把binlog发给从库;
2. 从库收到binlog以后,发回给主库一个ack,表示收到了;
3. 主库收到这个ack以后,才能给客户端返回“事务完成”的确认。

semi-sync的缺点:
1. 一主多从的时候,在某些从库执行查询请求会存在过期读的现象;
2. 在持续延迟的情况下,可能出现过度等待的问题。

#### 等主库位点方案
1. trx1事务更新完成后,马上执行showmaster status得到当前主库执行到的File和Position;
2. 再从库执行select master_pos_wait(File, Position, 1);
3. 如果结果大于等于0表示从库已经执行过了,可以查询

#### GTID方案
执行select wait_for_executed_gtid_set(gtid_set, 1);返回值是0就可以查询

