package main

import (
  "fmt"
  "regexp"
)

func main() {
  str := `<!doctype html>
<html>
<head>
</head>
<body>
<h1>标题</h1>
<p>内容1</p>
<p>内容2
阿克苏地方
隋东风搜索</p>
<p>内容3</p>
</body>
</html>`

  // 解析，编译正则表达式
  ret := regexp.MustCompile(`<p>(?s:(.*?))</p>`) // 单行模式(模式修饰符)?s， .可以匹配换行符

  // 提取需要数据, 返回[][]string
  data := ret.FindAllStringSubmatch(str, -1)
  for _, one := range data {
    fmt.Println("one[0]=", one[0])
    fmt.Println("one[1]=", one[1])
  }

}