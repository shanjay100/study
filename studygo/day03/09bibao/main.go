package main

import (
	"fmt"
)

//闭包

func f1(f func())  {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

//定义一个函数，对f2进行包装
//func f3(f func(int, int)) func() {
//	tmp := func() {
//		f()
//	}
//	return tmp
//}

//要求：f1(f2)

func f3(f func(int, int), x, y int) func()  {
	tmp := func() {
		//fmt.Println(x + y)
		////f2()
		f(x, y)
	}
	return tmp
}

func main()  {
	ret := f3(f2, 100, 200) //把原来需要传递两个int类型的参数包装成一个不需要传参的函数
	f1(ret)
}

//func main()  {
//	ret := liXing(f2, 100, 200)
//	ret()
//}
//
//func liXing(x func(int, int), m, n int) func() {
//	tmp := func() {
//		x(m, n)
//	}
//	return tmp
//	//tmp() //x(m, n)
//	//x(m, n)
//}
