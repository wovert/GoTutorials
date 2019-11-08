package main

import (
	"fmt"
)

func modify(p *int) {
	fmt.Println(p)
	*p = 123456789
	return 	
}

func main() {
	var a int = 5
	var p *int = &a

	fmt.Println(&a)

	fmt.Println(*p)
	*p = 100
	fmt.Println(a)

	var b int = 999
	p = &b
	*p = 5
	fmt.Println(a)
	fmt.Println(b)

	modify(&a)
	fmt.Println(a)

	// heap 上申请一片内存地址空间
	var pStr *string
	pStr = new(string)
	*pStr = "hello world"
	fmt.Printf("%v\n", *pStr)
	fmt.Printf("%s\n", *pStr)
	fmt.Printf("%q\n", *pStr)

	fmt.Println("=========")

	var x int = 100
	var y = x
	fmt.Printf("%v\n", &x)
	fmt.Printf("%v\n", &y)

}