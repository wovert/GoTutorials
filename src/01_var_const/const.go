package main

import "fmt"

func main() {
	// 不允许没有初始化声明常量
	// const b
	const (
		A = "test"
		B // 使用前一个常量值进行初始化(B)
		C // 使用前一个常量值进行初始化(B=>C)
		D = "TestD"
		E
		F
	)
	fmt.Println(B, C)
	fmt.Println(E, F)

	const (
		Mon = iota // 在常量组内使用 iota 初始化 0, 每次调用+1
		Tue
		Wed
		Thu
		Fir
		Sat
		Sun
	)
	fmt.Println(Mon, Tue, Wed, Thu, Fir, Sat, Sun)
}
