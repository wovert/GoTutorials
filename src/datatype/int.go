package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var n2 int64 = 10
	fmt.Printf("n2 的类型 %T n2占用的字节数为 %d ", n2, unsafe.Sizeof(n2))
	i1 := int8(3) // 默认int类型
	i2 := int16(30)
	i3 := int32(i1)
	fmt.Printf("i1=%d, type=%T\n", i1, i1)
	fmt.Printf("i2=%d, type=%T\n", i2, i2)
	fmt.Printf("i3=%d, type=%T\n", i3, i3)
}