package main

import(
	"fmt"
	"crypto/md5"
)


func main() {
  Md5Inst := md5.New()
  Md5Inst.Write([]byte("lisi"))
  Result := Md5Inst.Sum([]byte(""))
  
  fmt.Printf("%x\n\n", Result)
}