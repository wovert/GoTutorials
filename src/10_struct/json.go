package main

import (
	"fmt"
	"encoding/json"
)
type Dog struct {
	Name string `json:"name", db:"name", ini:"name"` // json解析时使用小写name
	Age int `json:"age"`
}


func newDog(name string, age int) *Dog {
	return &Dog {
		Name: name, // name 别的包拿不到
		Age: age,
	}
}

func main() {
	d1 := newDog("金毛", 20)
	
	// 结构体 ——> json, 序列化
	b, err := json.Marshal(d1)
	if err != nil {
		fmt.Printf("Marshal failed, err: %v, err")
		return
	}
	fmt.Printf("%#v\n", string(b))

	// json ——> 结构体, 反序列化
	str := "{name:\"wovert\", age:20}"
	var d2 Dog
	// 传指针视为能在json.Unmarshal内部修改p2值
	json.Unmarshal([]byte(str), &d2)
	fmt.Printf("%#v\n", d2)
}
