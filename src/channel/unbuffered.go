package main

import (
	"fmt"
	"time"
)

var ch2 chan int

func main() {
	// 创建一个无缓存的channel，即不能存储数据
	ch2 = make(chan int, 0) // 等价于 make(chan int)

	// len(ch) 缓冲区剩余未读取数据个数，cap(ch)通道容量
	fmt.Printf("len(ch)=%d, cap(ch)=%d\n", len(ch2), cap(ch2))

	// 新建一个协程
  go func (){
		for i:=0; i<10; i++ {
			fmt.Println("子协程：i = ", i)
			ch2 <- i // 管道写内容，如果没有人读取会阻塞等待. IO延迟
		}
	}()

	time.Sleep(time.Second * 2)

	for i:=0; i<10; i++ {
		num := <-ch2 // 如果没有数据写入到管道，则会阻塞
		fmt.Println("主协程, num = ", num) // 随机调度：打印是 IO 操作，耗时（访问硬件），阻塞排队。即子协程继续运行
	}
}