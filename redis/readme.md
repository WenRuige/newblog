#### `redis.md`





##### redis 命令
`Redis`几种命令类型`key`,`String`,`Hash`,`List`,`Set`,`Sorted Set`










##### `Redis key`回收策略
* `noviction`:返回错误当内存限制达到并且客户端尝试执行会让更多内存被使用的命令,
* `allkeys-lru`:尝试回收最少使用的键(`LRU`),使得新添加的数据有空间存放
* `volatile-lru`:尝试回收最少使用的键(`LRU`),但是仅仅限于过期时间的键,使得新添加的数据有空间存放
* `allkeys-random`: 回收随机的键
* `volatile-random`:回收仅在过期集合内随机的键
* `volatile-ttl`:回收在过期集合的键,并且优先回收存活时间(`TTL`)较短的键,使得新添加的数据有空间存放




##### 回收进程如何工作
* 一个客户端运行了新的命令,添加了新的数据
* `Redis`检查内存使用情况,如果大于`maxmemory`的限制,则根据设定好的策略进行回收
*  一个新的命令被设置&执行




