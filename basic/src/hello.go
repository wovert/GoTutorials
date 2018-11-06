package main

import "fmt"

func add(a int, b int) int {
	var sum int
	sum =  a + b
	return sum
}

func main() {
	var c int 
	c = add(100, 200)
	fmt.Println(c)
	fmt.Print("add(100,200)=", c)
}