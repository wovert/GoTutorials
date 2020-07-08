package main

import (
	"fmt"
)

// ini 配置文件解析器

type MysqlConfig struct {
	Host string `ini:"host"`
	Port int `ini:"port"`
	User string `ini:"user"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

type RedisConfig struct {
	Host string `ini:"host"`
	Port int `ini:"port"`
	Password string `ini:"password"`	
	Database string `ini:"database"`
}

func loadIni(fileName string, data interface{}) (err error) {
	// 0. 参数校验
	// 0.1 data参数必须是指针类型（因为在函数中对其赋值）
	// 0.2 data参数必须是结构体类型指针（因为配置文件中键值对儿操作）
	// 1. 读文件得到字节类型数据
	// 2. 一行一行得读数据
	// 2.1 注释就跳过
	// 2.2 []开头就表示节(section)
	// 2.3 []不开头键值对儿(key=value)

}

func main() {
	var mc MysqlConfig
	fileName := "./config.ini"
	err := loadIni(fileName, &mc)
	if err != nil {
		fmt.Printf("Load ini failed, err: %v\n", err)
		return
	}
	fmt.Println("Mysql Config:", mc.Host, mc.Port, mc.User, mc.Password)
}