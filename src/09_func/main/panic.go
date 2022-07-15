package main

import (
	"fmt"
)

func testPanic() (err error) {
	defer func() {
		fmt.Println("defer")
		// 获取panic错误信息
		if panicError := recover(); panicError != nil {
			fmt.Printf("%T, %#v\n", panicError, panicError)
			err = fmt.Errorf("%s", panicError)
		}
	}()
	fmt.Println("before")
	panic("自定义panic") // 终止程序，终止之前执行 defer 语句
	fmt.Println("after")
	return
}

func demo(i int) {
	// recover 错误兰姐 出现在 panic 错误之前
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	var arr [10]int
	arr[i] = 123 // err
	fmt.Println(arr)
}
func main01() {
	fmt.Println("before main")
	err := testPanic()

	fmt.Println("err=", err)
	fmt.Println("after main")
}
func main02() {
	fmt.Println("before main")
	demo(10)
	fmt.Println("after main")
}

func main() {
	//main01()
	main02()
}
