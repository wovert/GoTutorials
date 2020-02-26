package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//test()
  testWrite()
	//testRead()
	//copyFile()
	//readDir()
}

func test() {
	// 创建文件
	//f, err := os.Create("./tmp")
	//if err != nil {
	//	fmt.Println("create err:", err)
	//	return
	//} else {
	//	fmt.Println("create file successful!")
	//}
	//defer f.Close()
	//return

	// 只读方式打开文件
	//f, err := os.Open("./tmp")
	//if err != nil {
	//	fmt.Println("Open file error!")
	//} else {
	//	fmt.Println("Open file successful!")
	//}
	//defer f.Close()
	//return

	// 读写方式打开文件
	// 打开文件权限：O_RDONLY, O_WRONLY, O_RDWR
	// 0: 没有任何权限；1: 执行权限（可执行文件); 2: 写权限; 3: 写/执行
	// 4: 读权限; 5: 去/执行；6: 读权限于写权限， 7: 读权限于写权限，执行权限
	f, err := os.OpenFile("./tmp", os.O_RDWR, 6)
	if err != nil {
		fmt.Println("create err:", err)
	}
	defer f.Close()

	// 写入字符串内容
	_, err = f.WriteString("hello world")
	if err != nil {
		fmt.Println("write err:", err)
	} else {
		fmt.Println("Create file and Write successful!")
	}
}

func testWrite() {
	// 写文件操作
	f, err := os.OpenFile("./tmp", os.O_RDWR, 6)
	if err != nil {
		fmt.Println("OpenFile err:", err)
		return
	}
	defer f.Close()

	fmt.Println("Open file successful by Read Write Only")

	// 按照字符串写入，返回字节数
	//n, err := f.WriteString("12中\n") // 写入前三个123，中文3个字节
	//if err != nil {
	//	fmt.Println("WriteString err:", err)
	//	return
	//}
	//fmt.Println("写入字节数为 n =", n) // 写入字节数


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

func readDir() {
	fmt.Println("请输入查询的目录：")
	var path string
	fmt.Scan(&path)

	// open dir
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("Open file err:", err)
		return
	}
	defer f.Close()

	// 读取目录项
	info, err := f.Readdir(-1) // 读取目录中所有目录项

	// 遍历切片
	for _, fileInfo := range info {
		if fileInfo.IsDir() {
			fmt.Println(fileInfo.Name() , "是一个目录")
		} else {
			fmt.Println(fileInfo.Name() , "是一个文件")
		}
	}

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

	// 创建一个带有缓冲区（用户缓冲）的 reader
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

func copyFile() {
	args := os.Args // 获取命令行参数

	if args == nil || len(args) !=3 {
		fmt.Println("useage: xx src File destFile")
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

	buf := make([]byte, 1024) // 切片缓冲区
	for {
		// 从源文件读取内容，n为读取文件内容的长度
		n, err := srcFile.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			break;
		}

		if n == 0 {
			fmt.Println("文件处理完毕")
			break
		}

		// 切片截取
		tmp := buf[:n]

		// 把读取的内容写入到目的文件
		dstFile.Write(tmp)

	}
	srcFile.Close()
	dstFile.Close()
}