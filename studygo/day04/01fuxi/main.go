package main

import "fmt"

func xianShi(name string)  {
	fmt.Println("hello", name)
}

//函数作为参数
func liXiang(f func(string), name string)  {
	f(name)
}

//函数作为返回值
//函数内部声明匿名函数
func mingTian() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}

func jinTian(f func())  {
	f()
}

//闭包
func bibao(f func(string), name string) func() {
	return func() {
		f(name)
	}
}

func main()  {
	liXiang(xianShi, "理想")
	ret := mingTian()
	fmt.Printf("%T\n", ret)
	sum := ret(10, 20)
	fmt.Println(sum)
	fc := bibao(xianShi, "现实")
	jinTian(fc)
}
