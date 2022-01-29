package main

import (
	"fmt"
	"os"
	"io"
)

func main() {
	path := "./tmp"
	copyFile()
}

func copyFile() {
	args := os.Args // 获取命令行参数

	if args == nil || len(args) !=3 {
		fmt.Println("usage: xx src File destFile")
		return
	}

	srcPath := args[1]
	dstPath := args[2]
	fmt.Printf("srcPath=%s, dstPath=%s\n", srcPath, dstPath)

	if srcPath == dstPath {
		fmt.Println("errors: 源文件名 与 目标文件名相同")
		return
	}

	srcFile, err1 := os.Open(srcPath) // 打开源文件
	if err1 != nil {
		fmt.Println(err1)
		return
	}

	dstFile, err2 := os.Create(dstPath) // 创建目标文件
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	// 从读文件中获取数据，放到缓冲区
	buf := make([]byte, 1024) // 切片缓冲区，一次读取1024字节

	// 循环从读文件中，获取数据，原封不动的写到写文件中
	for {
		// 从源文件读取内容，n为读取文件内容的长度, 读取出来的有效字符(byte)数，读取的字节数
		n, err := srcFile.Read(buf) // n 从文件读取内容的长度
		if err != nil && err != io.EOF { // 文件出错，同时没有到结尾
			fmt.Println(err)
			break;
		}

		if n == 0 {
			fmt.Println("文件处理完毕")
			break
		}

		// 切片截取，从缓冲区中读取放入的字节数
		tmp := buf[:n]

		// 把读取的内容写入到目的文件
		dstFile.Write(tmp) // 读取多少，写多少

	}
	srcFile.Close()
	dstFile.Close()
}
