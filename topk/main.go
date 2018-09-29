package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// 方法1:首先进行排序增序排序,然后返回后k个数即可
func topk(nums []int, k int) []int {
	if len(nums) < k {
		return []int{}
	}
	sort.Ints(nums)
	return nums[len(nums)-k:]
}

//方法2:基于冒泡来进行部分排序
func topk1(nums []int, k int) []int {
	newNums := []int{}
	for i := 0; i < len(nums)-1; i++ {
		for j := i; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j+1], nums[j] = nums[j], nums[j+1]
			}
		}
		if i < k {
			newNums = append(newNums, nums[len(nums)-1-i])
		} else {
			break
		}
	}
	return newNums
}

// 方法3:基于快排来进行topk排序

func topk2() {

}

// 方法4 堆排序

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	h := &IntHeap{1, 4125, 1222, 12, 3, 42}
	heap.Init(h)
	heap.Push(h, 1)
	fmt.Println(h.Len())
	// pop的不是实际上堆排序的结果
	fmt.Println(h.Pop())

}
