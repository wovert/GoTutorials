package main

import (
	"fmt"
)

var (
	flag = "global"
	status = true
)
func main() {
	var a int
	var b string = "hello"
	c := true
	name, age := "wovert", 20
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("c=", c)
	fmt.Println("name=", name)
	fmt.Println("age=", age)
	fmt.Println("flag=", flag)
	fmt.Println("status=", status)

	// iota 常量自动生成器，每个一行，自动累加1 
	const (
		x = iota
		y = iota
		z = iota
	)
	fmt.Println("x=", x) // 0
	fmt.Println("y=", y) // 1
	fmt.Println("z=", z) // 2

	const xyz = iota
	fmt.Println("xyz=", xyz) // 0

	const (
		j = iota
		k
		l
	)
	fmt.Println("j=", j) // 0
	fmt.Println("k=", k) // 1
	fmt.Println("l=", l) // 2

	const (
		a1 = iota
		// 同一行，值都一样
		b1, b2, b3 = iota, iota, iota
		c1 = iota
	)
	fmt.Println("a1=", a1) // 0
	fmt.Println("b1=", b1) // 1
	fmt.Println("b2=", b2) // 1
	fmt.Println("b3=", b3) // 1
	fmt.Println("b1=", c1) // 2
}