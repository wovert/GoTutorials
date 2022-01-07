package main


import (
	"fmt"
	"strconv"
	// "strings"
)

func main() {
	// 转换为字符串后追加到字节数组
	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, true)
	// 追加的数，地上那个为指定10进制方式追加
	slice = strconv.AppendInt(slice, 1234, 10)
	slice = strconv.AppendQuote(slice, "abcd hello")

	fmt.Println("slice", string(slice)) // 转换string
	
	// 其他类型转换字符串
	var str string
	str = strconv.FormatBool(false)
	// f: 但因格式，以小数方式，-1指小数点位数（紧缩模式），64以float64处理
	str = strconv.FormatFloat(3.14, 'f', -1, 64)
	
	// 整型转字符串，常用
	str = strconv.Itoa(6666)

	fmt.Println("str=", str)

	// 字符串其他类型
	if flag, err := strconv.ParseBool("true"); err != nil {
		fmt.Println("err=", err)
	} else {
		fmt.Println("flag=", flag)
	}

	// 字符串转换为整型
	a, _ := strconv.Atoi("567")
	fmt.Println("a=", a)

	

}