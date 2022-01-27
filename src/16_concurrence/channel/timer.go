package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 定时方法
	time.Sleep(time.Second)

	fmt.Println("当前时间：", time.Now())

	// 2. 定时方法
	// 创建一个定时器，设置时间为2s, 2s后，往 myTimer 通道写当前系统时间
	myTimer := time.NewTimer(time.Second * 2) // 2s

	// 2s后，往myTimer.C写输入，有数据后就可以读取，没有数据不去读会一直阻塞
	nowTime := <-myTimer.C
	fmt.Println("现在时间：", nowTime)

	// 3. 定时方法，直接返回定时之后 类似 <-myTimer.C
	nowTime2 := <-time.After(time.Second * 2)
	fmt.Println("After现在时间：", nowTime2)

	// 4. 定时器停止和重置
	myTimer3 := time.NewTimer(time.Second * 3)

	// 重置定时时长为 1
	myTimer3.Reset(1 * time.Second)
	go func() {
		nowTime := <-myTimer3.C // 定时器停止，系统不会写入数据，因此不能读取时间。会阻塞，不会执行下面语句
		fmt.Println("子协程，定时完毕:", nowTime)
	}()

	myTimer3.Stop() // 定时器停止，即取消 time.Second * 3, <-mytimer3.C 阻塞

	for {
		;
	}
}
