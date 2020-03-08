package main

import (
  "fmt"
  "time"
)

var ch = make(chan int)

func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}

func Person1() {
	Printer("hello")
	ch <- 888 // 写入
}

func Person2() {
	<-ch // 读取
	Printer("world")
}

func main() {

	// 新建2个协程 2个人同时使用打印机
	go Person1()
	go Person2()

	// 死循环，不让主协程结束
	for {
		;
	}
}