package main

import (
	"fmt"
	"io"
	"os"
)

//文件操作

//func f1()  {
//	var fileObj *os.File
//	var err error
//	fileObj, err = os.Open("./main.go")
//	if err != nil {
//		fmt.Printf("open file failed, err:%v", err)
//	}
//	defer fileObj.Close()
//}

func f2()  {
	//打开要操作的文件
	fileObj, err := os.OpenFile("./test.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	//借助临时文件,因为没有办法直接在文件中插入内容
	tmpFile, err2 := os.OpenFile("./tmp.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err2 != nil {
		fmt.Printf("create tmp file failed, err:%v\n", err2)
		return
	}
	defer tmpFile.Close()
	//读取文件
	var ret [1]byte
	n, err1 := fileObj.Read(ret[:])
	if err1 != nil {
		fmt.Printf("read from file failed, err:%v", err1)
		return
	}
	//写入临时文件
	tmpFile.Write(ret[:n])
	//写入要插入的内容
	var s []byte
	s = []byte{'c'}
	tmpFile.Write(s)
	//接着把源文件后续的内容写入临时文件
	var x [1024]byte
	for {
		m, err3 := fileObj.Read(x[:])
		if err3 == io.EOF {
			tmpFile.Write(x[:m])
			break
		}
		if err3 != nil {
			fmt.Printf("read from file failed, err:%v\n", err3)
			return
		}
		tmpFile.Write(x[:m])
	}
	fileObj.Close()
	tmpFile.Close()
	//将临时文件重命名为原文件，完成替换
	os.Rename("./tmp.txt", "./test.txt")
}

func main()  {
	//f1()
	f2()
}
