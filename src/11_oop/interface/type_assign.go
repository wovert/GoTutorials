package main


import (
	"fmt"
)

func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if !ok {
		fmt.Println("猜错了")
	} else {
		fmt.Println("传进来的一个字符串：", str)
	}
}

func assign2(a interface{}) {
	fmt.Printf("%T\n", a)
	// 类型断言
	switch t := a.(type) {
		case string:
			fmt.Println("是一个字符串:", t)
		case int:
			fmt.Println("是一个int:", t)
		case int64:
			fmt.Println("是一个int64:", t)
		case bool:
			fmt.Println("是一个bool:", t)
	}
}


func main() {
  assign(100)
  assign2(100)
  assign2("是你")
  assign2(false)
}