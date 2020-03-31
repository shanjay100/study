package main

import (
	"fmt"
)

//函数类型

func f1()  {
	fmt.Println("hello")
}

func f2() int  {
	return 10
}

//函数也可以作为参数的类型
func f3(x func() int)  {
	ret := x()
	fmt.Print(ret)
}

func f4(x, y int) int  {
	return x + y
}


//函数还可以作为返回值
func f5(x func() int) func(int, int) int {
	return f4
}

func main()  {
	a := f1
	fmt.Printf("%T\n", a)
	b := f2
	fmt.Printf("%T\n", b)
	c := f3
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", f4)
	f7 := f5(f2)
	fmt.Printf("%T\n", f7)
	fmt.Println(f7)
}
