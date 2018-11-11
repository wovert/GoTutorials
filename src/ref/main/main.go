package main

import "fmt"

func modify(a *int) {
	*a = 10
	return
}

func main() {
	a := 5
	b := make(chan int, 1) // 存储地址

	fmt.Println("a=", a)
	fmt.Println("b=", b)

	// modify(a) // 值传递
	fmt.Println("a=", a) // 5
	modify(&a) // 引用传递
	fmt.Println("a=", a)
}