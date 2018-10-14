


* 首先从官网下载`qgis`安装包
* 安装`python3`





##### `mac`如何安装`python3`



```
brew search python

brew install python3
```


报错:
```
Linking /usr/local/Cellar/python/3.7.0... Error: Permission denied @ dir_s_mkdir - /usr/local/Frameworks
```

发现`/usr/local/`下面没有`Framework`,创建文件夹并给出组用户即可
```
sudo mkdir /usr/local/Frameworks $ sudo chown $(whoami):admin /usr/local/Frameworks
```

```
python3 -v    //显示Python3即可
```

