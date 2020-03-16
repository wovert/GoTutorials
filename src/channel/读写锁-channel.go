package main

import (
	"fmt"
	"math/rand"
	"time"
)


func readGo8(in <-chan int, i int) {
	//for {
		num := <-in // 从 channel 读取数据
		fmt.Printf("%dth 读协程，读出：%d\n", i, num)
	//}
}

func writeGo8(out chan<- int, i int) {
	//for {
		// 生成随机数
		num := rand.Intn(1000)

		// 写入到管道
		out <- num
		fmt.Printf("-----%dth 写协程，写入:%d\n", i, num)
		time.Sleep(time.Microsecond * 300)
	//}
}

func main() {
	// 播种随机数种子
	rand.Seed(time.Now().UnixNano())
	ch := make(chan int)

	for i:=0; i<5; i++ {
		go readGo8(ch, i+1)
	}

	for i:=0; i<5; i++ {
		go writeGo8(ch, i+1)
	}

	for {
		;
	}
}