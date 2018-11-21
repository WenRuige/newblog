package main

import "fmt"

// 选择排序
func selectSort(nums []int) []int {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			fmt.Println(nums[i], nums[j])
			if nums[j] < nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	return nums
}

func main() {

	nums := []int{1, 3, 2, 444, 2, 2, 2, 3, 4, 1}
	res := selectSort(nums)
	fmt.Println(res)
}
