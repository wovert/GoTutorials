package main

import (
	"fmt"
	_ "../init" // 调用 init初始化函数
)

type FuncType func(int, int) int

func Add(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func Div(a, b int) int {
	return a / b
}

func Calc(a, b int, fTest FuncType) int {
	fmt.Println("calc")
	result := fTest(a, b)
	return result
}

func main() {
	res := Calc(10, 20, Add)
	fmt.Println("res=", res)

	res = Calc(10, 20, Minus)
	fmt.Println("res=", res)

	res = Calc(10, 20, Mul)
	fmt.Println("res=", res)

	res = Calc(10, 20, Div)
	fmt.Println("res=", res)
}