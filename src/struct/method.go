package main

import "fmt"

type Dog struct {
	name string
	age int
}


func newDog(name string, age int) *Dog {
	return &Dog {
		name: name,
		age: age,
	}
}

// 方法是作用于特定类型的函数
// 接受者表示调用该方法的具体类型变量，多用于类型名首字母小写表示
func (d *Dog) run() {
	fmt.Printf("%s: 跑跑~", d.name)
}

func main() {
	d1 := newDog("金毛", 20)
	d1.run()
}
