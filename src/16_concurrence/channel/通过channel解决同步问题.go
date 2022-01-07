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

// Person1完之后执行 Person2
func Person1() {
	Printer("hello")
	ch <- 666 // 给管道写数据，发送
}

func Person2() {
	<-ch // 从管道读取数据，接受，如果通道没有数据就会阻塞
	Printer("world")
}

func main() {

	// 新建2个协程 2个人同时使用打印机
	go Person1()
	go Person2()

	// Person1()
	// Person2()

	// 死循环，不让主协程结束
	for {
		;
	}
}