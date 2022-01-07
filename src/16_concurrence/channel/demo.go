package main

import (
  "fmt"
  "time"
)

func main() {
	// 创建通道数据类型
	ch := make(chan string)

	defer fmt.Println("主协程序结束")

	go func(){
		defer fmt.Println("子协程调用完毕")
		for i:=0; i<2; i++ {
			fmt.Println("子协程i=",i)
			time.Sleep(time.Second)
		}
		ch <- "我是子协程，工作完毕" // 写入到通道，没人读取ch会阻塞
	}()

	str := <-ch // 没有数据前它会阻塞，读取通道并赋值给str变量
	fmt.Println("str=", str)
}