package main

import (
  "fmt"
)

// define inteface type
type Humaner interface {
	// 只有方式声明，由自定义类型实现
	sayHi()
}

type Student struct {
	name string
	id int
}

type Teacher struct {
  addr string
  group string
}

// Student 实现sayHi方法
func (stu *Student) sayHi() {
	fmt.Printf("Student [%s, %d] sayHi \n", stu.name, stu.id)
}

// Teacher 实现sayHi方法
func (tmp *Teacher) sayHi() {
	fmt.Printf("Teacher [%s, %s] sayHi \n", tmp.addr, tmp.group)
}

type MyStr string

// MyStr类型 实现 sayHi方法
func (tmp *MyStr) sayHi() {
	fmt.Printf("MyStr [%v] sayHi \n", &tmp)
}

func WhoSayHi(i Humaner) {
	i.sayHi()
}

func main() {
  // 定义接口类型变量
  var i Humaner

  s := &Student{name:"mike", id:10}
  i = s
//   i.sayHi()
  WhoSayHi(i)


  t := &Teacher{addr: "북경", group: "100"}
  i = t
//   t.sayHi(i)
  WhoSayHi(i)

  var str MyStr = "hello mike"
  i = &str
//   t.sayHi(i)
  WhoSayHi(i)

  arr := make([]interface{}, 3)
  arr[0] = 1
  arr[1] = "hello"
  arr[2] = Student{"mike", 22}

  for index, data := range arr {
	// 第一个返回值，第二个返回判断真假 
	// if value, ok := data.(int); ok {
	// 	fmt.Printf("x[%d]类型为int,内容为%d\n", index, value)
	// } else if value, ok := data.(string); ok {
	// 	fmt.Printf("x[%d]类型为string,内容为%s\n", index, value)
	// } else if value, ok := data.(Student); ok {
	// 	fmt.Printf("x[%d]类型为Student,内容为%v\n", index, value)
	// }

		switch value := data.(type) {
			case int:
			fmt.Printf("x[%d]类型为int,内容为%d\n", index, value)
			case string:
			fmt.Printf("x[%d]类型为string,内容为%s\n", index, value)
			case Student:
			fmt.Printf("x[%d]类型为Student,内容为%v\n", index, value)
		}
  }
}