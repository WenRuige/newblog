package main

import "fmt"

func heapSort(arr []int) []int {
	arrLen := len(arr)
	// 建堆,得出的第一个元素是最大的
	buildMaxHeap(arr, arrLen)

	/***
	eg :
	原始数组为:
	{20,30,90,40,70,110,60,10,100,50,80}
	第一次建堆排序后:
	{110,100,90,40,80,20,60,10,30,50,70}。

	将第一个元素和最后一个元素交换,确保最后一个元素是最大的



	***/

	for i := arrLen - 1; i >= 0; i-- {
		swap(arr, 0, i)
		arrLen -= 1
		// 从父节点开始下沉
		heapify(arr, 0, arrLen)
	}
	return arr
}

func buildMaxHeap(arr []int, arrLen int) {
	for i := arrLen / 2; i >= 0; i-- {
		heapify(arr, i, arrLen)
	}
}

// 下沉
func heapify(arr []int, i, arrLen int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	// 如果左儿子 小于当前的点,将最大的点记为左儿子
	if left < arrLen && arr[left] > arr[largest] {
		largest = left
	}
	// 如果有右儿子,右儿子同理
	if right < arrLen && arr[right] > arr[largest] {
		largest = right
	}
	// 如果最大的不是当前节点
	if largest != i {
		swap(arr, i, largest)
		// 递归
		heapify(arr, largest, arrLen)
	}
}

// 交换
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
func main() {

	arr := []int{1, 7, 4, 2, 23, 4, 125, 4, 2}
	res := heapSort(arr)
	fmt.Println(res)
}
