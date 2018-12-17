##### PHP



##### `php`常用数组操作函数


* `array_fill_keys`:函数使用指定的键和值填充数组
```
$keys = array("a","b","c","d");
$res = array_fill_keys($keys,"blue")
print_r($res)

===========
Array ( [a] => blue [b] => blue [c] => blue [d] => blue )

```