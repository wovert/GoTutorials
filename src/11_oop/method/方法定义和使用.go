package main

import "fmt"

// 定义结构体表示对象属性
type Student struct {
	Name string
	age  int
}

// 对象的行为用方法表示
// func (方法接受者)方法名(参数)(返回值){ return value}
// 值传递
func (stu Student) SayHiByValue() {
	stu.Name = "修改值传递对象的名字"
	fmt.Println("大家好，欢迎到Wovert大学。我是", stu.Name)
}

// 引用传递
func (stu *Student) SayHiByRef() {
	stu.Name = "修改引用对象的名字"
	fmt.Println("大家好，欢迎到Wovert大学。我是", stu.Name)
}
func main() {
	//
	stu := Student{"张三", 18}
	stu.SayHiByValue()
	fmt.Println("我是", stu.Name)

	stuRef := &Student{"李四", 20}
	stuRef.SayHiByRef()
	fmt.Println("我是", stuRef.Name)
}
