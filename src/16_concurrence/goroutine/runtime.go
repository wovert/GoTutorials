package main

import (
	"fmt"
	"time"
	"runtime"
)

func main() {
	// for i := 0; i < 1000; i++ {
	// 	// 并发执行函数，线程
	// 	go func(i int) { // 匿名函数
	// 		for {
	// 			// Printf IO操作
	// 			fmt.Printf("hello from goroutine %d\n", i)
	// 		}
	// 	}(i)
	// }
	var a [10]int;
	for i := 0; i < 10; i++ {
		// 并发执行函数，线程
		go func(i int) { // race condition
			for {
				a[i]++;
				// 手动交出控制权
				runtime.Gosched()
			}
		}(i)
	}
	fmt.Println("主协程继续执行")
	time.Sleep(time.Millisecond)
	fmt.Println(a)
	fmt.Println("==========")
	
}

// go run --race main.go