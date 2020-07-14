package mylogger

import (
	"os"
	"path"
	"fmt"
	"time"
)

type FileLogger struct {
	Level LogLevel
	filePath string // 日志文件保存的路径
	fileName string // 日志文件保存的文件明
	fileObj *os.File
	errFileObj *os.File
	maxFileSize int64
}

func NewFileLogger(levelStr, fp, fn string, maxSize int64) (*FileLogger) {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger {
		Level: logLevel,
		filePath: fp,
		fileName: fn,
		maxFileSize: maxSize,
	}
	err = fl.initFile() // 按照文件路径和文件名将文件打开
	if err != nil {
		panic(err)
	}
	return fl
}

func (f *FileLogger) initFile() (err error) {
	fileFullName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fileFullName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Open log file failed, err:%v\n", err)
		return err
	}

	errFileObj, err2 := os.OpenFile(fileFullName + ".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err2 != nil {
		fmt.Printf("Open error log file failed, err:%v\n", err2)
		return err2
	}

	// 日志文件都已经打开了
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}

func (f *FileLogger)log(level LogLevel, msg string, a ...interface{}) {
	if f.enable(level) {
		msg = fmt.Sprintf(msg, a...)
		now := time.Now()
		funcName, fileName, line := getInfo(3)
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d]  %s\n", now.Format("2006-01-02 15:04:05"), getLogType(level), fileName, funcName, line, msg)

		// 大于等于error级别
		if level >= ERROR {
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d]  %s\n", now.Format("2006-01-02 15:04:05"), getLogType(level), fileName, funcName, line, msg)
		}
	}
	// defer f.Close()
}

func (f *FileLogger)Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}

func (f *FileLogger) enable(logLevel LogLevel) bool {
	return f.Level <= logLevel
}

func (f *FileLogger) Debug(msg string, arg ...interface{}) {
	f.log(DEBUG, msg, arg...)
}

func (f *FileLogger) Trace(msg string, arg ...interface{}) {
	f.log(TRACE, msg, arg...)
}

func (f *FileLogger) Info(msg string, arg ...interface{}) {
	f.log(INFO, msg, arg...)
}

func (f *FileLogger) Warning(msg string, arg ...interface{}) {
	f.log(WARNING, msg, arg...)
}

func (f *FileLogger) Error(msg string, arg ...interface{}) {
	f.log(ERROR, msg, arg...)
}

func (f *FileLogger) Fatal(msg string, arg ...interface{}) {
	f.log(FATAL, msg, arg...)
}
