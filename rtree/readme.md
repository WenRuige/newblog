#### `R-tree`

`R-Tree`是多级平衡树,它是`B`树在多维空间上的扩展,在`R`树中存放的数据并不是原始数据
而是这些数据的最小多边形(`Minimal-Bounding-Box, MBR`),空间对象`MBR`被包含于`R`树的叶节点中
从叶子节点开始用矩形(`rectangle`)将空间框起来,父节点的矩形需要包含子节点,节点越往上框住的空间就越大



`R-tree`数据结构
* `R-tree`是`n`叉树,`n`称为`R-tree`的扇
* 每个节点对应一个矩形









 [参考文档2](http://xiaobaoqiu.github.io/blog/2014/06/17/r-tree/)
 [参考文档](https://blog.csdn.net/v_JULY_v/article/details/6530142)
 [Rtree实现](https://godoc.org/github.com/dhconnelly/rtreego)