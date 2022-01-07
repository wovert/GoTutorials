package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("当前时间：", time.Now())
	
	// 定时器
	timer := time.NewTimer(time.Second * 3)

	// 重置定时时长为 1
	timer.Reset(1 * time.Second)
	go func() {
		nowTime := <-timer.C // 定时器停止，系统不会写入数据，因此不能读取时间。会阻塞，不会执行下面语句
		fmt.Println("子协程，定时完毕:", nowTime)
	}()

	// timer.Stop() // 停止定时器，即取消 time.Second * 3

	for {}
}
