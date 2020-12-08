## mysql主备一致性
主备流程图
![master-to-slave-flow](masterToSlaveFlow.png)

主从启动后,从库会启动两个线程,一个io_thread,一个sql_thread,io_thread负责与主库建立连接,写入本地