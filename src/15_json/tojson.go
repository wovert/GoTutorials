package main

import(
	"fmt"
	"encoding/json"
)

type Student struct {
	name string // name = private 不显示在json
	age int // age = private
	Nickname string `json:"nickanme"` // public, 结构元素标签
	Height int `json:"height"` // public
}

func main() {
  // 对数组类型的json编码
  x := [5]int{1,2,3,4,5}
  s, err := json.Marshal(x)
  if err != nil {
    panic(err)
  }
  fmt.Println("数组类型的json编码:", string(s))

  // 对map类型的json编码
  m := make(map[string]interface{})
  m["zhangsan"] = 100.4
  m["lisi"] = 20.1
  m["subjects"] = []string{"Go", "C++", "Rust"}
  m["isok"] = true
  m2, err2 := json.Marshal(m)
  if nil != err2 {
    panic(err2)
  }
  fmt.Println("map编码生成json:", string(m2))

  // 对结构进行json编码
  stu := Student{"张三",24,"昵称", 30}
  // m3, err3 := json.Marshal(stu)
  m3, err3 := json.MarshalIndent(stu, "", " ") // 格式化编码
  if nil != err3 {
    panic(err3)
  }
  fmt.Println("结构生成json:", string(m3))

  // 对m3进行解码
  var m4 interface{}
  json.Unmarshal([]byte(m3), &m4)
  fmt.Printf("json解码生成map=%v", m4)

  // json解析到结构体

  // json解析道map
}