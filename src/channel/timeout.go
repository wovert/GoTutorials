package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	quit := make(chan bool)

	go func() {	// 子协程获取数据
		for {
			select {
				case num := <-ch:
					fmt.Println("num = ", num)
				case <-time.After(3 * time.Second):
					quit <- true
					//break // runtime.Goexit()
					goto label
			}
		}
		label:
			fmt.Println("break to label...")
	}()

	for i:=0; i<2; i++ {
		ch <- i
		time.Sleep(time.Second * 2)
	}

	<-quit	// 等待子协程通知，退出
	fmt.Println("finish!")
}
