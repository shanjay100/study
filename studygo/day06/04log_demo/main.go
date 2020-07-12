package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// log demo

func main()  {
	//新建log文件
	fileObj, err := os.OpenFile("./xx.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	//设置日志输出位置
	log.SetOutput(fileObj)
	for {
		log.Println("这是一条测试的日志")
		time.Sleep(time.Second * 2)
	}
}
