package main

import (
	"fmt"
)

func test(x int) {
	result := 100 / x
	fmt.Println("result=",result)
}

func main() {
  // 延迟调用
  defer fmt.Println("延迟调用4")
  defer fmt.Println("延迟调用3")
  defer fmt.Println("延迟调用2")
  
  // 内存出问题
  defer test(0)
  fmt.Println("主函数被调用...")
  defer fmt.Println("延迟调用1")
}