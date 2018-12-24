package main

import "fmt"

func main() {
	//var ch chan int = make(chan int, 4)
	//for i := 0; i < cap(ch); i++ {
	//	ch <- i
	//}
	//
	//for len(ch) > 0 {
	//	var value int = <-ch
	//	fmt.Println(value)
	//}

	var ch = make(chan int, 4)
	ch <- 1
	ch <- 2
	close(ch)

	for value := range ch {
		fmt.Println(value)
	}

}
