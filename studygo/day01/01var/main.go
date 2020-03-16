package main

import "fmt"

//Go语言中推荐使用驼峰式命名
//var student_name string
//var studentName string——推荐使用
//var StudentName string

//声明变量
//var name string
//var age int
//var isok bool

//批量声明
var (
	name string // ""
	age	int		// 0
	isOk bool	// false
)

func main() {
	name = "理想"
	age = 16
	isOk = true
	//Go中非全局变量声明后必须使用，不适用编译不了——看IDEA版本
	fmt.Print(isOk)
	fmt.Println()
	fmt.Printf("name:%s\n", name) //%s占位符，使用name这个变量的值去替换占位符
	fmt.Println(age)
	//声明变量同时赋值
	var s1 string = "who"
	fmt.Println(s1)
	//类型推导（根据值判断该变量是什么类型）
	var s2 = "20"
	fmt.Println(s2)
	//简短变量声明，只能在函数里面用，函数外不能使用:=
	s3 := "111"
	fmt.Println(s3)
	//同一个作用域中{}，不能重复声明同名变量
	//s1 := "222" //
}