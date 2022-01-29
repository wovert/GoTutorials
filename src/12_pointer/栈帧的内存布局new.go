package main

import (
	"fmt"
)

func test(m int) {
	var b int = 1000
	b += m
}
func main() {
	// var a int = 10
	// fmt.Println("&a", &a)

	var p *int

	// 在 heap 上申请一片内存地址空间
	p = new(int) // 默认类型的默认值
	*p = 1000

	fmt.Printf("%d\n", *p)
	fmt.Printf("%v\n", *p)
	fmt.Println("*p=", *p)
}
