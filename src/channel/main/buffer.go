package main

import (
  "fmt"
  "time"
)

var ch1 chan int

func main() {
	// 创建一个缓存的channel
	ch1 = make(chan int, 3)

	// len(ch) 缓冲区剩余个数，cap(ch)缓冲区大小
	fmt.Printf("len(ch)=%d, cap(ch)=%d\n", len(ch1), cap(ch1))

	// 新建一个协程
  go func (){
		for i:=0; i<10; i++ {
			ch1 <- i // 管道写内容，如果没有人读取会阻塞
			// fmt.Println("子协程：i = ", i)
			fmt.Printf("子协程[%d] len(ch)=%d, cap(ch)=%d\n", i, len(ch1), cap(ch1))
		}
	}()

	time.Sleep(time.Second * 2)

	for i:=0; i<10; i++ {
		num := <-ch1 // 没有内容会阻塞
		fmt.Println("num = ", num)
	}
}