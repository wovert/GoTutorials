package main

import (
	"fmt"
	"unsafe"
)

type S struct {
	A struct{}
	B struct{}
}

func main() {
	null := struct{}{}
	fmt.Println(unsafe.Sizeof(null)) // 0
	fmt.Printf("%p\n", &null)

	// 定义多个空结构体都指向同一个内存地址（全局区）
	null2 := struct{}{}
	fmt.Printf("%p\n", &null2)

	s := S{}
	fmt.Println(s.A) // {}
	fmt.Println(s.B) // {}
}
