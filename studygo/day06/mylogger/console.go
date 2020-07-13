package mylogger

import (
	"fmt"
	"time"
)

// 往终端写日志相关内容


// Logger日志结构体
type ConsoleLogger struct {
	Level LogLevel
}

// NewLog构造函数
func NewConsoleLog(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}

func (c ConsoleLogger) log(lv LogLevel, format string, a...interface{}) {
	if c.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		fmt.Printf("[%s] [%s] [%s:%s:%d]%s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
	}
}

func (c ConsoleLogger) enable(loglevel LogLevel) bool {
	return loglevel >= c.Level
}

func (c ConsoleLogger) Debug(format string, a ...interface{})  {
	//if c.enable(Debug) {
	//	log(Debug, format, a...)
	//}
	c.log(Debug, format, a...)
}

func (c ConsoleLogger) Info(format string, a ...interface{})  {
	//if c.enable(Info) {
	//	log(Info, format, a...)
	//}
	c.log(Info, format, a...)
}

func (c ConsoleLogger) Warning(format string, a ...interface{})  {
	//if c.enable(Warning) {
	//	log(Warning, format, a...)
	//}
	c.log(Warning, format, a...)
}

func (c ConsoleLogger) Error(format string, a ...interface{})  {
	//if c.enable(Error) {
	//	log(Error, format, a...)
	//}
	c.log(Error, format, a...)
}

func (c ConsoleLogger) Fatal(format string, a ...interface{})  {
	//if c.enable(Fatal) {
	//	log(Fatal, format, a...)
	//}
	c.log(Fatal, format, a...)
}