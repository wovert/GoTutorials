package main

import "fmt"

type Person struct {
	name string
	age int
}

// 构造函数
// 返回的是结构体还是结构体指针
// 当结构体比较大的时候，使用结构体指针，减少程序内存开销
func newPerson(name string, age int) *Person {
	return &Person {
		name: name,
		age: age,
	}
}

func main() {
	p1 := newPerson("老子", 20)
	p2 := newPerson("孔子", 19)
	fmt.Println(p1, p2)
}