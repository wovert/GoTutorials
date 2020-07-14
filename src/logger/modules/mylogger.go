package mylogger

import (
	"fmt"
	"strings"
	"errors"
	"runtime"
	"path"
)

// 在终端写日志相关内容

type LogLevel uint16

const (
	// 定义日志级别
	DEBUG LogLevel = iota
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
	UNKNOW
)

// Logger 日志结构
type ConsoleLogger struct {
  LogLevel
}

// 解析日志级别
func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
		case "debug":
			return DEBUG, nil
		case "trace":
			return TRACE, nil
		case "info":
			return INFO, nil
		case "warning":
			return WARNING, nil
		case "error":
			return ERROR, nil
		case "fatal":
			return FATAL, nil
		default:
			err := errors.New("无效的日志级别")
			return UNKNOW, err
	}
}

// NewLog
func NewLog(level string) ConsoleLogger {
	l, err := parseLogLevel(level)
	if err != nil {
		// panic(err)
	}
	return ConsoleLogger {
		LogLevel: l,
	}
}

func (l ConsoleLogger) enable(logLevel LogLevel) bool {
	return l.LogLevel <= logLevel
}

func getLogType(level LogLevel) string {
  switch level {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	}
	return "DEBUG"
}

// n调用函数层级 main时0，main函数里调用函数时1
func getInfo(skip int)(funcName, fileName string, line int) {
	// ok: ture获取成功状态， pc:函数信息， file:文件名， line:行号
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	funcName = strings.Split(funcName, ".")[1]

	// fmt.Println("funcName: ", funcName)
	// fmt.Println(file) // modules/mylogger.go
	// fmt.Println(line) // 114

	return
}