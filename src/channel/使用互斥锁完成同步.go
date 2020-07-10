package main

import (
	"fmt"
	"sync"
	"time"
)

// 使用锁完成同步 - 互斥锁

var mutex sync.Mutex // 创建一个互斥量，新建的互斥锁状态为 0：未加锁，锁只有一把


func printer2(str string) {
	mutex.Lock() // 访问共享数据之前，加锁
	for _, ch := range str {
		fmt.Printf("%c", ch)
		time.Sleep(time.Microsecond * 300)
	}
	mutex.Unlock() // 共享数据访问结束，解锁
}

func p_1() { // before
	printer2("hello ")
}

func p_2() { // after
	printer2(" word")
}

func main() {
	go p_1()
	go p_2()
	for {
		;
	}
}