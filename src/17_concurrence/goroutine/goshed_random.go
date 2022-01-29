package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [10]int
	for i:=0; i<10; i++ {
		go func(i int) { // race condition
			for {
				a[i]++
				runtime.Gosched() // 手动交出控制权
			}
		}(i)
	}
	time.Sleep(time.Millisecond * 1000)
	fmt.Println(a)
}
