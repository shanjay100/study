package main

import (
	"fmt"
)

//结构体
type person struct {
	name string
	age int
	gender string
	hobby []string
}

func main()  {
	//声明一个person类型的变量hu
	var hu person
	//通过字段赋值
	hu.name = "zhangsan"
	hu.age = 18
	hu.gender = "男"
	hu.hobby = []string{"唱", "跳", "RAP"}
	fmt.Println(hu)
	//访问变量hu的字段
	fmt.Printf("%T\n", hu)
	fmt.Println(hu.name)

	var lin person
	lin.name = "li"
	lin.age = 19
	fmt.Printf("type:%T value:%v\n", lin, lin)

	//匿名结构体，多用于临时场景
	var s struct{
		x string
		y int
	}
	s.x = "hei"
	s.y = 100
	fmt.Printf("type:%T value:%v\n", s,s)
}