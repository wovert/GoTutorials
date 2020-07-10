package main

import (
	"time"
	"logger/modules"
)
func main() {
	log := mylogger.NewLog("ddd")
	log.Debug("调试日志")
	log.Info("信息日志")
	time.Sleep(time.Second)
	log.Warning("警告日志")
	log.Error("错误日志")
	log.Fatal("致命日志")
}