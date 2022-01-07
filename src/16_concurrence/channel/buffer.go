package main

import (
  "fmt"
  "time"
)

var ch1 chan int

func main() {
	// 创建一个缓存的channel，存满3个元素之前，不会阻塞
	ch1 = make(chan int, 3)

	// len(ch) 查看当前通道的大小，cap(ch)缓冲区大小
	fmt.Printf("len(ch)=%d, cap(ch)=%d\n", len(ch1), cap(ch1))

	// 新建一个协程
  go func (){
		for i:=0; i<10; i++ {
			//fmt.Println("子协程开始")
			ch1 <- i // 管道写内容，如果没有人读取会阻塞
			// fmt.Println("子协程：i = ", i)
			fmt.Printf("子协程[%d] len(ch)=%d, cap(ch)=%d\n", i, len(ch1), cap(ch1))
		}
	}()

	time.Sleep(time.Second * 2)

	for i:=0; i<8; i++ {
		num := <-ch1 // 没有内容会阻塞
		fmt.Println("主协程 num = ", num) // 有延迟
	}
}