package main

import (
	"fmt"
	"runtime"
)

func init() {
	fmt.Printf("Map %v\n", m)
	info = fmt.Sprintf("os %s , arch %s", runtime.GOOS, runtime.GOARCH)
}

var m = map[int]string{1: "A", 2: "B"}

var info string

func main() {

	fmt.Println(info)
}

//  -- 当前代码包中所有全局变量的初始化会在代码包初始化函数执行前完成
//  -- 所有代码包初始化函数都会在main函数执行前执行完毕 既init在main前面
