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
	person := Human{"张三", 25, 60}
	fmt.Printf("%v", person)

	jane := Student{Human{"Jane",35,100},"blology"}
	fmt.Printf("%v", jane)
	fmt.Println()
	fmt.Printf("%v", jane.age)
	fmt.Println()
}
