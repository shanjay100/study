package main

import "fmt"

//嵌套结构体
type address struct {
	province string
	city string
}

type workPlace struct {
	province string
	city string
}

type person struct {
	name string
	age int
	//province string
	//city string
	address //匿名嵌套
	workPlace
}

type company struct {
	name string
	//province string
	//city string
	addr address
}

func main()  {
	p1 := person{
		name: "zhangsan",
		age:  10,
		address: address{
			province: "山东",
			city: "威海",
		},
		workPlace: workPlace{
			province: "北京",
			city:     "北京",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.name, p1.address.city)
	//fmt.Println(p1.city) //现在自己的结构体查找这个字段，找不到就去匿名嵌套的结构体中查找该字段
	fmt.Println(p1.address.city)
	fmt.Println(p1.workPlace.city)
}
