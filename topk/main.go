package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 3, 5, 2, 4, 6}
	res := topk(nums, 3)
	fmt.Println(res)
}

// 方法1:首先进行排序增序排序,然后返回后k个数即可
func topk(nums []int, k int) []int {
	if len(nums) < k {
		return []int{}
	}
	sort.Ints(nums)
	return nums[len(nums)-k:]
}
