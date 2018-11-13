package main

import "fmt"

func quickSort(values []int, left, right int) {
	if left < right {
		// 基准点为数组最左面的点
		tmp := values[left]
		i, j := left, right
		for {
			// 从右向左找到小于基准的数
			for values[j] >= tmp && i < j {
				j--
			}
			// 从左向右找到大于基准的数
			for values[i] <= tmp && i < j {
				i++
			}
			// 如果当前i>=j 停止当前搜索
			if i >= j {
				break
			}
			// 交换
			values[i], values[j] = values[j], values[i]
		}
		// 基准点与i点交换
		values[left] = values[i]
		values[i] = tmp

		quickSort(values, left, i-1)
		quickSort(values, i+1, right)
	}
}

func main() {
	arr := []int{1, 2, 12, 23, 4, 2, 34, 56, 7, 8, 6}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)

}
