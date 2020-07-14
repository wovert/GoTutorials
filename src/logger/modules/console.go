package mylogger

import (
	"fmt"
	"time"
)

func (c ConsoleLogger)log(level LogLevel, msg string, a ...interface{}) {
	if c.enable(level) {
		msg = fmt.Sprintf(msg, a...)
		now := time.Now()
		funcName, fileName, line := getInfo(3)
		fmt.Printf("[%s] [%s] [%s:%s:%d]  %s\n", now.Format("2006-01-02 15:04:05"), getLogType(level), fileName, funcName, line, msg)
	}
}

// arg 可变类型，可有可无
func (c ConsoleLogger) Debug(msg string, arg ...interface{}) {
	// fmt.Println(msg)
	// if l.Level > DEBUG {
	c.log(DEBUG, msg, arg...)
}

func (c ConsoleLogger) Trace(msg string, arg ...interface{}) {
	c.log(TRACE, msg, arg...)
}

func (c ConsoleLogger) Info(msg string, arg ...interface{}) {
	c.log(INFO, msg, arg...)
}

func (c ConsoleLogger) Warning(msg string, arg ...interface{}) {
	c.log(WARNING, msg, arg...)
}

func (c ConsoleLogger) Error(msg string, arg ...interface{}) {
	c.log(ERROR, msg, arg...)
}

func (c ConsoleLogger) Fatal(msg string, arg ...interface{}) {
	c.log(FATAL, msg, arg...)
}