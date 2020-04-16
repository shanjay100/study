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

func f2(x *person)  {
	(*x).gender = "nv" //根据内存地址找到原变量，修改的是原变量
	//x.gender = "nv" 语法堂，自动根据指针找对应的变量
}

func main() {
	var p person
	p.name = "zhangSan"
	p.gender = "nan"
	f(p)
	fmt.Println(p.gender) //nan
	f2(&p)
	fmt.Println(p.gender) //nv

	//1、结构体指针1
	var p2 = new(person)
	//p2.name = "lisi"
	fmt.Printf("%T\n", p2)
	fmt.Printf("%p\n", p2) //p2保存的值是一个内存地址
	fmt.Printf("%p\n", &p2) //p2的内存地址
	//(*p2)

	//2、结构体指针2
	//2.1、key.value初始化
	var p3 = person{
		name:"wangWu",
		gender:"nan",
	}
	fmt.Printf("%#v\n", p3)
	//2.2、使用值列表的形式初始化，值的顺序要和结构体定义的字段顺序一致
	p4 := person{
		"zhaoLiu",
		"nan",
	}
	fmt.Printf("%#v\n", p4)
}