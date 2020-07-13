package main

import (
  "fmt"
  "io"
  "net/http"
  "os"
  "strconv"
)
/**

1. 明确目标 Url
2. 发送请求，获取应答数据包。 http.Get(url)
3. 过滤 数据。提取有用信息。
4. 使用、分析得到数据信息。

https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=0		下一页 +50

https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=50

https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=100

https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=150

https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=300

1. 提示用户指定 起始、终止页。 创建working函数

2. 使用 start、end 循环 爬取每一页数据

3. 获取 每一页的 URL —— 下一页 = 前一页 + 50

4. 封装、实现 HttpGet() 函数，爬取一个网页的数据内容，通过 result 返回。

http.Get/ resp.Body.Close/ buf := make（4096）/ for { resp.Body.Read(buf)/   result += string(buf[:n])  return

5. 创建 .html 文件。 使用循环因子 i 命名。

6. 将 result 写入 文件 f.WriteString（result）。   f.close()  不推荐使用 defer 。
 */

func main() {
  var start, end int
  fmt.Printf("请输入爬取的起始页：")
  fmt.Scan(&start)
  fmt.Printf("请输入爬取的终止页：")
  fmt.Scan(&end)

  working(start, end)
}

// 爬取页面操作
func working(start, end int) {
  fmt.Printf("正在爬虫第%d页到%d页...\n", start, end)

  // loop every page 王者荣耀
  url := "https://tieba.baidu.com/f?kw=%E7%8E%8B%E8%80%85%E8%8D%A3%E8%80%80&ie=utf-8&cid=&tab=corearea&pn=" // 0,50,100,150...
  for i:=start; i<=end; i++ {
    url += strconv.Itoa((i-1)*50)
    result, err := HttpGet(url)
    if err != nil {
      fmt.Println("HttpGet err:", err)
      continue
    }
    //fmt.Println("result=", result)

    // 将读到的整网页数据，保存成一个文件
    f, err := os.Create("第" + strconv.Itoa(i) + "页" + ".html")
    if err != nil {
      fmt.Println("os.Create err:", err)
      continue
    }

    f.WriteString(result) // 写入内容
    f.Close() // 保存一个文件，关闭一个文件。 defer关键字是整个函数执行之后才会执行
  }
}

func HttpGet(url string) (result string, err error) {
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
