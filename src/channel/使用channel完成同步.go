package main

import (
	"fmt"
	"time"
)

// 使用channel完成同步

var c = make(chan int)

func printer(str string) {
	for _, ch := range str {
		fmt.Printf("%c", ch)
		time.Sleep(time.Microsecond * 300)
	}
}

func p1() { // before
	printer("hello ")
	c <- 98
}

func p2() { // after
  <- c
	printer(" word")
}

func main() {
	go p1()
	go p2()
	for {
		;
	}
}