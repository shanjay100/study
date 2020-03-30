package main

import "fmt"

//defer
//多个defer，后进先出（执行）
//多用于函数结束之前释放资源（文件句柄，数据库连接，socket连接）

func deferDemo()  {
	fmt.Println("start")
	defer fmt.Println("111") //defer把它后面的语句延迟到函数即将返回的时候再执行
	defer fmt.Println("222")
	defer fmt.Println("333")
	fmt.Println("end")
}

func main()  {
	deferDemo()
}
