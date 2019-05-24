package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 字符串长度
	str := "1a中"
	fmt.Printf("str length is %v\n", len(str))

	for i := 0; i < len(str); i++ {
		fmt.Printf("str[%v]=%c\n", i, str[i])
	}
	fmt.Println("-----------------")
	str1 := []rune(str)
	for i := 0; i < len(str1); i++ {
		fmt.Printf("str1[%v]=%c\n", i, str1[i])
	}

	// 字符串转整数
	n, err := strconv.Atoi("12")
	if err != nil {
		fmt.Println("转换错误\n", err)
	} else {
		fmt.Println("转换结果是\n", n)
	}
	
	// 整数转字符串：
	str3 := strconv.Itoa(1234)
	fmt.Printf("str=%v整数转字符%T\n", str3, str3)

	// 字符串转 []byte 
	var bytes = []byte("hello World")
	fmt.Printf("bytes=%v\n", bytes)

	// []byte转字符串
	str4 := string([]byte{97,98,99})
	fmt.Printf("str=%s\n", str4)

	// 10进制转2,8,16进制
	str5 := strconv.FormatInt(123, 2)
	fmt.Printf("123对应的二进制是=%v\n", str5)
	str6 := strconv.FormatInt(123, 16)
	fmt.Printf("123对应的16进制是=%v\n", str6)

	// 查找子串是否在指定的字符串中
	isContain := strings.Contains("seafood", "foo")
	fmt.Printf("foo在seafood字符%v\n", isContain)
}
