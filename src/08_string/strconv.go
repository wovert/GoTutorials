package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 字符串修改
	s2 := "白萝卜"
	s3 := []rune(s2) // 字符串强制转换成一个rune切片
	s3[0] = '红' // rune(int32)
	fmt.Println(string(s3)) // rune切片强制转换成字符串

	c1 := "红"
	c2 := '红'

	fmt.Printf("c1:%T c2:%T\n", c1, c2)

	c3 := "H" // string
	c4 := byte('H') // byte(uint8)
	fmt.Printf("c3:%T c4:%T\n", c3, c4)
	fmt.Printf("%d\n", c4)

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