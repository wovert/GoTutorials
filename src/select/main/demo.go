package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := make(chan int) // 数据通信
  quit := make(chan bool) // 是否退出

  go func() {	// 写数据
		for i:=0; i<5; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
		quit <- true // 通知主协程退出
		runtime.Goexit() // 退出自己协程
	}()

  for {
  	select {
			case num := <- ch:
				fmt.Println("读到：", num)
			case <-quit:
				return // 终止进程
				//break // 跳出select
		}
  	fmt.Println("======")
	}

}