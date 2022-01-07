package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company string `json:"-"` // 不会输出到屏幕
	Subjects []string `json:"subjects"` // 二次编码
	IsOk bool `json:"string"`
	Price float64 `json:"string"`
}

func main() {
  // json to struct
  jsonBuf := `
		{
			"company": "wovert",
			"subjects": [
				"Go",
				"C++",
				"Python",
				"Test"
			],
			"isok": true,
			"price": 666.66
		}
	`
  var tmp IT // 定义结构体变量
  err := json.Unmarshal([]byte(jsonBuf), &tmp)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

  //fmt.Println("tmp=",tmp)
  fmt.Printf("json解码生成结构体tmp=%+v\n", tmp)


	type Subject struct {
		Subjects []string `json:"subjects"`
	}

  var tmp2 Subject
	err2 := json.Unmarshal([]byte(jsonBuf), &tmp2)
	if err2 != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Printf("tmp2=%+v\n", tmp2)


}