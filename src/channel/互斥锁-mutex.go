package main

import (
	"fmt"
	"sync"
	"time"
)


// 使用"锁"完成同步 —— 互斥锁

var mutex2 sync.Mutex // 创建互斥量，新建的互斥锁状态为 0:未加锁，锁只有一把

func printer88(str string) {
	fmt.Printf("\nstart str=%s\n", str)
	mutex2.Lock() // 访问共享数据之前，加锁
	for _, ch := range str { // str共享数据不能同时访问
		fmt.Printf("%c", ch)
		time.Sleep(time.Millisecond * 300)
	}
	mutex2.Unlock() // 共享数据访问结束，解锁
	fmt.Printf("\nend str=%s\n", str)
}

func person10() { // 先执行
	printer88("hello")
}

func person20() { // 后执行
	printer88("world") // 读取数据
}

func main() {
	go person10();
	go person20();
	for {
		;
	}
}