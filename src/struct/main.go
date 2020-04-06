package main

import (
	"fmt"
	"unsafe"
)

type Human struct {
	name string
	age uint8
	weight uint
}

type Student struct {
	Human // 匿名字段，默认Student 包含了Human的所有字段
	speciality string
}

type Person struct {
	name string
	sex byte
	age int
}

// 方法
func test(man Person) {
	man.age = 100
	fmt.Println("test man size:=", unsafe.Sizeof(man))
}

func ptest(man *Person) {
	man.age = 100
	fmt.Println("ptest man size:=", unsafe.Sizeof(man))
}


func main() {
	// 顺序初始化
	var man = Person{"wovert", 'm', 18}
	fmt.Println("man=", man)

	// 部分初始化
	man2 := Person{name:"wovert", age:18 }
	fmt.Println("man2=", man2)
	man2.sex = 'm'

	// 索引成员变量
	fmt.Println("man2.name=", man2.name)


	// 普通变量的赋值和使用
	var man3 Person
	man3.name = "mike"
	man3.sex = 'f'
	man3.age = 20
	fmt.Println("man3=", man3)
	fmt.Println("main man3 size:=", unsafe.Sizeof(man3))


	// 结构体变量的比较和赋值
	// 1. compare: 只能使用 == or !=
	fmt.Println("man == man2 ?", man == man2) // true
	fmt.Println("man == man3 ?", man == man3) // false

	// 相同类型的结构体赋值
	var tmp Person
	fmt.Println("tmp=", tmp)


	tmp = man3

	// 函数内部使用结构体传参
	test(man3)	// 值传递
	fmt.Println("man3=", man3)

	ptest(&man3)	// 引用传递
	fmt.Println("man3=", man3)

	//  结构体变量的地址 = 结构体首个元素的地址
	fmt.Printf("&man = %p\n", &man3)
	fmt.Printf("&man.name = %p\n", &(man3.name))

	// 结构体指针
	var person1 *Person = &Person{"wovert", 'f', 19}
	fmt.Println("person1=", person1)

	// 堆创建结构体对象
	person2 := new(Person)
	person2.name = "xiuhong"
	person2.age = 20
	fmt.Println("person2=", person2)


	return


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
