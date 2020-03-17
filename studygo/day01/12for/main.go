package main

import "fmt"

// for循环
// for 初始语句; 添加表达式; 结束语句{
//     循环体语句
//}

func main(){
	// 基本格式
	for i := 0; i < 10; i++{
		fmt.Println(i)
	}

	////变种1
	//var a = 10
	//for ; a < 15; a++{
	//	fmt.Println(a)
	//}
	//
	////变种2
	//var b = 15
	//for b < 20 {
	//	fmt.Println(b)
	//	b++
	//}

	// 无限循环

	//for range循环
	s := "hello沙河"
	for c, v := range s {
		fmt.Printf("%d %c \n", c, v)
	}
}
