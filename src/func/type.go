package main

import "fmt"

func main() {
	type sum func(x,y int)int
	var f sum = func(x,y int) int {
		return x + y
	}
	fmt.Println(f(3,4))
}