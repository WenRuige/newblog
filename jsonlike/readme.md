#### 几种常见的解析协议

`Json`是纯文本解析协议,优点是可读性高,使用简单方便,但是它解析非常费时,解析内存耗费高,以及数据量大等问题,
在移动场景对性能要求较高的情况下,选择json作为通信协议并不是最佳的方式



##### `msgpack`
`it's like json but fast and small`,它的数据格式与`json`类似,但是在存储多字节,数组,数字等做了优化


```
// 27Bytes

{"compact":true,"schema":0}


```
其中大括号,引号冒号 占了9个字节(来表示毫无意义的数据)

```
MsgPack表达:
82               A7 compact        C3      A6  schema       00

(2-element map)  (7-Byte string)   true    (6-Byte string)  0integer
```

使用`MessagePack Code Generator`代码生成工具来进行生成

```
//go:generate msgp -o=stuff.go -tests=false
type Person struct {
	Name       string `msg:"name"`
	Address    string `msg:"address"`
	Age        int    `msg:"age"`
	Hidden     string `msg:"-"`  // this field is ignored
	unexported bool              // this field is also ignored
}
```

```
有一些简单的参数
-o: - output file name ( default is {input}_gen.go )
-tests: 是否生成test文件和benchmarks文件

```




##### `protolbuf`
> what is protobuf




##### `grpc`
