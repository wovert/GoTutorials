package main

import "fmt"

type people struct {
	Name string
}

func main01() {
	// 定义接口类型变量
	var i interface{} = &people{"张三"}
	// 类型断言
	if v, ok := i.(*people); ok {
		fmt.Println(v.Name)
	} else {
		fmt.Println("类型匹配失败")
	}
}

func main02() {
	// 定义接口类型变量
	var i interface{} = &people{"李四"}
	switch i.(type) {
	case people: // 对象类型
		fmt.Println("people")
	case *people: // 指针类型
		fmt.Println("*people")
	default:
		fmt.Println("类型匹配失败")
	}
}

func main() {
	//main01()
	main02()
}
