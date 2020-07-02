package main

import "fmt"

func funcA() {
  fmt.Println("a")
}

func funcB() {
	// 刚刚打开数据库连接
	defer func() {
		err := recover() // 找到panic错误，并正常处理业务逻辑
		if err != nil {
			fmt.Println("释放数据库连接")
			fmt.Println(err)
		}
	}()
	// panic("出现了严重的出错啦！！！") // 程序崩溃退出
  fmt.Println("b")
}

func funcC() {
  fmt.Println("c")
}

func main() {
	funcA()
	funcB()
	funcC()
}