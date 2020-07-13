package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

//往文件里面写日志相关代码

//FileLogger文件日志结构体
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

//文件初始化
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

//判断是否需要记录该日志
func (f *FileLogger) enable(loglevel LogLevel) bool {
	return loglevel >= f.Level
}

//判断文件是否需要切割——判断大小
func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed, err:%v\n", err)
		return false
	}
	//如果当前文件大小，大于等于日志文件的最大值，返回true
	return fileInfo.Size() >= f.maxFileSize
	//fmt.Printf("文件大小是：%dB\n", fileInfo.Size())
}

//切割文件
func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	//需要切割日志文件
	nowStr:= time.Now().Format("200060102150405000")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("open file info failed, err:%v\n", err)
		return nil, err
	}
	logName := path.Join(f.filePath, fileInfo.Name()) //拿到当前日志文件的完整路径
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr) //拼接一个日志文件备份的名字
	//1、关闭当前的日志文件
	file.Close()
	//2、备份一下rename xx.log --> xx.log.bak202007131436
	os.Rename(logName, newLogName)
	//3、打开一个新的日志文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new log file faile:%v\n", err)
		return nil, err
	}
	//4、将打开的新日志文件对象赋值给f.fileObj
	return fileObj, nil
}

//记录日志的方法
func (f *FileLogger) log(lv LogLevel, format string, a...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		if f.checkSize(f.fileObj) {
			newFile, err := f.splitFile(f.fileObj) //日志文件
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d]%s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
		if lv >= Error {
			if f.checkSize(f.errFileObj) {
				newFile, err := f.splitFile(f.errFileObj) //日志文件
				if err != nil {
					return
				}
				f.errFileObj = newFile
			}
			//如果要记录的日志，大于等于Error级别，还需在err日志文件中再记录一遍
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d]%s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
		}
	}
}

//Debug日志
func (f *FileLogger) Debug(format string, a ...interface{})  {
	f.log(Debug, format, a...)
}

//Info日志
func (f *FileLogger) Info(format string, a ...interface{})  {
	f.log(Info, format, a...)
}

//Warning日志
func (f *FileLogger) Warning(format string, a ...interface{})  {
	f.log(Warning, format, a...)
}

//Error日志
func (f *FileLogger) Error(format string, a ...interface{})  {
	f.log(Error, format, a...)
}

//Fatal日志
func (f *FileLogger) Fatal(format string, a ...interface{})  {
	f.log(Fatal, format, a...)
}

//关闭日志
func (f *FileLogger)close()  {
	f.fileObj.Close()
	f.errFileObj.Close()
}