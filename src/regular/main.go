package main

import (
	"fmt"
	"regexp"
)

func main() {
  isok, _ := regexp.Match("[a-zA-Z]{3}", []byte("cgw"))
  fmt.Printf("%v\n", isok)
  
  isok2, _ := regexp.MatchString("[a-zA-Z]{3}", "cw1ab2")
  fmt.Printf("%v\n", isok2)
  
  // 成功返回解析器，返回结构体
  reg := regexp.MustCompile(`a.c`)
  if reg == nil {
    fmt.Println("err = ", reg)
    return
  }

  buf := "abc azc a7c aac 888 a9c tac 6.66 7.66 20.2039"
  // 返回匹配信息作为切片， -1:匹配所有
  res := reg.FindAllStringSubmatch(buf, -1) // [[abc] [azc] [a7c] [aac] [a9c]]
  fmt.Printf("res =%T\n", res)
  fmt.Println("res = ", res)
  
  // 提取小数
  reg2 := regexp.MustCompile(`\d+\.{1}\d+`)
  if reg2 == nil {
    panic(reg2)
  }
  result := reg2.FindAllStringSubmatch(buf, -1) // -1:匹配全部， 1:匹配第一个，2:匹配前两个
  fmt.Println("result = ", result)

  // 提取网页内容
  buf = `<div>
    <h1>标题</h1>
    <span>你好</span>
    </div>`
  reg3 := regexp.MustCompile(`<div>(?s:(.*?))</div>`) // 模式：?s: 单行模式(匹配\n)，x>=0次匹配
  if reg3 == nil {
    panic(reg3)
  }
  res = reg3.FindAllStringSubmatch(buf, -1)
  for _, one := range res {
    fmt.Println("one[0]=", one[0])
    fmt.Println("one[1]=", one[1])
  }
}