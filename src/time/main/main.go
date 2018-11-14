package main

import (
	"fmt"
	"time"
)

func test() {
	time.Sleep(time.Millisecond * 100)
}

func main() {
	now := time.Now()
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("2006/1/2 15:4:5"))
	start := time.Now().UnixNano() // 纳秒
	test()
	end := time.Now().UnixNano() // 纳秒
	fmt.Printf("cost:%d us\n", (end-start)/1000) // 微妙


}