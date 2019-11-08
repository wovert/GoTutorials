package main

import "fmt"

type Human struct {
	name string
	age uint8
	weight uint
}

type Student struct {
	Human // 匿名字段，默认Student 包含了Human的所有字段
	speciality string
}

func main() {
	person := Human{name: "张三", age: 25}
	fmt.Printf("%v", person)

	lisi := &Human{name: "李四", age: 25}
	fmt.Printf("%T\n", lisi)
	fmt.Printf("%v\n", lisi)

	var h Human
	var p1 *Human
	p1 = &h
	p1.name = "王武"
	(*p1).age = 30
	fmt.Printf("%T\n", p1)
	fmt.Printf("%v\n", p1)

	// 通过new申请结构体
	p2 := new(Human)
	p2.name = "wovert"
	p2.age = 29
	fmt.Printf("%T\n", p2)
	fmt.Printf("%v\n", p2)


	jane := Student{Human{"Jane",35,100},"blology"}
	fmt.Printf("%v", jane)
	fmt.Println()
	fmt.Printf("%v", jane.age)
	fmt.Println()

	jack := Student{Human{"Jane", 30, 30},"how"}
	fmt.Printf("%v", jack)
	fmt.Println()
}
