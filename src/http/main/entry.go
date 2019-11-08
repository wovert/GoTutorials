package main

import (
	"fmt"
	"net/http"
  "io/ioutil"
  "strings"
)

func main() {
//   host := "127.0.0.1:8080"

//   http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
// 	  w.Write([]byte("hell world"))
//   })

//   http.ListenAndServe(host, nil)

  res, err := http.Get("http://www.baidu.com")
  if err != nil {
	  panic(err)
  }

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