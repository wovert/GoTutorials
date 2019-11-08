package main

import "fmt"

func main() {
	t1 := 3.3 + 4.4i
	fmt.Println("t1 type is %T", t1)

	// 取得实部和虚部
	fmt.Println("read(t1) = ", real(t1), ", imag(t1) = ", imag(t1))
}