package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "How do you do"
	s2 := strings.Split(s1, " ")
	m1 := make(map[string]int, 10)
	for _, w := range s2 {
		if _, ok := m1[w]; !ok {
			m1[w] = 1
		} else {
			m1[w]++
		}
	}
	for key, value := range m1 {
		fmt.Println(key, value)
	}
}