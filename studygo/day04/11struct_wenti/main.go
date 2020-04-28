package main

import (
	"fmt"
)

//结构体遇到的问题

//1、myInt(100)是个啥
type myInt int

type person struct {
	name string
	age int
}

func (m myInt) hello()  {
	fmt.Println("我是一个int")
}

func main()  {
	//声明一个int32类型的变量x，它的值是10
	//方法一
	//var x int32
	//x = 10
	//方法二
	//var x int32 = 10
	//方法三
	//var x = int32(10)
	//方法四
	x := int32(10)
	fmt.Println(x)

	//声明一个myInt类型的变量m，它的值是100
	//方法一
	//var m myInt
	//m = 100
	//方法二
	//var m myInt = 100
	//方法三
	//var m = myInt(100)
	//方法四
	m := myInt(100) //强制类型转换
	fmt.Println(m)
	//m := myInt(100)
	//m.hello()

	//结构体初始化

	//方法一
	var p person
	p.name = "zhangsan"
	p.age = 18
	fmt.Println(p)
	var p1 person
	p1.name = "lisi"
	p1.age = 18
	fmt.Println(p1)
	//方法二
	//键值对初始化
	var p2  = person {
		name: "wangwu",
		age:  100,
	}
	fmt.Println(p2)
	//值列表初始化
	var p3 = person{
		"laoliu",
		0,
	}
	fmt.Println(p3)

	//声明同时初始化
	s1 := []int{1, 2, 3, 4}
	m1 := map[string]int{
		"stu1" : 1,
		"stu2" : 2,
		"stu3" : 3,
	}
	fmt.Println(s1, m1)
}

//为什么要有构造函数
func newPerson(name string, age int) *person {
	//别人调用我，我能给他一个person类型的变量
	return &person{
		name: name,
		age:  age,
	}
}