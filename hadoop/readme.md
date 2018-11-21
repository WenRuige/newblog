#### hadoop


##### 安装

TBD



#### `hdfs`的一些命令

```
hdfs dfs -ls /    查看文件

hdfs dfs -mkdir  / test //创建test目录

hdfs dfs -copyFromLocal 原始文件 目标路径  // 将本地文件上传至hdfs
 
hdfs dfs -copyToLocal 目录文件 本机目录     // 将远程文件下载到本地

hdfs dfs -ls /test    //like linux ls

hdfs dfs -chmod  777 /test/文件 //变更权限
```





#### `yarn`


[参考资料](https://www.jianshu.com/p/4a65fd034871)