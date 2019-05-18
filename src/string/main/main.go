package main

import (
	"fmt"
	"strings"
	"strconv"
)

func urlProcess(url string) string {
	result := strings.HasPrefix(url, "http://")
	if !result {
		url = fmt.Sprintf("http://%s", url)
	}
	return url
}

func pathProcess(path string) string {
	result := strings.HasSuffix(path, "/")
	if !result {
		path = fmt.Sprintf("%s/", path)
	}
	return path
}

func main() {
	var str = "hello world\n"
	var str2 = `hello \n` // 原样输出

	var a int
	var b bool
	var c byte = 'a' // 字符类型

	fmt.Println(str)
	fmt.Println(str2)

	fmt.Printf("%v\n", a) // 0
	fmt.Printf("%v\n", b) // false
	fmt.Printf("%v\n", c) // 97

	fmt.Printf("%t\n", a) // int=0
	fmt.Printf("%t\n", b) // false
	fmt.Printf("%t\n", c) // uint8=97

	fmt.Printf("%b\n", 1000) // 1111101000
	fmt.Printf("%e\n", 10222200.203023) // 1.022220e+07
	fmt.Printf("%q\n", "This is content") // "This is content"
	fmt.Printf("%x,%X\n", 16, 17) // 10, 11
	fmt.Printf("%p\n", &a) // 0xc000048058

	s := fmt.Sprintf("a=%d\n", 100)
	fmt.Printf("%q\n", s)

	s1 := "hello"
	s2 := " world"
	fmt.Println(s1 + s2)
	fmt.Printf("len(str1+str2)=%d\n", len(s1+s2))
	


	var (
		url string
		path string
	)
	fmt.Scanf("%s,%s", &url, &path)
	var ss1 string = urlProcess(url)
	var ss2 string = pathProcess(path)
	fmt.Printf("ss1=%s, ss2=%s\n", ss1, ss2)

	// -1 strings.Index(s, str) int
	fmt.Printf("%d\n", strings.Index(url, "w"))
	fmt.Printf("%d\n", strings.LastIndex(url, "w"))

	sss1 := strconv.Itoa(1000)
	fmt.Println("itoa:", sss1)

	number, err := strconv.Atoi(sss1)
	if err != nil {
		fmt.Println("can not conver to int", err )
	}

	fmt.Println("number:", number)

	i, err := strconv.Atoi("-42")
	if err != nil {
		fmt.Println("can not conver to int", err )
	}
	fmt.Printf("i=%d\n", i) // -42



}


