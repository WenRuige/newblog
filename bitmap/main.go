package main

import (
	"bytes"
	"fmt"
)

const size = 64

// unsafe sizeof uint64 8

func NewBitMap(nbits int) *BitSet {
	wordsLen := (nbits - 1) >> 6
	temp := BitSet(make([]uint64, wordsLen+1, wordsLen+1))
	return &temp
}

// BitSet
type BitSet []uint64

//查找单词索引
func (this *BitSet) wordIndex(bitIndex uint64) int {
	// bitIndex >> 6  64
	return int(bitIndex >> 6)
}

// 设置bitSet
func (this *BitSet) set(bitIndex uint64) {
	index := this.wordIndex(bitIndex)
	//  (bitIndex & 0x1F); 0x1F = 64
	//(*this)[index] = (*this)[index] | uint64(0x01) << (bitIndex %64)
	// 或操作将对应的index 里面的值置成1
	(*this)[index] |= uint64(0x01) << (bitIndex % 64)
}

// 展示当前位图
func (this *BitSet) display() {
	fmt.Println(len(*this))

}

// 清除数据
func (this *BitSet) clear(bitIndex uint64) {
	index := this.wordIndex(bitIndex)
	(*this)[index] ^= (uint64(0x01) << (bitIndex % 64))
}

func (this *BitSet) ToString() string {
	var temp uint64
	strAppend := &bytes.Buffer{}
	for i := 0; i < len(*this); i++ {
		temp = (*this)[i]
		for j := 0; j < 64; j++ {
			if temp&(uint64(0x01)<<uint64(j)) != 0 {
				strAppend.WriteString("1")
			} else {
				strAppend.WriteString("0")
			}
		}
	}
	return strAppend.String()
}

func main() {
	bit := NewBitMap(10)
	bit.set(10)
	bit.set(63)
	bit.set(1)
	//bit.clear(10)
	//fmt.Println(bit.ToString())

	fmt.Println(100%64, 100&64)
}
