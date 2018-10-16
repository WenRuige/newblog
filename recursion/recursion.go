package main

import "fmt"

// 阶乘计算(递归实现)
func factorial(num int) int {
	if num == 1 {
		return 1
	}
	return num * factorial(num-1)
}

// 阶乘计算(非递归实现)
func factorialV2(num int) int {
	sum := 1
	for i := 1; i <= num; i++ {
		sum = sum * i
	}
	return sum
}

// 斐波那契递归实现
func fibonacci(num int) int {
	if num < 2 {
		return 1
	}
	return fibonacci(num-1) + fibonacci(num-2)
}

// 斐波那契非递归实现
func fibonacciV2(num int) int {
	return 0
}
func main() {
	res := fibonacci(8)
	fmt.Println(res)
	//res := factorialV2(5)
	//fmt.Println(res)
}
