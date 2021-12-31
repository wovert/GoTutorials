package main

import "fmt"

func main() {
	stack := []string{}

	stack = append(stack, "a")
	stack = append(stack, "b")
	stack = append(stack, "c")

	// pop
	x := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Println("弹出:", x)
	fmt.Printf("queue=%#v\n", stack)

	x = stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Println("弹出:", x)
	fmt.Printf("queue=%#v\n", stack)

	x = stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Println("弹出:", x)
	fmt.Printf("queue=%#v\n", stack)
}
