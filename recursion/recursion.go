package main

import "fmt"

func factorial(num int) int {
	if num == 1 {
		return 1
	}
	return num * factorial(num-1)
}

func factorialV2(num int) int {
	sum := 1
	for i := 1; i <= num; i++ {
		sum = sum * i
	}
	return sum
}
func main() {

	res := factorialV2(5)
	fmt.Println(res)
}
