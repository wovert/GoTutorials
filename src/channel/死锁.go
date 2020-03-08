package main

import "fmt"

func deadlock1(){
	ch := make(chan int)

	ch <- 720 // 写入数据，写端阻塞，不执行向下执行代码

	num := <-ch // 读取数据

	fmt.Println("num=", num)
}

func deadlock2(){
	ch := make(chan int)
	num := <-ch // 阻塞
	fmt.Println("num=", num)

	go func(){
		ch <- 293
	}()
}

func deadlock3(){
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for {
			select {
				case num := <-ch1: // ch1读操作
					ch2 <- num // ch2写操作
			}
		}
	}()

	for {
		select {
			case num := <-ch2: // ch2读操作
				ch1 <- num // ch1写操作
		}
	}
}

func main() {
	//deadlock1()
	//deadlock2()
	deadlock3()
}
