package main

import "fmt"

func main() {
	var i interface{}
	i = 10

	// 值, 值得判断 := 接口变量.(数据类型)
	if value, ok := i.(int); ok {
		fmt.Println("整型数据：", value)
	} else {
		fmt.Println("错误")
	}
}