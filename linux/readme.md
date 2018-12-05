#### shell




##### 查看某个进程的进程id

```
ps -ef | grep nginx | grep -v grep | awk '{print $2}'
```



##### `pkill`&`kill`区别


