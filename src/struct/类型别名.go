package main

import "fmt"

type yourInt = int // 类型别名

func main() {
	var n yourInt
	n = 100
	fmt.Println("类型别名 yourInt n=",n)
}