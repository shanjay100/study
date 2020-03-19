package main

import "fmt"

// goto

func main()  {
	////跳出多层for循环
	//var flag = false
	//for i := 0; i < 10; i++{
	//	for j :=0; j < 10; j++{
	//		if j == 2{
	//			flag = true
	//			break //跳出内层的for循环
	//		}
	//		fmt.Printf("%v-%v\n", i, j)
	//	}
	//	if flag{
	//		break //跳出内层的for循环
	//	}
	//}

	//goto+label实现多层for循环
	for i := 0; i < 10; i++{
		for j :=0; j < 10; j++{
			if j == 2{
				//设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
breakTag: 	//退出标签
	fmt.Println("结束for循环")
}
