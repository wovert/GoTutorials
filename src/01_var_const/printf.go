package main

import "fmt"

func main() {
	var name string = "kk"
	var number int = 100
	var salary float32 = 50000

	fmt.Printf("%T, %v, %#v\n", name, name, name)
	fmt.Printf("%T, %v, %#v\n", number, number, number)
	fmt.Printf("%T, %v, %#v\n", salary, salary, salary)
}
