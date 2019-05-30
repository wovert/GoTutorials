package main

import (
	"fmt"
)

type Person struct {
	name string
	age int
}

type Student struct {
	Person
	age int
	speciality string
}

func main() {
	// person := Person{"张三", 23}
	person := Person{age:23, name:"张三"}
	fmt.Printf("%v\n", person.name)

	student := Student{Person{"wovert", 29}, 50, "maths"}
	fmt.Printf("%v\n", student)
	fmt.Printf("%v\n", student.Person.name)
	fmt.Printf("%v\n", student.age)
}