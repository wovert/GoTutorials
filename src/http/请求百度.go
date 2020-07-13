package main

import (
	"fmt"
	"net/http"
  "io/ioutil"
  "os"
  "strings"
)

func errFunc(err error, info string) {
  if err != nil {
    fmt.Println(info, err)
    // return // 返回当前函数调用
    // runtime.Goexit() // 结束当前go程
    os.Exit(1) // 结束当前进程
  }
}

func main() {
//   host := "127.0.0.1:8080"

//   http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
// 	  w.Write([]byte("hell world"))
//   })

//   http.ListenAndServe(host, nil)

  res, err := http.Get("http://www.baidu.com")
  //if err != nil {
	//  panic(err)
  //}
  errFunc(err, "链接地址失败")

  defer res.Body.Close()
  
  // 获取body信息
  body, err := ioutil.ReadAll(res.Body)
  // fmt.Println(string(body))

  url := "http://www.baidu.com"
  dataType := "application/x-www-form-urlencoded"
  res, err = http.Post(url, dataType, strings.NewReader("id=1"))
  if err != nil {
    panic(err)
  }
  body, err = ioutil.ReadAll(res.Body)
  fmt.Println(string(body))
}