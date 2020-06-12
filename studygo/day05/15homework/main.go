package main

import (
	"bufio"
	"fmt"
	"os"
)

//日志库
//需求：
//1、可以往不同的输出位置记录日志
//2、日志分为五种级别
//logger.Trace()
//logger.Debug()
//logger.Warning()
//logger.Info()
//logger.Error()

//获取用户输入时如果有空格

func userScan()  {
	var s string
	fmt.Print("请输入内容：")
	fmt.Scanln(&s)
	fmt.Printf("你输入的内容是：%s\n", s)
}

func useBufio()  {
	var s string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入内容：")
	s, _ = reader.ReadString('\n')
	fmt.Printf("你输入的内容是：%s\n", s)
}

func main()  {
	userScan()
	fmt.Fprintf(os.Stdout, "这是一条日志记录！")
	os.OpenFile("./xx.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	fmt.Fprintln(fileObj)
}

