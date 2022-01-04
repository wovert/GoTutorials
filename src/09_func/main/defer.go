package main

import "fmt"

func f() {
	fmt.Println("start func")

	// 先进后出执行 defer
	defer fmt.Println("延迟执行语句1") // 函数返回之前执行
	defer func() {
		fmt.Println("延迟执行语句2") // 函数返回之前执行
	}()
	defer fmt.Println("延迟执行语句3") // 函数返回之前执行
	fmt.Println("end func")
}

func f1() int {
	// 1. return 5
	// 2. x=6
	// 3. 
	x := 5
	// 函数体不管是否发生错误，都会执行 defer 语句（文件关闭等）
	defer func() {
		x++ // 修改的是 x 不是返回值
	}()
	return x
}

func f2() (x int) {
	// 1. return x=5
	// 2. x=6
	// 3. RET 6
	defer func() {
		x++
	}()
	return 5 // 返回值=x
}

func f3() (y int) {
	// 1. return y=x=5
	// 2. defer x=6
	// 3. return y=5
	x := 5
	defer func() {
		x++
	}()
	return x // 返回值 = x = y = 5
}

func f4()(x int) {
	defer func(x int) {
		x++ // 改变是函数中的副本
	}(x)
	return 5 // 返回值 = x = 5
}

func f5()(x int) {
	defer func(x int) int {
		x++ // 改变是函数中的副本
		return x
	}(x)
	return 5 // 返回值 = x = 5
}

func f6()(x int) {
	defer func(x *int) {
		(*x)++ // 改变是函数中的副本
	}(&x)
	return 5 // 1. 返回值 = x = 5, 2. defer x=6, 3. RET返回
}

func main() {
	f()
	fmt.Println(f1())
	// fmt.Println(f2())
	// fmt.Println(f3())
	// fmt.Println(f4())
	// fmt.Println(f5())
	// fmt.Println(f6())
}