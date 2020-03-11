package main

import (
	"fmt"
	"time"
)

func main() {

	// 创建一个判断是否终止的channel
	quit := make(chan bool)

	fmt.Println("now:", time.Now())

	// 创建周期定时器
	myTicker := time.NewTicker(time.Second) // 1s

	i := 0
	go func(){
		for {
			nowTime := <-myTicker.C
			i++
			fmt.Println("nowTime:", nowTime)
			if i == 6 {
				quit <- true // 解除主go程阻塞，即 <-quit
				fmt.Println("结束")
				break // return runtime.Goexit
			}
		}
	}()

	//for{}
	<-quit // 子协程 循环获取<-myTicker.C 期间，一直阻塞
}
