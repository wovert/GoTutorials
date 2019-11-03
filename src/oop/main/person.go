package main


import (
	"fmt"
)

type Person struct {
	name string
	age uint8
}

func (person *Person) setNameAndAge(name string, age uint8) {
	person.name = name
	person.age = age
}
func (person *Person) getNameAndAge()(string, uint8) {
	return person.name, person.age
}
func (person *Person) setName(name string) {
	person.name = name
}
func (person *Person) setAge(age uint8) {
	person.age = age
}
func (person *Person) getName() string {
	return person.name
}
func (person *Person) getAge() uint8 {
	return person.age
}

type Student struct {
	Person
	school string
}

func (student *Student) getSchool() string {
	return student.school
}

// 定义接口
type Animal interface {
	Fly()
	Run()
}

type Bird struct {}
// 实现接口

// 实现 Fly 接口
func (bird Bird) Fly() {
	fmt.Println("Bird is flying!!!")
}
// 实现 Run 接口
func (bird Bird) Run() {
	fmt.Println("Bird is running!!!")
}

func main() {
	person := new(Person)
	// person.setNameAndAge("wovert", 20)
	person.name = "wovert"
	person.age = 30
	name, age := person.getNameAndAge()
	fmt.Printf("name=%s, age=%d\n",name, age)

	person2 := new(Person)
	// person.setNameAndAge("wovert", 20)
	person2.name = "xiuhong"
	person2.age = 50
	name2, age2 := person2.getNameAndAge()
	fmt.Printf("name2=%s, age2=%d\n",name2, age2)

	stu := new(Student)
	stu.name = "lingyima"
	stu.age = 18
	stu.school = "Hight 20"
	school := stu.getSchool()
	fmt.Printf("stu scholl is %s\n", school)

	var animal Animal // 声明接口变量
	bird := new(Bird) // 实例化类
	
	// 一个类只要实现了接口要求的所有函数，这个类实现了该接口
	animal = bird // 对象实例赋值给接口变量
	animal.Fly()
	animal.Run()


	// 任何实例都满足空接口 interface{}
	// var v1 interface{} = 1
	// var v2 interface{} = "abc"
	var v3 interface{} = 3.14
	// var v4 interface{} = make(map[string])

	// v3是否为float64类型 && ok是否为真
	if v,ok := v3.(float64); ok {
		fmt.Println(v, ok)
	} else {
		fmt.Println("v3不是float64类型，且 ok 为false")
	}

	var v4 interface{} = "string value"
	switch v4.(type) {
		case int:
		case float32:
		case float64:
			fmt.Println("this is flaot64")
		case string:
			fmt.Println("this is string")
	}
}