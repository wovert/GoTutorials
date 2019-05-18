package main

import (
	"fmt"
)

func main() {
	var a int = 10
	var p *int = &a
	a = 100

	fmt.Println("a=", a)
	*p = 200 // 借助 a 变量的地址，操作 a  对应的空间
	fmt.Println("a=", a)
	fmt.Println("*p=", *p)
}