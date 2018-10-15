#### `mac` 安装`qgis`的两种方法


> 什么是`qgis`

`TBD`





* 首先从官网下载`qgis`安装包
* `qgis`需要安装python版本号为`python3.6.x`



##### 介绍`mac`如何安装`python3`依赖

```
brew search python

brew install python3
```


执行上述两行代码发现报错:
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


安装完成结果发现`qgis`需要安装`python3.6`,搜了一下`brew search python3@`并不能制定版本号,那么有两种解决方案

* 从源码安装
* 通过更改`repo`来进行安装

```
cd /usr/local/Homebrew/Library/Taps/homebrew/homebrew-core/Formula

git log python.rb

```
**此处需要关闭homebrew的自动更新功能,否则会把分支更新到master上**

之后再次执行`brew install python3`即可


然后我发现即时`python3 -v` 显示的是`3.6.x`但是还是无法安装,遂搜索了一下
```
If you have installed Python3 with Homebrew, 
you can use symlinks to comply with the installer needs (as suggested by @shongololo).
It seems, the installer needs Python3 installed at (thanks @shongololo): /Library/Frameworks/Python.framework/Versions/3.6/bin/python3.
```

[资料](https://gis.stackexchange.com/questions/274381/installing-qgis3-on-mac)

所以是因为安装器需要将`python3`安装到制定的目录,那么我们做一下软链即可eg:

```
sudo ln -s /usr/local/Cellar/python/3.6.5_1/Frameworks/Python.framework /Library/Frameworks/Python.framework

```

然后发现成功安装了,但是打开会提示某些模块不存在,我们可以直接使用`pip`来进行安装

```
pip3 install owslib PyYaml psycopg2 jinja2 pygments plotly
```




##### 另外一种安装方式

```
brew tap osgeo/osgeo4mac
brew install osgeo/osgeo4mac/qgis3
```

`ps`这种方式我没有安装成功,安装后闪退

##### 题外:安装中一些工具的学习

`locate`

TBD


##### 防止使用`brew install xxx`,使得每次都`updating homebrew`

```
export HOMEBREW_NO_AUTO_UPDATE=true
```



#### 其他参考链接

[Installing QGIS on a Macbook](http://www.geothread.net/installing-qgis-on-a-macbook/)