package main

import (
	"fmt"
	"strconv"
	"net/http"
	"io"
	"regexp"
	"os"
)

func SaveImg(idx int, url string, page chan int)  {
	path :="C:/itcast/img/" + strconv.Itoa(idx+1) + ".jpg"
	f, err := os.Create(path)
	if err != nil {
		fmt.Println(" http.Get err:", err)
		return
	}
	defer f.Close()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(" http.Get err:", err)
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		f.Write(buf[:n])
	}
	page <- idx
}

func main()  {
	url := "https://www.douyu.com/g_yz"

	// 爬取 整个页面，将整个页面全部信息，保存在result
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet err:", err)
		return
	}
	// 解析编译正则
	ret := regexp.MustCompile(`data-original="(?s:(.*?))"`)
	// 提取每一张图片的 url
	alls := ret.FindAllStringSubmatch(result, -1)

	page := make(chan int)
	n := len(alls)

	for idx, imgURL := range alls {
		//fmt.Println("imgURL:", imgURL[1])
		go SaveImg(idx, imgURL[1], page)
	}

	for i:=0; i<n; i++ {
		fmt.Printf("下载第 %d 张图片完成\n", <- page)
	}

}
// 获取一个网页所有的内容， result 返回
func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 4096)
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