package main

import (
"fmt"
"io"
"net/http"
"os"
"strconv"
)

/**
横向规律：
https://movie.douban.com/top250?start=0&filter=
https://movie.douban.com/top250?start=25&filter=
https://movie.douban.com/top250?start=50&filter=

纵向规律：
电影名称：<img width"100" alt="[电影名称]"  `<img width"100" alt="(?s:(.*?))"`
分数：<span class="rating_num" property="v:average">[分数]</span>  `<span class="rating_num" property="v:average">(?s:(.*?))</span>`
评分人数：<span>[评分人数]人评价</span>   `<span>(?s:(.*?))人评价</span>`
*/

var url = "https://movie.douban.com/top250?filter=&start=0" // 0,50,100,150...

func main() {
	var start, end int
	fmt.Printf("请输入爬取的起始页：")
	fmt.Scan(&start)
	fmt.Printf("请输入爬取的终止页：")
	fmt.Scan(&end)

	working(start, end)
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
func working(start, end int) {
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
