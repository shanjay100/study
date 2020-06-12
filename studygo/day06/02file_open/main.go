package main

import (
	"fmt"
	"os"
)

//文件操作

func f1()  {
	var fileObj *os.File
	var err error
	fileObj, err = os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
	}
	defer fileObj.Close()
}

func main()  {
	f1()
}
