package main

import (
	"fmt"
	"unicode"
)

func main() {
	// 1. 判断字符串中汉字的数量
	// 难点是判断一个字符是汉字
	s := "Hello你好안녕"
	var count int
	for _, c := range s {
		if unicode.Is(unicode.Han, c) {
			count++
		}
	}
	fmt.Println(count)
}