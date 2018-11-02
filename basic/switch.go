package main

import "fmt"

func main() {
	x := 0
	switch x {
		case 6:
			fmt.Print("星期六")
			break
		case 0:
			fmt.Print("星期日")
			break
		default:
			fmt.Print("周中")
	}
}