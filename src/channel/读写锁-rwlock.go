package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var rwMutex3 sync.RWMutex // 锁只有一把，2个属性
var value int // 定义全局变量，模拟共享数据

func readGo2(i int) {
		rwMutex3.RLock() // 以读模式加锁，同时不能写入数据
		num := value // 读取数据，共享数据
		fmt.Printf("%dth 读协程，读出：%d\n", i, num)
		rwMutex3.RUnlock()	// 以读模式解锁
}

func writeGo2(i int) {
	// 生成随机数
	num := rand.Intn(1000)

	rwMutex3.Lock()	// 以写模式加锁，独占，同时不能读取
	value = num
	fmt.Printf("-----%dth 写协程，写入:%d\n", i, num)
	time.Sleep(time.Microsecond * 200)
	rwMutex3.Unlock() // 解锁
}

func main() {
	// 播种随机数种子
	rand.Seed(time.Now().UnixNano())

	for i:=0; i<10; i++ {
		go readGo2(i+1)
	}

	for i:=0; i<10; i++ {
		go writeGo2(i+1)
	}


	for {
		;
	}
}