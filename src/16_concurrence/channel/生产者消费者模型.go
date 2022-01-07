package main

import (
	"fmt"
	"time"
)

// 生产者：发送数据端
func producer(out chan<- int, value int) {
	for i:=0; i<50; i++ {
		fmt.Printf("生产者[%d]，生产[%d]\n", value, i)
		out <- value*i
	}
	close(out)
}

// 消费者：接受数据端
func consumer(in <- chan int) {
	for num := range in {
		fmt.Println("消费者，消费：", num)
		time.Sleep(time.Second)
	}

}

func main() {
	//ch := make(chan int) // 同步通信
	ch := make(chan int, 6) // 异步通信

	// 子协程生产者
	for i:=0; i<50; i++{
		go producer(ch, i)
	}


	// 主协程消费者
	consumer(ch)

	for {
		;
	}
}
