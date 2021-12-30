package main

import (
	"fmt"
	"strconv"
)

func main() {
	// int32 => float32
	var i32Val = 257
	var f32Val = float32(i32Val)
	fmt.Printf("int32(%#v) 转换 float32(%#v)\n", i32Val, f32Val)

	// int32 => int8
	var i8Val = int8(i32Val)
	fmt.Printf("int32(%#v) 转换 int8(%#v)\n", i32Val, i8Val)

	// string => float64
	var sFloat64 = "3.14159265926"
	f64Val, _ := strconv.ParseFloat(sFloat64, 64)
	fmt.Printf("string(%#v) 转换 float64(%#v)\n", sFloat64, f64Val)

	// float64 => string
	// 将 64 位浮点型的数字转换为字符串，其中 fmt 表示格式（其值可以是 'b'、'e'、'f' 或 'g'），prec 表示精度，bitSize 则使用 32 表示 float32，用 64 表示 float64。
	f64_2_string := strconv.FormatFloat(f64Val, 'f', 2, 64)
	fmt.Printf("float64(%#v) 转换 string(%#v)\n", f64Val, f64_2_string)

	// string => uint64
	var sUI64Val = "255"
	ui64Val, _ := strconv.ParseUint(sUI64Val, 10, 64) // base 10, up to 64 bits
	fmt.Printf("string(%#v) 转换 uint64(%d)\n", sUI64Val, ui64Val)

	// uint64 => string
	ui64_2_string := strconv.FormatUint(ui64Val, 10)
	fmt.Printf("uint64(%d) 转换 string(%#v)\n", ui64Val, ui64_2_string)

	// 	string => int
	s_2_int, _ := strconv.Atoi(ui64_2_string)
	fmt.Printf("string(%#v) 转换 int(%#v)\n", ui64_2_string, s_2_int)

	// 	int => string
	i_2_string := strconv.Itoa(s_2_int)
	fmt.Printf("int(%#v) 转换 string(%#v)\n", s_2_int, i_2_string)
}

func demo() {
	var n1 int32 = 12
	var n2 int64
	var n3 int8

	n2 = int64(n1) + 20
	n3 = int8(n1) + 20
	fmt.Println("n2=", n2, "n3=", n3)
}

func demo2() {
	var n1 int32 = 12
	var n3 int8
	var n4 int8
	n4 = int8(n1) + 127 // 编译通过，但是溢出
	// n3 = int8(n1) + 128 // 编译不通过（int8:127~-128）
	fmt.Println(n3, n4)
}
