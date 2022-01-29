package main

import "fmt"

func fibonacci(ch <-chan int, quit <-chan bool) {
	for {
		select {
			case num := <-ch:
				fmt.Print(num, " ")
			case <-quit:
				return // 等效于 runtime.Goexit()
		}
	}
}

func main() {
	ch := make(chan int) // 数字通信
	quit := make(chan bool) // 程序是否结束

	// 消费者，从channel 对内容
	go fibonacci(ch, quit) // 打印操作

	x, y := 1, 1
	for i:=0; i<20; i++ {
		ch <- x
		x, y = y, x+y
	}
	quit <- true
}