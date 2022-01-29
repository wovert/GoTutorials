package main

import (
	"fmt"
)

func test(m int) {
	var b int = 1000
	b += m
}
func main() {
	var a int = 10
	var p *int = &a

	a = 100
	fmt.Println("a=", a)

	test(10)

	*p = 250 // 借助a变量的地址，操作a对应空间
	fmt.Println("a=", a)
	fmt.Println("*p=", *p)

	a = 1000
	fmt.Println("*p=", *p)
}
