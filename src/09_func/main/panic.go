package main


import (
	"fmt"
)

func test() (err error) {
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
func main() {
	fmt.Println("before main")
	err := test()

	fmt.Println("err=", err)
	fmt.Println("after main")
}