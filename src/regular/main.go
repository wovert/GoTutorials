package main

import (
	"fmt"
	"regexp"
)

func main() {
  isok, _ := regexp.Match("[a-zA-Z]{3}", []byte("cgw"))
  fmt.Printf("%v\n", isok)

  isok2, _ := regexp.MatchString("[a-zA-Z]{3}", "cw1ab2")
  fmt.Printf("%v\n", isok2)

  reg := regexp.
}