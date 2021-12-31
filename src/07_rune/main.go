package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	ascii := "abc中国人"
	fmt.Println([]byte(ascii))
	fmt.Println([]rune(ascii))
	fmt.Println(len(ascii))                    // 字节
	fmt.Println(len([]rune(ascii)))            // 字数
	fmt.Println(utf8.RuneCountInString(ascii)) // 字数

	fmt.Println(string([]byte(ascii)))
	fmt.Println(string([]rune(ascii)))
}
