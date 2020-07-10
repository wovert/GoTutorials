package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("当前时间：", time.Now())

	// 1. 定时方法
	time.Sleep(time.Second)

	// 2. 定时方法
	// 创建定时器 2s之后，操作系统向time结构体写入系统当前时间
	myTimer := time.NewTimer(time.Second * 2) // 2s

	// 定时满，系统自动写入系统时间。读操作 chan类型，不去读会一直阻塞
	nowTime := <-myTimer.C
	fmt.Println("现在时间：", nowTime)

	// 3. 定时方法，直接放回定时之后
	nowTime2 := <-time.After(time.Second * 2)
	fmt.Println("现在时间：", nowTime2)

	// 4. 定时器停止和重置
	myTimer3 := time.NewTimer(time.Second * 3)

	// 重置定时时长为 1
	myTimer3.Reset(1 * time.Second)
	go func() {
		nowTime := <-myTimer3.C // 定时器停止，系统不会写入数据，因此不能读取时间。会阻塞，不会执行下面语句
		fmt.Println("子协程，定时完毕:", nowTime)
	}()

	myTimer3.Stop() // 定时器停止，即取消 time.Second * 3

	for {
		;
	}
}
