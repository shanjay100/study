package main

import "fmt"

//构造函数

type person struct {
	name string
	age int
}

type dog struct {
	name string
}

//构造函数：约定俗成用new开头
//返回的是结构体还是结构体指针？
//当结构体比较大时，尽量使用结构体指针，减少程序的内存开销
func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

func newDog(name string) dog {
	return dog{
		name,
	}
}

func main()  {
	p1 := newPerson("lisi", 18)
	d1 := newDog("ahuang")
	fmt.Println(p1, d1)
}