#### `hotkey` & `bigkey`


解决方案:

- 使用本地缓存(进程内缓存),需要考虑本地缓存与redis数据一致性
- 利用分片的算法特性,对key进行打散处理





##### 如何发现bigkey和hotkey

