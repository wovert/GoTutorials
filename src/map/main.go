package main


import (
	"fmt"
	"strings"
)


func wordCount(str string) map[string]int {
	s := strings.Fields(str) // 将字符串拆分成字符串切片
	m := make(map[string]int)


	for i:=0; i<len(s); i++ {
		if _, ok := m[s[i]]; ok {
			m[s[i]] = m[s[i]] + 1
		} else {
			m[s[i]] = 1
		}
	}

	return m

}

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

	m4 := make(map[int]string, 2)
	m4[0] = "make"
	m4[1] = "Jack"
	m4[2] = "White"
	fmt.Println("m4 = ", m4)
	fmt.Println("len = ", len(m4))


	str := "How are you and I I I love you very much"
	mRes := wordCount(str)
	for k, v := range mRes {
		fmt.Printf("%q;%d\n", k, v)
	}
}