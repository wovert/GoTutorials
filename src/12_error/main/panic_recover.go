package main

import (
	"fmt"
	"errors"
)

func test() {
	// defer + recover 来捕获处理异常
	defer func() {
		// recover()内置函数，可以捕获到异常
		if err := recover(); err != nil {
			fmt.Println("err=", err)	
			fmt.Println("发送邮件给admin@amdin.com")
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

// 读取配置文件的init.conf 的信息
// 文件传入错误，返回自定义错误
func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("读取文件错误")
	}
}

func test2() {
	err := readConf("config.ini0")
	if err != nil {
		// 读取文件发送错误，输出错误，并终止程序
		panic(err)
	}

	fmt.Println("test2继续执行剩余代码...")
}

func main() {
	test()
	fmt.Println("测试结束")
	// test2()
}