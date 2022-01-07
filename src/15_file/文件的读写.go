package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
)
func WriteFile(path string) {
	// 1. 创建文件
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("create err:", err)
		return
	} else {
		fmt.Println("create file successful!")
	}
	// 2. 使用完毕，需要关闭文件
	defer f.Close()

	// 3. 写入内容
	var buf string
	for i:=0; i<10; i++ {
		buf = fmt.Sprintf("i=%d\n", i)
		// fmt.Println("buf=", buf)

		_, err := f.WriteString(buf)
		if err != nil {
			fmt.Println("error=", err)
		}
		// fmt.Println("n=", n)
	}
}

func ReadFileLine(path string) {
	// 打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	// 关闭文件
	defer f.Close()

	// 新建一个缓冲区，把内容先放到缓冲区
	r := bufio.NewReader(f)

	for {
		// 遇到 \n 结束读取，但是 \n 也读取进入
		buf, err := r.ReadBytes('\n')
		if err != nil {
			// 文件已经结束
			if err == io.EOF {
				break
			}
			fmt.Println("err = ", err)
		}
		fmt.Printf("buf = %s#\n", string(buf))
	}

}

func ReadFile(path string) {
	// 打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	// 关闭文件
	defer f.Close()

	buf := make([]byte, 1024*2) // 2k大小

	// n 代表从文件读取内容的长度
	// 指定大小读取
	n, err1 := f.Read(buf)

	// 文件出错，同时没有到结尾
	if err1 != nil && err1 != io.EOF {
		fmt.Println("err1=", err1)
		return
	}
	fmt.Println("buf=", string(buf[:n]))
}

func main() {
	path := "./tmp"
	// WriteFile(path)
	ReadFile(path)
	ReadFileLine(path)
}