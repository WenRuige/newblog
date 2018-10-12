#### `DFS` 深度优先搜索




> 什么是深度优先搜索




**leetcode .78**
```
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
```


对于这道题,可能第一想法是用递归来做
eg1:



但是也可以使用非递归的方式来实现
eg2:首先简述一些执行逻辑
第一轮 n = 1 , sets = [],[1]

第二轮 n = 2 , sets = [],[1],[2],[1,2]

第三轮 n = 3 , sets = [],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]
```
func subsets(nums []int) [][]int {
	empty := []int{}
	sets := [][]int{empty}
	for _, n := range nums {
		newSets := make([][]int, 0)
		for _, set := range sets {
			tmp := make([]int, len(set)+1)
			copy(tmp, append(set, n))
			newSets = append(newSets, tmp)
		}
		sets = append(sets, newSets...)
	}
	return sets
}
```