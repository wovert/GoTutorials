package main

import "fmt"

func test_pipe() {
	pipe := make(chan int, 3)
	pipe <- 1
	pipe <- 2
	pipe <- 3
	// pipe <- 4

	var t int
	t =<- pipe // 从管道Pipe取值放到t变量中

	pipe <- 4 // 覆盖了第一个从1改版为4

	fmt.Println(t)
	fmt.Println(len(pipe))

}