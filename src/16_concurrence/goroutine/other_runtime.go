package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("GOROOT=", runtime.GOROOT())
	fmt.Println("version=", runtime.Version())
	fmt.Println("cpu=", runtime.NumCPU())
	fmt.Println("arch=", runtime.GOARCH)
}