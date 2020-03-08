package main


import (
	"fmt"
)


func main() {
	ch3 := make(chan int)
	go func() {
		for i :=0; i<8; i++ {
			ch3 <- i
		}
		close(ch3) // 写端，写完数据主动关闭 channel
		//ch3 <- 10
	}()

	for {
		if num, ok := <- ch3; ok == true {
			fmt.Println("读到数据：", num)
		} else {
			n := <- ch3
			fmt.Println("关闭后, 读取chanel为", n)
			break
		}
	}
}