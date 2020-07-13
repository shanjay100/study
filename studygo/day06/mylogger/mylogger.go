package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

type LogLevel uint16

//Logger 接口
type Logger interface {
	Debug(format string, a ...interface{})
	Info(format string, a ...interface{})
	Warning(format string, a ...interface{})
	Error(format string, a ...interface{})
	Fatal(format string, a ...interface{})
}

const (
	//定义日志级别
	Unknown  LogLevel = iota //级别为0
	Debug  //级别为1
	Trace
	Info
	Warning
	Error
	Fatal

)

func parseLogLevel(s string) (LogLevel, error)  {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return Debug, nil
	case "trace":
		return Trace, nil
	case "info":
		return Info, nil
	case "warning":
		return Warning, nil
	case "error":
		return Error, nil
	case "fatal":
		return Fatal, nil
	default:
		err := errors.New("无效的日志级别")
		return Unknown, err
	}
}

func getLogString(lv LogLevel) string {
	switch lv {
	case Debug:
		return "Debug"
	case Trace:
		return "Trace"
	case Info:
		return "Info"
	case Warning:
		return "Warning"
	case Error:
		return "Error"
	case Fatal:
		return "Fatal"
	default:
		return "Debug"
	}
}

func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	funcName = strings.Split(funcName, ".")[1]
	return
}
