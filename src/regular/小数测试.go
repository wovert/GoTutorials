package main

import (
  "fmt"
  "regexp"
)

func main() {
  str := "3.14 123.23 .62 ha 1.0 abc 7. abc.3 33.3 133."

  // 解析，编译正则表达式
  //ret := regexp.MustCompile(`[0-9]+\.[0-9]+`) // 参数s正则表达式，返回结构体，``表示原生字符串
  ret := regexp.MustCompile(`\d+\.\d+`) // 参数s正则表达式，返回结构体，``表示原生字符串


  // 提取需要数据, 返回[][]string
  data := ret.FindAllStringSubmatch(str, -1)
  fmt.Printf("data=%s", data) // [[abc] [a7c] [azc]]

}