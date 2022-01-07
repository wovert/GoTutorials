package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	timeout := make(chan int, 1)
	
	go func() {
	  // time.Sleep(time.Second)
	  timeout <- 1
	}()

	go func() {
	  // time.Sleep(time.Second)
	  ch <- 30
	}()



	time.Sleep(time.Second)

	// 监听channel上数据流动, ch,timeout 谁先来，随机选择执行
	select {
		case <-ch:
			fmt.Println("come to read ch!")
		case <-timeout:
			fmt.Println("come to timeout!")
		// default:
		// 	fmt.Println("default执行语句")
	}
}