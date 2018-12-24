#####Mysql












##### mysql支持任意ip远程连接
```

use mysql;

# %为所有ip都可以远程访问
update user set host = '%' where user = 'root';
# 推送设置到内存或重启服务器也行
FLUSH PRIVILEGES;

```







##### `mysql`隔离级别
* 未提交读(`READ UNCOMMITTED`):一个事务还没提交时,它做的变更就能被别的事务看见
* 提交读(`READ COMMITTED`):一个事务提交之后,它做的变更才能被其他事务看见
* 可重复读(`REPEATABLE READ`):一个事务执行过程中,看到的数据,总是跟这个事务启动时看到的数据是一致的
* 串行化(`SERIALIZABLE`):对于同一行记录,“写”会加“写锁”,"读”会加“读锁”。当出现读写锁冲突的时候，后访问的事务必须等前一个事务执行完成，才能继续执行。