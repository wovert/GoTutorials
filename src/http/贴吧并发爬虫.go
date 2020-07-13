package main

import (
  "fmt"
  "io"
  "net/http"
  "os"
  "strconv"
)

/**
并发版百度爬虫：

1. 封装 爬取一个网页内容的 代码 到  SpiderPage（index）函数中

2. 在 working 函数 for 循环启动 go 程 调用 SpiderPage() —— > n个待爬取页面，对应n个go程

3. 为防止主 go 程提前结束，引入 channel 实现同步。 SpiderPage（index，channel）

4. 在SpiderPage() 结尾处（一个页面爬取完成）， 向channel中写内容 。 channel <- index

5.  在 working 函数 添加新 for 循环， 从 channel 不断的读取各个子 go 程写入的数据。 n个子go程 —— 写n次channel —— 读n次channel

 */

// loop every page 王者荣耀
var url = "https://tieba.baidu.com/f?kw=%E7%8E%8B%E8%80%85%E8%8D%A3%E8%80%80&ie=utf-8&cid=&tab=corearea&pn=" // 0,50,100,150...

func main() {
  var start, end int
  fmt.Printf("请输入爬取的起始页：")
  fmt.Scan(&start)
  fmt.Printf("请输入爬取的终止页：")
  fmt.Scan(&end)

  working2(start, end)
}

// 爬取单个页面的函数
func SpiderPage(i int, page chan int){
  url += strconv.Itoa((i-1)*50)
  result, err := HttpGet2(url)
  if err != nil {
    fmt.Println("HttpGet err:", err)
    return
  }
  //fmt.Println("result=", result)

  // 将读到的整网页数据，保存成一个文件
  f, err := os.Create("第" + strconv.Itoa(i) + "页" + ".html")
  if err != nil {
    fmt.Println("os.Create err:", err)
    return
  }

  f.WriteString(result) // 写入内容
  f.Close() // 保存一个文件，关闭一个文件。 defer关键字是整个函数执行之后才会执行

  page <- i // 与主 go 程完成同步
}


// 爬取页面操作
func working2(start, end int) {
  fmt.Printf("正在爬虫第%d页到%d页...\n", start, end)

  page := make(chan int)

  // 10个 子go程序写
  for i:=start; i<=end; i++ {
    go SpiderPage(i, page)
  }

  // 10个 子go程序读
  for i:=start; i<=end; i++ {
    fmt.Printf("第%d个页爬取完成\n", <-page)
  }

}

func HttpGet2(url string) (result string, err error) {
  resp, err1 := http.Get(url)
  if err1 != nil {
    err = err1 // 将封装函数内部的错误，传出给调用者
    return
  }
  defer resp.Body.Close()

  // 循环读取网页数据，传出给调用者
  buf := make([]byte, 4096)
  for {
    n, err2 := resp.Body.Read(buf)
    if n == 0 {
      fmt.Println("读取网页完成")
      break
    }
    if err2 != nil && err2 != io.EOF {
      fmt.Println("发送错误")
      err = err2
      return
    }
    result += string(buf[:n]) // 累加每一次循环的 buf 数据，存入 result 一次性返回
  }
  return
}
