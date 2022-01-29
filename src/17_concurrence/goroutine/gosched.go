package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for i:=0; i<100;i++ {
			fmt.Println("子协程 i=", i)
		}
	}()

	for i:=0; i<100;i++ {
		// 出让当前CPU时间轮片，交给其他go协程执行完，在回来执行此协程
		runtime.Gosched()
		fmt.Println("主协程 i=", i)
	}
}