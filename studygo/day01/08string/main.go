package main

import (
	"fmt"
	"strings"
)

func main()  {
	// \本来是具有特殊含义的，我需要告诉程序我写的\就是一个单纯的\
	path := "\"D:\\Go\\rc\\github.com\\shanjay100\\studygo\\day01\""
	fmt.Println(path)
	path1 := "'D:\\Go\\rc\\github.com\\shanjay100\\studygo\\day01'"
	fmt.Println(path1)
	s := "i'm ok"
	fmt.Println(s)
	// 多行字符串
	var s2 = `
君不见
黄河之水天上来
奔流到海不复回
	`
	fmt.Println(s2)
	s3 := `D:\Go\src\github.com\shanjay100\studygo\day01\`
	fmt.Println(s3)
	//字符串相关操作
	//字符串长度
	fmt.Println(len(s3))
	//字符串拼接
	name := "理想"
	world := "dsb"
	ss := name + world
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s", name, world)
	fmt.Println(ss1)
	//字符串分隔
	ret := strings.Split(s3, "\\")
	fmt.Println(ret)
	//字符串包含判断
	fmt.Println(strings.Contains(ss, "理性"))
	fmt.Println(strings.Contains(ss, "理想"))
	// 字符串前缀、后缀
	fmt.Println(strings.HasPrefix(ss, "理想"))
	fmt.Println(strings.HasSuffix(ss, "理想"))

	//字符串位置
	s4 := "abcdef"
	fmt.Println(strings.Index(s4, "c"))
	fmt.Println(strings.LastIndex(s4, "e"))

	//字符串拼接
	fmt.Println(strings.Join(ret, "+"))
}
