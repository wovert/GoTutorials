package main

import (
	"fmt"
)

func main() {
	a := 10
	p := &a // 一级指针 指向变量的地址

	var pp **int
	pp = &p // 二级指针 指向一级指针的地址

	var ppp ***int
	ppp = &pp // 三级指针 指向二级指针的地址
	
	fmt.Printf("%T\n", p) // *int
	fmt.Printf("%T\n", pp) // **int
	fmt.Printf("%T\n", ppp) // ***int

	*p = 100
	fmt.Printf("a=%d\n", a)

	**pp = 1000
	fmt.Printf("a=%d\n", a)

	***ppp = 10000
	fmt.Printf("a=%d\n", a)
}