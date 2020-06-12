package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//打开文件

func readFromFile()  {
	fileObj, err := os.Open("./main.go")
	if err != nil{
		fmt.Println("open file failed, err：", err)
		return
	}
	//记得关闭文件，通常用defer
	defer fileObj.Close()
	//读文件
	//var tmp = make([]byte, 128) //指定读的长度
	var tmp [128]byte
	for {
		n, err1 := fileObj.Read(tmp[:])
		if err1 != nil {
			fmt.Printf("read from file failed, err:%v", err1)
			return
		}
		fmt.Printf("读了%d个字节\n", n)
		fmt.Println(string(tmp[:n]))
		if n < 128 {
			return
		}
	}
}

//利用bufio这个包读取文件——一次读取一行
func readFromFileByBufio()  {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	//记得关闭文件
	defer fileObj.Close()
	//创建一个用来从文件读内容的对象
	reader := bufio.NewReader(fileObj)
	for {
		line, err1 := reader.ReadString('\n')
		if err1 == io.EOF {
			return
		}
		if err1 != nil {
			fmt.Printf("read line failed, err:%v", err1)
			return
		}
		fmt.Print(line)
	}
}

//利用ioutil读取文件
func readFromFileByIoutil()  {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	//记得关闭文件
	defer fileObj.Close()
	ret, err1 :=ioutil.ReadFile("./main.go")
	if err1 != nil {
		fmt.Printf("read file failed, err%v\n", err1)
	}
	fmt.Print(string(ret))
}

func main()  {
 	//readFromFile()
 	//readFromFileByBufio()
 	readFromFileByIoutil()
}
