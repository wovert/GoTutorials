package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//test()
  //testWrite()
	testRead()
}

func test() {
	//f, err := os.Create("./tmp")
	//f, err := os.Open("./tmp")
	// 打开文件权限：O_RDONLY, O_WRONLY, O_RDWR
	f, err := os.OpenFile("./tmp", os.O_RDWR, 6)
	if err != nil {
		fmt.Println("create err:", err)
	}

	defer f.Close()

	_, err = f.WriteString("hello world")
	if err != nil {
		fmt.Println("write err:", err)
		return
	}

	fmt.Println("Create file successful!")
}

func testWrite() {
	// 写文件操作
	f, err := os.OpenFile("./tmp", os.O_RDWR, 6)
	if err != nil {
		fmt.Println("OpenFile err:", err)
		return
	}
	defer f.Close()

	fmt.Println("successful!")

	// 按照字符串写入，返回字节数
	//n, err := f.WriteString("12中\n") // 写入前三个123
	//if err != nil {
	//	fmt.Println("WriteString err:", err)
	//	return
	//}
	//fmt.Println("n =", n) // 写入字节数

	// 按位置写入
	// 偏移量(正：像文件尾偏，负：向文件头偏)
	// 偏移起始位置（os.SeekStart, os.SeekEnd, os.SeekCurrent）
	// 返回值：从文件起始位置，到当前文件读写指针位置的偏移量
	off, _ := f.Seek(5, io.SeekStart)
	//off, _ := f.Seek(-5, io.SeekEnd)
	fmt.Println("off :", off)
	n, _ := f.WriteAt([]byte("345"), off) // 从偏移off开始写入
	fmt.Println("Write n :", n) // 3

}

func testRead() {
	// 写文件操作
	f, err := os.OpenFile("./tmp", os.O_RDWR, 6)
	if err != nil {
		fmt.Println("OpenFile err:", err)
		return
	}
	defer f.Close()

	fmt.Println("successful!")

	// 创建一个带有缓冲区的 reader
	reader := bufio.NewReader(f)
	for {
		buf, err := reader.ReadBytes('\n') // 读一行数据
		if err != nil && err == io.EOF{
			fmt.Println(buf) // 默认读取ASCII码
			fmt.Println(string(buf))
			fmt.Println("文件读取完毕")
			return
		} else if err != nil {
			fmt.Println("ReadBytes err:", err)
			return
		}
		fmt.Println(buf) // 默认读取ASCII码
		fmt.Println(string(buf))
	}

}