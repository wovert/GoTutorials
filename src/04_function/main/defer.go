package main

import (
	"fmt"
)

func sum(n1 int, n2 int) int {
	defer fmt.Println("n1=", n1) // 压栈
	defer fmt.Println("n2=", n2) // 压栈
	n1++
	n2++
	res := n1 + n2
	fmt.Println("res=", res)
	return res
}

func closeFile() {
	file = openfile("1.txt")
	defer file.close()
}

func closeDb() {
	connect = openData()
	defer connect.close()
}

func main() {
	res := sum(10, 20)
	fmt.Println("sum=", res)
}