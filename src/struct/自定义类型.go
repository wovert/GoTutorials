package main

import "fmt"

// type 后面跟的是类型
type myInt int // 自定义类型

func main() {
	var n myInt
	n = 100
	var c rune
	c = '中'
	fmt.Println("自定义类型 myInt n=",n)
	fmt.Printf("c=%T", c)
}