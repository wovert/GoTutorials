package main

import (
	"fmt"
	"time"
)

func sing() {
	for i :=1; i <= 100; i++{
		fmt.Printf("正在唱歌，隔壁泰山%d\n", i)
	}
}

func dance() {
	for i :=1; i <= 100; i++{
		fmt.Printf("正在跳舞，赵三街舞%d\n", i)
	}
}


func demo01() {
	// 子协程
	go sing()

	// 子协程
	go dance()

	// 死循环，不让主协程结束
	for {
		;
	}
}


func demo02() {
	// 创建一个 子 go协程
	go func () {
		for i:=0; i<5; i++ {
			fmt.Println("----go routine----")
			time.Sleep(time.Second)
		}
	}()

	for i:=0; i<5; i++ {
		fmt.Println("----I'm demo02----")
		time.Sleep(time.Second)
		if i == 2 {
			break;
		}
	}
}

func main() {
	//demo01()
	demo02()

}