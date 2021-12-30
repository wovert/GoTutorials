package main

import "fmt"

func main() {
	var (
		achar        byte = 'A'
		aint         byte = 65
		unicodePoint rune = '中'
	)

	fmt.Println(achar, aint)
	fmt.Println(unicodePoint)

	fmt.Printf("%d %b %o %x %U %c %c %c\n", achar, 15, 15, 15, unicodePoint, unicodePoint, achar, 65)

}
