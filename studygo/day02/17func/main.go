package main

import "fmt"

//函数

//函数存在的意义
//函数是一段代码的封装
//把一段逻辑抽象出来封装到一个函数中，命名，每次用到它时直接用函数名调用就可以了
//使用函数能够让代码结构更清晰、更简洁

//函数的定义

func sum(x int, y int)(ret int){
	return x + y
}

//没有返回值
func f1(x int, y int) {
	fmt.Println(x + y)
}

//没有参数没有返回值
func f2()  {
	fmt.Println("f2")
}

//没有参数但有返回值
func f3() int{
	ret := 3
	return ret
}

//返回值可以命名也可以不命名
//命名返回值就相当于再函数中声明一个变量
func f4(x int, y int)(ret int) {
	ret = x + y
	return //使用命名返回值可以在return后省略
}

//多个返回值
func f5() (int, string)  {
	return 1, "沙河"
}

//参数类型简写:当参数中
func f6(x, y int) int  {

}

func main()  {
	r := sum(1, 2)
	fmt.Println(r)
}
