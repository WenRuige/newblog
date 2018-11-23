package main

import "fmt"

// 基于golang的快速排序算法
// 利用分治法,可将快速排序分为三步
// 1,在所有的数据集,选择一个元素作为基准
// 2,对于小于基准的元素都移到元素的左边,对于大于基准的元素都移到元素的右边,这个操作称为分区操作(partition),分区结束后,基准所处的位置就是最终排序后的位置
// 时间复杂度是o(n * logn),最坏是 o(n ^ 2)
// 快速排序详解:http://wiki.jikexueyuan.com/project/easy-learn-algorithm/fast-sort.html

func quickSort(values []int, left int, right int) {
	// 设置结束节点
	if left < right {

		// 设置分水岭,一般是最左面的元素
		temp := values[left]

		// 设置哨兵
		i, j := left, right
		for {
			// 从右向左找，找到第一个比分水岭小的数
			for values[j] >= temp && i < j {
				j--
			}

			// 从左向右找，找到第一个比分水岭大的数
			for values[i] <= temp && i < j {
				i++
			}

			// 如果哨兵相遇，则退出循环
			if i >= j {
				break
			}

			// 交换左右两侧的值
			values[i], values[j] = values[j], values[i]
		}

		// 将分水岭移到哨兵相遇点
		values[left] = values[i]
		values[i] = temp

		// 递归，左右两侧分别排序
		quickSort(values, left, i-1)
		quickSort(values, i+1, right)
	}
}

func QuickSort(values []int) {
	quickSort(values, 0, len(values)-1)
}

func main() {
	nums := []int{1, 23, 142, 123, 123, 2, 312, 41, 1}
	QuickSort(nums)
	fmt.Println(nums)
}
