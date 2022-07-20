package main

import (
	"fmt"
	"reflect"
)

// 结构体的 tag 设置
type People struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

func main() {
	// tag 时类型的一部分，用于通过反射获取字段的相关 tag 设置

	// 通过反射获取结构体成员
	typeOfCat := reflect.TypeOf(People{})
	// 遍历结构体成员
	for i := 0; i < typeOfCat.NumField(); i++ {
		field := typeOfCat.Field(i)
		tag := field.Tag.Get("json")
		fmt.Println(tag)
	}
}
