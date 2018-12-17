#####Mysql












##### mysql支持任意ip远程连接
```

use mysql;

# %为所有ip都可以远程访问
update user set host = '%' where user = 'root';
# 推送设置到内存或重启服务器也行
FLUSH PRIVILEGES;

```