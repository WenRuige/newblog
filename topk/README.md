#### top K 




基于排序算法来进行计算,使用`golang`来实现

```
// 方法1:首先进行排序增序排序,然后返回后k个数即可
func topk(nums []int, k int) []int {
	if len(nums) < k {
		return []int{}
	}
	sort.Ints(nums)
	return nums[len(nums)-k:]
}

```

部分排序,只查找出前`k`个最大值
方法2: