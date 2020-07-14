package main

import (
	"time"
	"logger/modules"
)
func main() {
	// log := mylogger.NewLog("info")
	log := mylogger.NewFileLogger("info", "./", "access.log", 10*1024*1024) // 10MB
	for {
		log.Debug("调试日志")
		log.Info("信息日志")
		id := 10010
		name := "wovert"
		log.Warning("警告日志, id:%d, name:%s", id, name)
		log.Error("错误日志")
		log.Fatal("致命日志")
		time.Sleep(time.Second * 2)
	}
	defer log.Close()
}