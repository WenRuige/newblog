package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 3, 5, 2, 4, 6}
	res := topk1(nums, 3)
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
