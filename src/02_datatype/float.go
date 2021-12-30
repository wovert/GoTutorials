package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	var f1 float64
	f1 = 3.14
	var f2 float32
	f2 = 0.5
	fmt.Printf("f1 的类型 %T f1占用的字节数为 %d\n", f1, unsafe.Sizeof(f1))
	fmt.Printf("f2 的类型 %T %v f1占用的字节数为 %d\n", f2, f2, unsafe.Sizeof(f2))
	fmt.Printf("%5.2f\n", f2)
	fmt.Printf("Sprintf: %T\n", fmt.Sprintf("%5.2f\n", f2))
	fmt.Println(math.MaxFloat64) // float64最大值
	fmt.Println(math.MaxFloat32) // float32最大值

	var val1 = 1.00
	fmt.Println(val1 == 1.00)
}