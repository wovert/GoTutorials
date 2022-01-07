package main

import (
	"fmt"
	"time"
)

func main() {
	// 构建一个通道
	ch := make(chan string)

	// 开启一个并发匿名函数
	go func() {
		defer fmt.Println("子协程调用完毕")

		for i:=0; i<2; i++ {
			fmt.Println("子协程 i=", i)
			time.Sleep(time.Second)		
		}
		// 通过通道通知 main 的 goroutine
		ch <- "end"
	}()

	fmt.Println("等待子协程序结束")
	// 等待 goroutine
	str := <-ch // 没有数据钱，阻塞
	
	fmt.Println("主协程执行完毕")
	fmt.Println("str=", str)
}