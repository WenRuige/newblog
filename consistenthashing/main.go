package main

import (
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
)

type UInt32Slice []uint32

func (s UInt32Slice) Len() int {
	return len(s)
}

func (s UInt32Slice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s UInt32Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type Hash func(data []byte) uint32

type Map struct {
	hash     Hash
	replicas int               // 复制因子
	keys     UInt32Slice       // 已排序的hash切片
	hashMap  map[uint32]string // 节点哈希和key的map
}

func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[uint32]string),
	}

	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

// 增加节点
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		// 结合复制因子计算所有虚拟节点的hash值，并存入m.keys中，同时在m.hashMap中保存哈希值和key的映射
		for i := 0; i < m.replicas; i++ {
			hash := m.hash([]byte(strconv.Itoa(i) + key)) //调用crc32做了一次hash
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}
	// 对所有虚拟节点的哈希值进行排序，方便之后进行二分查找(基于keys做索引)
	sort.Sort(m.keys)
}

// 是否为空
func (m *Map) IsEmpty() bool {
	return len(m.keys) == 0
}

// Get 方法根据给定的对象获取最靠近它的那个节点key
func (m *Map) Get(key string) string {
	if m.IsEmpty() {
		return ""
	}
	// 通过crc-hash化
	hash := m.hash([]byte(key))

	// 通过二分查找获取最优节点，第一个节点hash值大于对象hash值的就是最优节点
	idx := sort.Search(len(m.keys), func(i int) bool { return m.keys[i] >= hash })

	// 如果查找结果大于节点哈希数组的最大索引，表示此时该对象哈希值位于最后一个节点之后，那么放入第一个节点中
	if idx == len(m.keys) {
		idx = 0
	}

	return m.hashMap[m.keys[idx]]
}

func main() {
	m := New(10, nil)
	m.Add("112")
	m.Add("111")
	m.Add("99")
	res := m.Get("10")
	fmt.Println(res)
}
