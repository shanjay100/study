package main

import "fmt"

//结构体是值类型

type person struct {
	name, gender string
}

//Go语言中，函数传参永远是拷贝
func f(x person)  {
	x.gender = "nv" //修改的是副本的gender
}

func main() {
	var p person
	p.name = "zhangsan"
	p.gender = "nan"
	f(p)
	fmt.Println(p.gender)
}