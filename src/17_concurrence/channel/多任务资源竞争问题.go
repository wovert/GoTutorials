package main

import (
  "fmt"
  "time"
)

func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}

func Person1() {
	Printer("hello")
}

func Person2() {
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