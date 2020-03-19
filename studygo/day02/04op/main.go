package main

import "fmt"

func main()  {
	var (
		a = 5
		b = 2
	)
	//算术运算符
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++ //单独语句，不能放在=号右边赋值
	b-- //单独语句，不能放在=号右边赋值
	fmt.Printf("%d %d\n", a, b)

	//关系运算符
	fmt.Println(a == b) //Go语言是强类型，相同类型的变量才能比较
	fmt.Println(a != b) //Go语言是强类型，相同类型的变量才能比较
	fmt.Println(a >= b) //Go语言是强类型，相同类型的变量才能比较
	fmt.Println(a > b) //Go语言是强类型，相同类型的变量才能比较
	fmt.Println(a <= b) //Go语言是强类型，相同类型的变量才能比较
	fmt.Println(a < b) //Go语言是强类型，相同类型的变量才能比较

	//逻辑运算符
	//如果年龄大于18岁且年龄小于60岁
	age := 22
	if age > 18 && age < 60 {
		fmt.Println("苦逼上班族")
	}else {
		fmt.Println("不用上班")
	}
	// 如果年龄小于18岁 或者 年龄大于60岁
	if age < 18 || age > 60 {
		fmt.Println("不用上班")
	}else {
		fmt.Println("苦逼上班族")
	}
	// not取反，原来为真就为假，原来为假就为真
	isMarrid := false
	fmt.Println(isMarrid) //false
	fmt.Println(!isMarrid) //true

	//位运算：针对的是二进制
	//5 ：101
	//2 ：10
	//可能应用场景 ip、权限、

	//&：按位与（两位均为1才为1）
	fmt.Println(5 & 2) //000
	//|：按位或（两位有1个1就为1）
	fmt.Println(5 | 2)	//111
	//^ :按位异或（两位不一样则为1）
	fmt.Println(5 ^ 2) //111
	//<< ：将二进制位左移指定位数
	fmt.Println(5 << 1) //1010
	fmt.Println(1 << 10) //10000000000
	//>> ：将二进制位右移指定位数
	fmt.Println(5 >> 1) //10
	fmt.Println(5 >> 2) //1
	fmt.Println(5 >> 3) //0

	//var m = int8(1)	//只能存8位
	//fmt.Println(m << 10) //超出变量范围

	//赋值运算符，用来给变量赋值
	var x int
	x = 10
	x += 1 // x = x + 1
	x -= 1 // X = X - 1
	x *= 2 // x = x * 2
	x /= 2 // x = x / 2
	x %= 2 // x = x % 2
	x <<= 2 // x = x << 2
	x &= 2 // x = x & 2
	x |= 3 // x = x | 3
	x ^= 3 // x = x ^ 3
	x >>= 4 // x = x >> 4
	fmt.Println(x)
}
