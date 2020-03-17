package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 生产者：发送数据端
func producer3(out chan<- int, index int) {
	for i:=1; i<10; i++ {
		num := rand.Intn(800)
		fmt.Printf("生产者%dth, 生产%d\n", index, num)
		out <- num
	}
	close(out)
}

// 消费者：接受数据端
func consumer3(in <- chan int, index int) {
	for num := range in {
		fmt.Printf("----消费者%dth, 消费%d\n", index, num)
		//time.Sleep(time.Millisecond * 300)
	}
}

func main() {
	ch := make(chan int) // 同步通信
	//ch := make(chan int, 6) // 异步通信
	rand.Seed(time.Now().UnixNano())

	// 子协程生产者
	for i:=1; i<5; i++{
		go producer3(ch, i)
	}


	// 主协程消费者
	for i:=1; i<5; i++{
		go consumer3(ch, i)
	}

	for {
		;
	}
}
