package main

import (
  "fmt"
  "regexp"
)

func main() {
  str := "abc a7c cat mfc 8cs azc cba"

  // 解析，编译正则表达式
  //ret := regexp.MustCompile(`a.c`) // 参数s正则表达式，返回结构体，``表示原生字符串
  ret := regexp.MustCompile(`a[0-9]c`) // 参数s正则表达式，返回结构体，``表示原生字符串


  // 提取需要数据, 返回[][]string
  data := ret.FindAllStringSubmatch(str, -1)
  fmt.Printf("data=%s", data) // [[abc] [a7c] [azc]]

}