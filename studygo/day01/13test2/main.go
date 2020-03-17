package main

import "fmt"

func main()  {
	//编写代码打印9*9乘法表
	for a := 1; a < 10; a++{
		for b := 1; b <= a; b++{
			fmt.Printf("%d*%d=%d\t", b, a, a*b)
		}
		fmt.Println()
	}
	fmt.Println("------------------------")
	//编写代码打印倒立9*9乘法表
	for c := 9; c >= 1; c--{
		for d := 1; d <= c; d++{
			fmt.Printf("%d*%d=%d\t", d, c, c*d)
		}
		fmt.Println()
	}
}
