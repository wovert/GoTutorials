package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 空接口可以存储任意类型数据，但是无法计算操作
	var i interface{} = 123
	var a interface{} = 234
	// fmt.Println( i + a)

	// 通过反射获取数据类型
	v1 := reflect.TypeOf(i)
	v2 := reflect.TypeOf(a)
	if v1 == v2 {
		// 通过反射获取数据
		v1 := reflect.ValueOf(i).Int()
		v2 := reflect.ValueOf(a).Int()
		result := v1 + v2
		fmt.Printf("%v(%T) + %v(%T) = %v(%T))\n", v1, v1, v2, v2, result, result)
	}
}
