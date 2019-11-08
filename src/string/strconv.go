package main

import (
	"fmt"
	"strconv"
)

func main() {
  slice := make([]byte, 0, 1024)

  // 追加布尔类型 true
  slice = strconv.AppendBool(slice, true)

  // 10进制数追加1234
  slice = strconv.AppendInt(slice, 1234, 10)

  // 追加字符串
  slice = strconv.AppendQuote(slice, "hello world")

  // 转换 string
  fmt.Println("slice = ", string(slice))

  // 其他类型转换字符串
  var str string
  str = strconv.FormatBool(false)

  // ‘f' 指打印格式，以小数方式； -1指小数点位数（紧缩模式），64以float64
  str = strconv.FormatFloat(3.147, 'f', -1, 64)
  
  str = strconv.Itoa(6666)
  str = strconv.FormatInt(-30, 10)
  str = strconv.FormatUint(2000, 10)

  fmt.Println("str =", str)

  // 字符串转其他类型
  flag, err := strconv.ParseBool("true")
  if err != nil {
	fmt.Println("flag =", flag)
  } else {
	fmt.Println(" err =", err)
  }

  // 字符串转换整型
  a, _ := strconv.Atoi("456")
  fmt.Println(" a =", a)


}