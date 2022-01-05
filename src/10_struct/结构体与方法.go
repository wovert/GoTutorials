package main

import "fmt"

type integer int

type Animal struct {
	Number int
}

type Ani Animal // alias

func main() {
	var i integer = 1000
	var j int = 100

	j = int(i) // 强制转换类型
	fmt.Println(j)

	var a Animal
	a = Animal{10}

	var b Ani
	a = Animal(b) // 强制转换

	fmt.Println("a=", a)
}
