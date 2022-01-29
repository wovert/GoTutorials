package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个 子 go协程
	go func () {
		i := 0
		for {
			i++
			fmt.Println("----子协程----: i=", i)
			time.Sleep(time.Second)
		}
	}()

	i := 0
	for {
		i++
		fmt.Println("----主协程----: id=", i)
		time.Sleep(time.Second)

		if i == 2 {
			break
		}
	}
}
