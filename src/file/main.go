package main

import (
	"fmt"
	"os"
)

func main() {
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
	}

  fmt.Println("Create file successful!")

}
