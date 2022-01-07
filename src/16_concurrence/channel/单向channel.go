package main

import "fmt"

func main() {
	ch := make(chan int) // 双向 channel

	go func() {
		var sendCh chan <- int = ch // 单向写channel

		sendCh <- 120 // 阻塞
	}()

	// 阻塞
	var recvCh <-chan int = ch // 单向读channel
	num := <- recvCh

	fmt.Println("num=", num)

	// 反向赋值
	//var c5 chan int = recvCh // error

	/////////// 函数传参 ////////////////
  c2 := make(chan int)
  go func(){
  	send(c2)
	}()
	recv(c2)
}

// 参数：单向写channel
func send(out chan <- int) {
	out <- 299
	close(out)
}

// 传参：单向读channel
func recv(in <- chan int) {
	n := <-in
	fmt.Println("读到n=", n)
}