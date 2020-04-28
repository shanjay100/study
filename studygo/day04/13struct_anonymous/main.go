package main

import "fmt"

//匿名字段
//字段较少且坚定的场景
//不常用！！！

type person struct {
	string
	int
}

func main()  {
	p1 := person{
		"zhangsan",
		100,
	}
	fmt.Println(p1)
	fmt.Println(p1.string)
	fmt.Println(p1.int)
}