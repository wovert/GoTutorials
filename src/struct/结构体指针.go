package main

import (
	"fmt"
)

type person struct {
	name string
	gender string
	age int
}

func main() {
	// map类型 结构体数组指针
	m := make(map[int]*[3]person)
	fmt.Printf("%T\n", m)

	m[1] = new([3]person) // 返回值类型为*[3]person
	m[1] = &[3]person{
		{name: "钢铁侠", gender: "男", age: 18},
		{name: "超人", gender: "男", age: 20},
		{name: "女超人", gender: "女", age: 39},
	}

	m[2] = new([3]person) // 返回值类型为*[3]person
	m[2] = &[3]person{
		{name: "钢铁侠2", gender: "男", age: 18},
		{name: "超人2", gender: "男", age: 20},
		{name: "女超人2", gender: "女", age: 39},
	}

	for k, v := range m {
		fmt.Println(k, *v)
	}
}