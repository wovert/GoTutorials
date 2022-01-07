package main

import (
	"fmt"
	"time"
)

func main() {
	// 构建一个通道
	ch := make(chan int)

	// 开启一个并发匿名函数
	go func() {
		fmt.Println("start goroutine")
		// 通过通道通知 main 的 goroutine
		for i :=3 ; i >= 0; i-- {
			// 发送3到0之间的数值
			ch <- i

			// 每次发送完时等待
			time.Sleep(time.Second)
		}
		fmt.Println("exit goroutine")
	}()

	fmt.Println("wait goroutine")
	
	// 便利接受通道数据
	for data := range ch {
		// 打印通道数据
		fmt.Println(data)

		// 当遇到数据0时，退出接受循环
		if data == 0 {
			break
		}
	}
	
	
	fmt.Println("all done")
}