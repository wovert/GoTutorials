package main

import (
	"fmt"
	"../errors"
)

func MyDiv(a, b int) (result int, err error) {
	err = nil
	if b == 0 {
		err = errors.New("被除数不能为0")
		// return result, err
		return
	}
	result = a / b
	return
}

func testa() {
	fmt.Println("aaaaaa")
}

func testb(x int) {
	fmt.Println("bbbbbb")
	// 显示调用 panic 函数，导致程序中断
	// panic("this is panic func.")

	defer func() {
		// recover() // 可以打印panic的错误信息		
		// 产生了panic异常
		if err := recover(); err != nil {
			fmt.Println(err)
		}

	}() // 调用匿名函数

	var a [10]int
	// 当x为20的时数组越界，产生一个Panic，导致程序崩溃
	a[x] = 111 // panic: runtime error: index out of range

}

func testc() {
	fmt.Println("cccccc")
}

func main() {
	err := fmt.Errorf("%s\n", "this is normal error")
	fmt.Println("err =",err)

	// 接口的使用
	err2 := errors.New("this is normal err2")
	fmt.Println("err2 =", err2)

	result, err := MyDiv(10, 0)
	if err != nil {
		fmt.Println("err =", err)
	} else {
		fmt.Println("result =", result)
	}

	// 致命错误（数组越界，内存溢出）

	testa()
	// testb(1)
	testb(20)
	testc() // 不会执行
}