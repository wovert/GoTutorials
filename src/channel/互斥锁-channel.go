package main

import (
	"fmt"
	"time"
)

var ch89 = make(chan int)

func printer8(str string) {
	for _, ch := range str {
		fmt.Printf("%c", ch)
		time.Sleep(time.Millisecond * 300)
	}
}

func person1() { // 先执行
	printer8("hello")
	ch89 <- 11
}

func person2() { // 后执行
	<- ch89 // 读取数据，没有数据会阻塞
	printer8("world") // 读取数据
}

func main() {
	go person1();
	go person2();
	for {
		;
	}
}