package main



import (
"sync"
"container/list"
"encoding/json"
"errors"
"fmt"
)

//lru 最近最少使用算法
//核心思想是:如果数据最近被访问过,那么将来被访问的几率更高
//常见的方式是用一个链表保存数据
//每当缓存命中,则将数据移到链表头部
//当链表满了,将尾部数据丢弃
type cacheItem struct {
	Key string
	Val interface{}
}

type LRU struct {
	maxNum int
	curNum int
	mutex  sync.Mutex
	data   *list.List
}

func (l *LRU) add(key string, value interface{}) error {
	if e, _ := l.exist(key); e {
		return errors.New(key + "已存在")
	}
	//如果当前长度等于最大长度
	if l.maxNum == l.curNum {
		//清空
		l.clear()
	}
	//加锁
	l.mutex.Lock()
	//当前值+1
	l.curNum ++
	data, _ := json.Marshal(cacheItem{key, value})
	l.data.PushFront(data)
	l.mutex.Unlock()
	return nil
}

func (l *LRU) set(key string, value interface{}) error {
	e, item := l.exist(key)
	if !e {
		return l.add(key, value)
	}
	l.mutex.Lock()
	data, _ := json.Marshal(cacheItem{key, value})
	item.Value = data
	l.mutex.Unlock()
	return nil
}

//清理数据
func (l *LRU) clear() interface{} {
	l.mutex.Lock()
	l.curNum --;
	v := l.data.Remove(l.data.Back())
	l.mutex.Unlock()
	return v
}

//判断key是否存在
func (l *LRU) exist(key string) (bool, *list.Element) {
	var data cacheItem
	//循环链表
	for v := l.data.Front(); v != nil; v = v.Next() {
		json.Unmarshal(v.Value.([]byte), &data)
		if key == data.Key {
			return true, v
		}
	}
	return false, nil
}

//获取数据
func (l *LRU) get(key string) interface{} {
	e, item := l.exist(key)
	if !e {
		return nil
	}
	l.mutex.Lock()
	l.data.MoveToFront(item)
	l.mutex.Unlock()
	var data cacheItem
	json.Unmarshal(item.Value.([]byte), &data)
	return data.Val
}

//删除数据
func (l *LRU) del(key string) error {
	e, item := l.exist(key)
	if !e {
		return errors.New(key + "不存在")
	}
	l.mutex.Lock()
	l.curNum --
	l.data.Remove(item)
	l.mutex.Unlock()
	return nil

}

//返回长度

func (l *LRU) len() int {
	return l.curNum
}

//打印链表

func (l *LRU) print() {
	var data cacheItem
	for v := l.data.Front(); v != nil; v = v.Next() {
		json.Unmarshal(v.Value.([]byte), &data)
		fmt.Println("key:", data.Key, "value", data.Val)
	}
}

//创建一个新的LRU
func LRUNew(maxNum int) *LRU {
	return &LRU{
		maxNum: maxNum,
		curNum: 0,
		mutex:  sync.Mutex{},
		data:   list.New(),
	}
}

func main() {
	lru := LRUNew(5)
	lru.add("1111", 1111)
	lru.add("1211", 1111)
	lru.add("1311", 1111)
	lru.add("1411", 1111)
	lru.add("1511", 1111)
	lru.print()

	fmt.Println(lru.get("1211"));
	lru.print();
	fmt.Println("________________")
	//再次添加元素，如果超过最大数量，则删除链表尾部元素，将新元素添加到链表头
	lru.add("6666", 6666);
	lru.print();
}

