package main

import (
	"fmt"
	"runtime"
)

func main() {

	n := runtime.GOMAXPROCS(2) // 设置最大的CPU数量
	fmt.Println("默认CPU核数：", n) // 4


	m := runtime.GOMAXPROCS(1) // 设置最大的CPU数量
	fmt.Println("上一次CPU核数：", m) // 2

	for {
		go fmt.Print(0) // 子协程
		fmt.Print(1) // 主协程
	}
}