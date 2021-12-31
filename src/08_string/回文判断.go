package main

import (
	"fmt"
)

func main() {
	// 上海自来水来自海上 s[0] s[len(s)-1]
	// 山西运煤车煤运西山
	// 黄山落叶松叶落山黄
	s := "上海自来水来自海上"
	// 把字符串中的字符放到一个[]rune中
	r := make([]rune, 0, len(s))
	for _, c := range s {
		r = append(r, c)
	}
	fmt.Println("[]rune:", r) // 中文ASCII没有，所以用rune,utf-8编码
	for i := 0; i < len(r)/2; i++ {
		// 上s[0] s[len(s)-1]

		// 海s[0] s[len(s)-2]
		// 自s[0] s[len(s)-3]
		// 来s[0] s[len(s)-4]
		if r[i] != r[len(r)-1-i] {
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")
}

