package main

import (
	"fmt"
	"time"
)

func main() {
	// 延时方法1
	timer := time.NewTimer(2*time.Second)
	<-timer.C
	fmt.Println("时间到1")

	// 延时方法2
	time.Sleep(2 * time.Second)
	fmt.Println("时间到2")

	// 延时方法3
	<- time.After(2 * time.Second)
	fmt.Println("时间到3")
}
