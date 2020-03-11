package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	timeout := make(chan int, 1)
	
	go func() {
	  time.Sleep(time.Second)
	//   ch <- 30 
	  timeout <- 1
	}()

	// time.Sleep(time.Second)

	// 坚挺channel上数据流动
	select {
	case <-ch:
		fmt.Println("come to read ch!")
	case <-timeout:
		fmt.Println("come to timeout!")
	// default:
	// 	fmt.Println("default执行语句")
	}
}