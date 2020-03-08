package main

import (
	"fmt"
	"time"
)

func main() {
	//singAndDance()
	demo02()
}

func singAndDance() {

	// 共同抢占 CPU 时间轮片
	go sing() // 子协程
	go dance() // 子协程

	// 死循环，不让主协程结束
	for {
		;
	}
}

func sing() {
	for i :=1; i <= 100; i++{
		fmt.Printf("正在唱歌，隔壁泰山%d\n", i)
		if (i % 10 == 0) {
			// 放弃CPU时间片执行，CPU执行其他协程程序，即 dance()函数
			time.Sleep(100 * time.Microsecond)
		}
	}
}

func dance() {
	for i :=1; i <= 100; i++{
		fmt.Printf("正在跳舞，赵三街舞%d\n", i)
		if (i % 10 == 0) {
			// 放弃CPU时间片执行，CPU执行其他协程程序，即 sing()函数
			time.Sleep(100 * time.Microsecond)
		}
	}
}





func demo02() {
	// 创建一个 子 go协程
	go func () {
		for i:=0; i<5; i++ {
			fmt.Println("----go routine----:", i)
			time.Sleep(time.Second)
		}
	}()

	for i:=0; i<5; i++ {
		fmt.Println("----I'm main routine----:", i)
		time.Sleep(time.Second)
		if i == 2 {
			break;
		}
	}
}