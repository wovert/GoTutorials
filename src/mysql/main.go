package main

import (
	"fmt"
	"database/sql" // 系统包
	_ "github.com/go-sql/driver/mysql"
)

func main() {
  db, err := sql.Open("mysql", "root:Am2r1c@n*8cg@tcp(127.0.0.1:3306)/test?charset=utf8")
  if err != nil {
	  panic(err)
  }
  
}