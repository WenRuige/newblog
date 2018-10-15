package main

import "fmt"

func strviolentcracking(str, partten string) int {

	i, j := 0, 0

	for i < len(str) && j < len(partten) {
		if str[i] == partten[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
	}

	if j == len(partten) {
		return i - j
	}
	return -1 // 没找到

}

func main() {

	res := strviolentcracking("eatnothing", "nothing")
	fmt.Println(res)
}
