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
}