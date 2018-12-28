#### shell




##### 查看某个进程的进程id

```
ps -ef | grep nginx | grep -v grep | awk '{print $2}'
```



##### `pkill`&`kill`区别






#### 想grep某个关键字并且看它的上下文
```
cat all.log | greep -C 50 "11"
```

