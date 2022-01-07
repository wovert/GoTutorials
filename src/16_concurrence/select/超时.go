package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	o := make(chan bool)

	go func() {
		for {
			select {
			case num := <- c:
				fmt.Println("num = ", num)
			case <-time.After(5 * time.Second):
				fmt.Println("超时")
				o <- true
				//break
				goto label
			}
		}
		label:
			fmt.Println("break to label...")
	}()

	for i:=0; i<2; i++ {
		c <- i
		time.Sleep(time.Second * 2)
	}

	//c <- 666 // 注释掉，引发 timeout
	<-o
	fmt.Println("Finish")
}
