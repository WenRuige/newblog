package main

//import "fmt"
//
//type Trie struct {
//	Root *Node
//}
//
//type Node struct {
//	Char    rune
//	HasData bool
//	Data    interface{} //存放得数据
//	Child   map[rune]*Node
//}
//
////构建一颗树
//func NewTrie() *Trie {
//	return &Trie{
//		Root: NewNode(' '),
//	}
//}
//
////构建一个节点
//func NewNode(char rune) *Node {
//	node := Node{
//		Char:    char,
//		HasData: false,
//		Data:    "",
//		Child:   make(map[rune]*Node),
//	}
//	return &node
//}
//
////添加
//func (t *Trie) AddNode(key string, data interface{}) {
//	parent := t.Root
//	keyRune := []rune(key)
//	for _, v := range keyRune {
//		//判断是否存在此节点
//		node, ok := parent.Child[v]
//		if !ok {
//			//如果不存在则构建节点
//			node := &Node{
//				Char: v,
//			}
//			//加入到父节点得子节点
//			parent.Child[v] = node
//		}
//		//将父节点改成当前找到（或者添加）得节点
//		parent = node
//	}
//	//给最后得节点添加数据
//	parent.Data = data
//	parent.HasData = true
//}
//
////查找数据
//func (t *Trie) SearchNode(key string, limit int) (res []interface{}) {
//	keyRune := []rune(key)
//	parent := t.Root
//	//查找到key所能找到得节点
//	for _, v := range keyRune {
//		//判断是否存在此节点
//		node, ok := parent.Child[v]
//		if !ok {
//			//不存再此节点，则不存在
//			return
//		}
//		//将父节点改成当前找到得节点
//		parent = node
//	}
//	//将获取到得节点与其后代节点全部找出来
//	var queue []*Node //定义一个当前找到得节点或者子节点得切片
//
//	queue = append(queue, parent)
//	for len(queue) > 0 {
//		//定义一个存放子节点得切片
//		var childQueue []*Node
//		for _, v := range queue {
//			//如果该节点下存在数据，则将数据加入到返回得结果中
//			if v.HasData {
//				res = append(res, v.Data)
//			}
//			//如果数据达到了限定数量，则直接返回
//			if len(res) == limit {
//				return
//			}
//			//将下属得子节点加入到子节点切片中
//			for _, vi := range v.Child {
//				childQueue = append(childQueue, vi)
//			}
//		}
//		//将子节点切片赋给queue，继续子节点得数据查找
//		queue = childQueue
//	}
//	return
//}
//
//func main() {
//	obj := NewTrie()
//	obj.AddNode("h", 12)
//	obj.AddNode("ha", 1)
//
//	res := obj.SearchNode("h", 1)
//	fmt.Println(res)
//}
