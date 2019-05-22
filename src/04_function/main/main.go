package main

import (
	"fmt"
)

type myFunType func(int, int) int

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

func myFun(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}

func myFun2(funvar myFunType, num1 int, num2 int) int {
	return funvar(num1, num2)
}

// 支持返回值命名
func getSumAndSub(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

func main() {
	type myInt int
	var num1 myInt
	var num2 int

	num1 = 40
	num2 = int(num1)
	fmt.Println("num1=", num1)
	fmt.Println("num2=", num2)

	res := myFun2(getSum, 500, 500)
	fmt.Println("res=", res)

	a, b := getSumAndSub(100, 50)
	fmt.Printf("a=%v, b=%v\n", a, b)
}