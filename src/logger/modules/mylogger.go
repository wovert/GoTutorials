package mylogger

import (
	"fmt"
	"time"
	"strings"
	"errors"
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
type Logger struct {
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
func NewLog(level string) Logger {
	l, err := parseLogLevel(level)
	if err != nil {
		// panic(err)
	}
	return Logger {
		LogLevel: l,
	}
}

func (l Logger) enable(logLevel LogLevel) bool {
	return l.LogLevel <= logLevel
}

func (l Logger) Debug(msg string) {
	// fmt.Println(msg)
	// if l.Level > DEBUG {
	if l.enable(DEBUG) {
		now := time.Now()
		fmt.Printf("[%s] [DEBUG]  %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}

func (l Logger) Trace(msg string) {
	if l.enable(TRACE) {
		now := time.Now()
		fmt.Printf("[%s] [TRACE] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}

func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		now := time.Now()
		fmt.Printf("[%s] [INFO] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}

func (l Logger) Warning(msg string) {
	if l.enable(WARNING) {
		now := time.Now()
		fmt.Printf("[%s] [WARNING] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}

func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		now := time.Now()
		fmt.Printf("[%s] [ERROR] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}

func (l Logger) Fatal(msg string) {
	if l.enable(FATAL) {
		now := time.Now()
		fmt.Printf("[%s] [FATAL] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}