package main

import "fmt"

func main() {
	//deadlock1()
	//deadlock2()
	deadlock3()
}

// 单go程自己死锁
func deadlock1(){
	ch := make(chan int)

	ch <- 720 // 写端阻塞，不会向下执行代码
	num := <-ch // 读取数据

	fmt.Println("num=", num)
}

// go程间channel访问顺序导致死锁
func deadlock2(){
	ch := make(chan int)

	num := <-ch // 读段阻塞（没有数据可读），先读后写导致死锁
	fmt.Println("num=", num)

	// 子go程序调用位置放错
	go func(){
		ch <- 293
	}()
}

// 多go程，多channel交叉死锁
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