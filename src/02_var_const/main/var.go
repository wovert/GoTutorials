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
}