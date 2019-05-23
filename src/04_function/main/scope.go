package main

import (
	"fmt"
)

var Age int = 20
// Name := "Tom" // error: 等价于 
var Name string
// Name = "Jerry"

func init() {
	Name = "Jimy"
}

func main() {
	fmt.Println("Age is ", Age)
	fmt.Println("Name is ", Name)
}