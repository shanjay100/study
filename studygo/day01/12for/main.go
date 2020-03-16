package main

import "fmt"

// for循环

func main(){
	//基本格式
	for i := 0; i < 10; i++{
		fmt.Println(i)
	}

	//变种1
	var a = 10
	for ; a < 15; a++{
		fmt.Println(a)
	}

	//变种2
	var b = 15
	for b < 20 {
		fmt.Println(b)
		b++
	}
}
