#### 选择排序

选择排序:顾名思义,就是直接从待排序的数组里面选择一个最小(最大)的数字出来,然后进行交换,直至到最后

```
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
```