package main

import (
	"fmt"
	"time"
)

// 生产者：发送数据端
func producer(out chan<- int) {
	for i:=0; i<10; i++ {
		fmt.Println("生产者：", i*i)
		out <- i*i
	}
	close(out)
}

// 消费者：接受数据端
func consumer(in <- chan int) {
	for num := range in {
		fmt.Println("消费者拿到：", num)
		time.Sleep(time.Second)
	}

}

func main() {
	//ch := make(chan int) // 同步通信
	ch := make(chan int, 6) // 异步通信

	// 子协程生产者
	go producer(ch)

	// 主协程消费者
	consumer(ch)
}
