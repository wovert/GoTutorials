package main

import (
  "fmt"
  "time"
)

// 全局定义chanel, 用来完成数据同步
var ch = make(chan int)

func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data) // stdout
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}

func write() {
	Printer("hello")
	fmt.Println("\n开始写通道值")
	ch <- 888 // 写入
}

func read() {
	fmt.Println("\n开始读取通道值")
	channelValue := <-ch // 读取，读取数据之前一直阻塞
	fmt.Println("通道值：", channelValue)
	Printer("world")
}

func main() {

	// 新建2个协程 2个人同时使用打印机
	go write()
	go read()

	// 死循环，不让主协程结束
	for {
		;
	}
}