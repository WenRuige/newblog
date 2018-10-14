


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


安装完成结果发现`qgis`需要安装`python3.6`,搜了一下`brew search python3@`并不能制定版本号,那么有两种解决方案

* 从源码安装
* 通过更改`repo`来进行安装

```
cd /usr/local/Homebrew/Library/Taps/homebrew/homebrew-core/Formula

git log python.rb

```






##### 题外:`locate`


##### 防止使用`brew install xxx`,使得每次都`updating homebrew`

```
export HOMEBREW_NO_AUTO_UPDATE=true
```
