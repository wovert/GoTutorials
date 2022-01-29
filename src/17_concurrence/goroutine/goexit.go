package main

import (
	"fmt"
	"runtime"
)

func testGoexit() {
	//defer fmt.Println("cccccccccccc")
	//return // 下面语句不会执行，返回当前函数出
	//fmt.Println("ddddddddddd")

	//fmt.Println("cccccccccccc")
	//return // 下面语句不会执行，返回当前函数出
	//defer fmt.Println("ddddddddddd")


	// 调用Goexit 之前的 defer 定义的语句会执行
	defer fmt.Println("cccccccccccc")

	// 终止所在的协程序（man()函数的go func()退出），不会执行下面 defer 语句
	runtime.Goexit()
	defer fmt.Println("ddddddddddd")
}

func main() {
	go func() {
		fmt.Println("aaaaaaaaaaa")
		
		testGoexit()

		fmt.Println("bbbbbbbbbbb")
	}()

	for {;}

}