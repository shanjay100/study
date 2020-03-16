package main

import (
	"fmt"
)

var	(
	a int
	b float64
	c bool
	d string
)

func main()  {
	//编写代码分别定义一个整型、浮点型、布尔型、字符串型变量，使用fmt.Printf()搭配%T分别打印出上述变量的值和类型。
	a = 10
	b = 1.123
	c = true
	d = "1234"
	fmt.Printf("a:%v b:%v c:%v d:%v\n", a, b, c, d)
	fmt.Printf("a:%T b:%T c:%T d:%T\n", a, b, c, d,)
	//编写代码统计出字符串"hello沙河小王子"中汉字的数量
	s := "hello沙河小王子"
	char := []rune(s)
	d := 0
	for i:= 0; i < len(char); i++{
		if len(string(char[i])) >= 3{
			d++
		}
		fmt.Printf("%v\n", string(char[i]))
	}
	fmt.Printf("汉字总数量：%v\n", d)
}

