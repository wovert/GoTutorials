package main


import (
	"fmt"
)

func main() {
	// key: string类型
	// value: string类型
	m1 := map[string]string {
		"code": "1",
		"msg": "没有信息",
		"data": "no data",
	}
	// 删除元素
	delete(m1, "data")
	for k, v := range m1 {
		fmt.Println(k,"=",v)
	}

	// 是否存在 key
	if cause, ok := m1["cause"]; ok {
		fmt.Println("cause=",cause)
	} else {
		fmt.Println("没有cause")
	}
	fmt.Println("==========")
	fmt.Printf("%v\n", m1)
	fmt.Printf("%v\n", m1["msg"])

	m2 := make(map[string]int) // m2 == empty map
	var m3 map[string]int // m3 == nil
	fmt.Println(m1, m2, m3)
}