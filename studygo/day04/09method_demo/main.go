package main

import "fmt"

//方法

type dog struct {
	name string
}

//构造函数
func newDog(name string) dog {
	return dog{
		name,
	}
}

//方法是作用于特定类型的函数
func (d dog)wang()  {
	fmt.Println("%s:wangwangwang", d.name)
}

func main()  {
	d1 := newDog("ahuang")
	d1.wang()
}
