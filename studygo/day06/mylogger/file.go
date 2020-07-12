package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

//往文件里面写日志相关代码

type FileLogger struct {
	Level LogLevel
	filePath 	string //日志文件保存的路径
	fileName 	string //日志文件保存的文件名
	fileObj *os.File
	errFileObj *os.File
	maxFileSize int64 //最大文件大小
}

//NewFileLogger构造函数
func NewFileLogger(levelstr, fp, fn string, maxSize int64) *FileLogger {
	LogLevel, err := parseLogLevel(levelstr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level:       LogLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	err = fl.initFile() //安装文件路径和文件名打开文件
	if err != nil {
		panic(err)
	}
	return fl
}

func (f *FileLogger)initFile() (error) {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err:%v\n", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFileName + ".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err log file failed, err:%v\n", err)
		return err
	}
	//日志文件已打开
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}

func (f *FileLogger) enable(loglevel LogLevel) bool {
	return loglevel >= f.Level
}

func (f *FileLogger) log(lv LogLevel, format string, a...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d]%s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
		if lv >= Error {
			//如果要记录的日志，大于等于Error级别，还需在err日志文件中再记录一遍
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d]%s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
		}
	}
}

func (f *FileLogger) Debug(format string, a ...interface{})  {
	f.log(Debug, format, a...)
}

func (f *FileLogger) Info(format string, a ...interface{})  {
	f.log(Info, format, a...)
}

func (f *FileLogger) Warning(format string, a ...interface{})  {
	f.log(Warning, format, a...)
}

func (f *FileLogger) Error(format string, a ...interface{})  {
	f.log(Error, format, a...)
}

func (f *FileLogger) Fatal(format string, a ...interface{})  {
	f.log(Fatal, format, a...)
}

func (f *FileLogger)close()  {
	f.fileObj.Close()
	f.errFileObj.Close()
}