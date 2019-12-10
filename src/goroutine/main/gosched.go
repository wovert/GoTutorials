package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for i:=0; i<100;i++ {
			fmt.Println("this is goroutine test i=", i)
		}
	}()

	for i:=0; i<100;i++ {
		runtime.Gosched() // 出让当前CPU时间
		fmt.Println("this is main test i=", i)
	}
}