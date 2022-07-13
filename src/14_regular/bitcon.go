package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

// 爬取指定url 的页面，返回 result
func HttpGetDB(url string) (result string, err error) {

	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()

	buf := make([]byte, 4096)
	// 循环爬取整页数据
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		result += string(buf[:n])
	}
	return
}

func Save2file(filmName, filmScore, peopleNum [][]string) {
	path := "第 "
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("os.Create err:", err)
		return
	}
	defer f.Close()

	n := len(filmName) // 得到 条目数。 应该是 25
	// 先打印 抬头  电影名称        评分         评分人数
	f.WriteString("电影名称" + "\t\t\t" + "评分" + "\t\t" + "评分人数" + "\n")
	for i := 0; i < n; i++ {
		f.WriteString(filmName[i][1] + "\t\t\t" + filmScore[i][1] + "\t\t" + peopleNum[i][1] + "\n")
	}
}

// 爬取页面数据信息
func SpiderPageDB(quit chan int) {
	// 获取 url
	url := "https://price.btcfans.com/"

	// 封装 HttpGet2 爬取 url 对应页面
	result, err := HttpGetDB(url)
	if err != nil {
		fmt.Println("HttpGet2 err:", err)
		return
	}
	fmt.Println("result=", result)

	// 解析，编译正则表达式
	ret := regexp.MustCompile(`<table class="highlight borderNone"><colgroup>(?s:(.*?))</colgroup><tbody>(?s:(.*?))</tbody></table>`) // 单行模式(模式修饰符)?s， .可以匹配换行符

	// 提取需要数据, 返回[][]string
	data := ret.FindAllStringSubmatch(result, -1)
	for _, one := range data {
		fmt.Println("one[0]=", one[0])
		fmt.Println("one[1]=", one[1])
	}

	// // 解析、编译正则表达式 —— 电影名称：
	// ret1 := regexp.MustCompile(`<img width="100" alt="(?s:(.*?))"`)
	// // 提取需要信息
	// filmName := ret1.FindAllStringSubmatch(result, -1)

	// // 解析、编译正则表达式 —— 分数：
	// pattern := `<span class="rating_num" property="v:average">(?s:(.*?))</span>`
	// ret2 := regexp.MustCompile(pattern)
	// // 提取需要信息
	// filmScore := ret2.FindAllStringSubmatch(result, -1)

	// // 解析、编译正则表达式 —— 评分人数：
	// ret3 := regexp.MustCompile(`<span>(?s:(\d*?))人评价</span>`)
	// //ret3 := regexp.MustCompile(`<span>(.*?)人评价</span>`)

	// // 提取需要信息
	// peopleNum := ret3.FindAllStringSubmatch(result, -1)

	// // 将提取的有用信息，封装到文件中。
	// Save2file(filmName, filmScore, peopleNum)

	// // 与主go程配合 完成同步
	// quit <- 1
}

func toWork() {
	quit := make(chan int) //防止主go 程提前结束

	go SpiderPageDB(quit)

	fmt.Printf("爬取完毕\n", <-quit)
}

func main() {
	toWork()
}
