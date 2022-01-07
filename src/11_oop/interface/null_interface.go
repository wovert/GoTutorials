package main


import (
	"fmt"
)
// 空接口
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

func main() {
	// interface{} 空接口类型
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "微明"
	m1["age"] = 300
	m1["married"] = true
	m1["hobby"] = [...]string{"唱","跳"}
	fmt.Println(m1)
	show(false)
	show(20)
	show(m1)
	show(nil)
}