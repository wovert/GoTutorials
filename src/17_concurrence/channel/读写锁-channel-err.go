package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 读写模式，一把锁
var rwMutex sync.RWMutex

func readGo(in <-chan int, i int) {
	for {
		rwMutex.RLock() // 以读模式加锁
		num := <-in // 读取数据（但没有数据，会阻塞）
		fmt.Printf("-----%dth 读 协程,读出：%d\n", i, num)
		rwMutex.RUnlock()	// 以读模式解锁
	}

}

func writeGo(out chan<- int, i int) {
	// 生成随机数
	num := rand.Intn(1000)

	rwMutex.Lock()	// 以写模式加锁 (因为readGo已经有锁定，所以不能拿到锁，所以会阻塞)
	out <- num // 因为加锁，不能读取通道数据
	fmt.Printf("%dth 写协程，写入:%d\n", i, num)
	//time.Sleep(time.Microsecond * 300) // 方法实现现象
	rwMutex.Unlock() // 解锁
}

func main() {
	// 播种随机数种子
	rand.Seed(time.Now().UnixNano())

	//quit := make(chan bool) // 关闭主协程channel
	ch := make(chan int) // 数据传递的channel

	for i:=0; i<5; i++ {
		go readGo(ch, i+1)
	}

	for i:=0; i<5; i++ {
		go writeGo(ch, i+1)
	}

	for {
		;
	}
}