package main

import (
	"encoding/json"
	"fmt"
)


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

  // create map
  m := make(map[string]interface{}, 4)

  err := json.Unmarshal([]byte(jsonBuf), &m)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

  fmt.Printf("json解码生成map m=%+v\n", m)


  // error
	//var str string
  //str = m["company"] // error
  //fmt.Println("str=", str)

	// 类型断言
	for key, value := range m {
		//fmt.Printf("%v=>%v\n", key, value)
		switch data := value.(type) {
			case string:
				str := data
				fmt.Printf("map[%s]的值类型为string, value=%s\n", key, str)
			case bool:
				fmt.Printf("map[%s]的值类型为bool, value=%v\n", key, data)
			case float64:
				fmt.Printf("map[%s]的值类型为float64, value=%v\n", key, data)
			case []string: // 不会执行
				fmt.Printf("map[%s]的值类型为[]string, value=%v\n", key, data)
			case []interface{}:
				fmt.Printf("map[%s]的值类型为[]interface, value=%v\n", key, data)
		}
	}
}