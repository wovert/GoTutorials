package main

import "fmt"

type OrderInfo struct {
	id int
}

// 生成订单 —— 生产者
func producer(out chan<- OrderInfo) {
	// 模拟生成10分订单
	for i := 0; i < 10; i++ {
		order := OrderInfo{id: i + 1}
		out <- order
		fmt.Println("生成订单id为：", i+1)
	}
	close(out)
}

// 处理订单 —— 消费者
func consumer(in <-chan OrderInfo) {
	for order := range in {
		// 模拟处理订单
		fmt.Println("订单 id为：", order.id)
	}
}

func main() {
	ch := make(chan OrderInfo)
	go producer(ch) // 只写
	consumer(ch)    // 只读
}
