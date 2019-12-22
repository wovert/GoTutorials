package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("当前时间：", time.Now())
	// 创建定时器，2s之后，系统向结构体写入时间
	myTimer := time.NewTimer(time.Second * 2) // 2s
	nowTime := <-myTimer.C // 读操作 chan类型
	fmt.Println("现下时间：", nowTime)

	nowTime2 := <-time.After(time.Second * 2)
	fmt.Println("现下时间：", nowTime2)

	// 定时器停止和重置
	myTimer3 := time.NewTimer(time.Second * 3)
	myTimer3.Reset(1 * time.Second) // 重置定时时长为 1
	go func() {
		nowTime := <-myTimer3.C // 定时器停止，系统不会写入数据，因此不能读取时间
		fmt.Println("子协程，定时完毕:", nowTime)
	}()
	//myTimer3.Stop() // 定时器停止，即取消 time.Second * 3
	for {

	}
}
