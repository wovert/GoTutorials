package main

import (
	"fmt"
	_ "unsafe"
	"strconv"
)

func main() {
	basic2string_1()
	basic2string_2()
	basic2string_3()
}

// 方法1：基本数据类型转字符串
func basic2string_1() {
	var num1 int = 90
	var num2 float64 = 23.456
	var b bool = true
	// var mychar byte = 'h'
	var str string

	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%.2f", num2)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str=%q\n", str, str)	
}

// 方法2：基本数据类型转字符串
func basic2string_2() {
	var num3 int = 99
	var num4  float64 = 23.456
	var b2 bool = true
	var str string

	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.FormatBool(b2)
	fmt.Printf("str type %T str=%q\n", str, str)
}

// 方法3：基本数据类型转字符串
func basi2string_3() {
	var num1 = 4567
	var str string
	str = strconv.Itoa(int(num1))
	fmt.Printf("str type %T str=%q\n", str, str)
}