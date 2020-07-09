package main

import "fmt"

type Person1 struct {
	name string
	age int
}

type Animal1 struct {
	*Person1
	color string
}

func main() {
	// 1
	a1 := Animal1{
		&Person1{"名字", 18},"红色",
	}
	fmt.Printf("a1=%T\n", a1)

	// 2
	var a2 = new(Animal1)
	a2.color = "red"
	a2.Person1 = new(Person1)
	a2.Person1.name = "name2"
	a2.Person1.age = 20
	fmt.Printf("a2=%T\n", a2)

	// 3
	var a3 = new(Animal1)
	a3.color = "green"
	a3.Person1 = &Person1{"name3", 29}
	fmt.Printf("a3=%T\n", a3)
}