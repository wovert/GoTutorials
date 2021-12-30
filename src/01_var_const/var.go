package main

import (
	"fmt"
)

// 包级别
var packageVar string = "package Var"
var (
	flag   = "global"
	status = true
)

func main() {
	// 函数级别
	var a int
	var b string = "hello"
	c := true
	var d *int
	d = &a
	name, age := "wovert", 20
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("c=", c)
	fmt.Println("d=", &d)
	fmt.Println("name=", name)
	fmt.Println("age=", age)
	fmt.Println("flag=", flag)
	fmt.Println("status=", status)

	{
		// 块级别
		var inlineBlockVar string = "inline block var"
		fmt.Println(inlineBlockVar)
		{
			// 子块级别
			var innerInlineBlockVar string = "inner Inline Block var"
			fmt.Println(innerInlineBlockVar)
		}
	}

	// iota 常量自动生成器，每个一行，自动累加1
	const (
		x    = iota
		y    = 100
		z    = iota
		_    = iota
		last = iota
	)
	fmt.Println("x=", x)       // 0
	fmt.Println("y=", y)       // 100
	fmt.Println("z=", z)       // 2
	fmt.Println("last=", last) // 4

	// 每新增一行常量声明将 iota 计数一次
	const (
		d1, d2 = iota + 1, iota + 2 // 1, 2
		d3, d4 = iota + 1, iota + 2 // 2, 3
	)
	fmt.Println("d1=", d1) // 1
	fmt.Println("d2=", d2) // 2
	fmt.Println("d3=", d3) // 2
	fmt.Println("d4=", d4) // 3

	// 定义数量级
	const (
		_  = iota
		KB = 1 << (10 * iota) // 1 << 10
		MB = 1 << (10 * iota) // 1 << 20
		GB = 1 << (10 * iota) // 1 << 30
		TB = 1 << (10 * iota) // 1 << 40
		PB = 1 << (10 * iota) // 1 << 50
	)

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
		c1         = iota
	)
	fmt.Println("a1=", a1) // 0
	fmt.Println("b1=", b1) // 1
	fmt.Println("b2=", b2) // 1
	fmt.Println("b3=", b3) // 1
	fmt.Println("b1=", c1) // 2

	var no int = 20
	var noPointer *int = &no

	fmt.Printf("no=%v\n", no)
	fmt.Printf("noPointer=%v\n", noPointer)

	// *noPointer // 指针变量值
	*noPointer = 30 // 设置指针变量的值
	fmt.Printf("no=%v\n", no)
	fmt.Printf("noPointer=%v\n", *noPointer) // 读取指针变量值

	//////////
	var x1 int = 10
	var xP *int = &x1
	var xPP **int = &xP

	fmt.Printf("x1=%v\n", x1)
	fmt.Printf("xP=%v\n", *xP)
	fmt.Printf("xPP=%v\n", **xPP)

	fmt.Printf("xP=%v\n", xP)
	fmt.Printf("xPP=%v\n", xPP)
	fmt.Printf("*xPP=%v\n", *xPP)

}
