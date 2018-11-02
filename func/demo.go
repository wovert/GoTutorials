package main

import "fmt"

func swap(a int, b int) (int,int) {
	return b,a
}

func add(a *int) *int {
	*a = *a + 100
	return a
}
func main() {
	a := 1
	b := 2
	add(&a)
	a, b = swap(a, b)
	fmt.Printf("a=%d b=%d", a, b)
}