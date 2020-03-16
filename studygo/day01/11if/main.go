package main

import "fmt"

// if条件判断
func main()  {
	age := 19
	if age > 18 {
		fmt.Println("澳门首家上线啦：")
	} else {
		fmt.Println("改干嘛啦")
	}
	
	//其他格式
	if age > 35{
		fmt.Println("人到中年")
	} else if age > 18 {
		fmt.Println("青年")
	} else {
		fmt.Println("好好学习")
	}

	//作用域
	//age变量此时仅在if条件判断语句中生效
	if age1 := 19; age1 > 18{
		fmt.Printf("澳门")
	} else {
		fmt.Println("作业")
	}
	// fmt.Printf(age1) //此处无法找到age1
}