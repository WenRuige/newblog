#### top K 




#### 方法1:基于排序算法来进行计算,使用`golang`来实现

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


#### 方法2:部分排序,只查找出前`k`个最大值


```
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
```


#### 方法三:基于快排来找出前`k`个最大值

