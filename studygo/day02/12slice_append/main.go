package main

import "fmt"

//关于append删除切片中得某个元素
func main()  {
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13, 15, 17}
	s1 := a1[:]

	//删除索引为1的元素“3”
	s1 = append(s1[0:1], s1[2:]...)
	fmt.Println(s1) //[1 5 7 9 11 13 15 17]
	fmt.Println(a1) //[1 5 7 9 11 13 15 17 17]
}
