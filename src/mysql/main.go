package main

import (
	"fmt"
	"database/sql" // 系统包
	_ "github.com/go-sql-driver/mysql"
)

func main() {
  db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test?charset=utf8")
  if err != nil {
	panic(err)
  }
  fmt.Println("数据库连接成功！")
//   fmt.Printf("%v", db)

  // 插入数据
//   sql, err1 := db.Prepare("INSERT user_info SET nickname=?,departname=?")
//   if err1 != nil {
// 	fmt.Println("SQL预加载错误")
// 	panic(err1)
//   }
//   res, sqlErr := sql.Exec("李四", "技术部")
//   if sqlErr != nil {
// 	  fmt.Println("执行SQL语句错误")
// 	  panic(sqlErr)
//   }
//   id, err2 := res.LastInsertId()
//   if err2 != nil {
// 	  fmt.Println("数据返回错误")
// 	  panic(err2)
//   }
//   fmt.Println("id=", id)

  // 更新数据
//   sql, err1 := db.Prepare("UPDATE user_info SET nickname=? where id=?")
//   if err1 != nil {
//     fmt.Println("sql错误")
//   }
//   id := 1
//   res, err2 := sql.Exec("张三", id)
//   if err2 != nil {
// 	panic(err2)
//   }
//   affect, err3 := res.RowsAffected() // 影响的行数
//   if err3 != nil {
// 	panic(err3)
//   }

//   fmt.Printf("%v\n", affect)

  // 查询
//   rows, err1 := db.Query("SELECT * FROM user_info WHERE id=1")
//   rows, err1 := db.Query("SELECT * FROM user_info")
//   if err1 != nil {
// 	  panic(err1)
//   }
//   for rows.Next() {
// 	var id int
// 	var nickname string
// 	var department string
// 	var create_at string
// 	_ = rows.Scan(&id, &nickname, &department, &create_at)
// 	fmt.Println(id, nickname, department, create_at)
//   }

  // 删除
  sql, err1 := db.Prepare("delete from user_info where id=?")
  if err1 != nil {
	panic(err1)
  }
  id := 1
  res, err2 := sql.Exec(id)
  if err2 != nil {
	panic(err2)
  }
  fmt.Println("res=%v\n", res)

}