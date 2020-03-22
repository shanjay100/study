package main

import (
	"fmt"
	"sort"
)

//切片的练习

func main()  {
	//追加
	var a = make([]int, 5, 10) //创建切片，长度5， 容量10
	fmt.Println(a)
	for i := 0; i < 10; i++ {
		a = append(a, i) //追加10个元素，分别是[0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
	}
	fmt.Println(a)
	fmt.Println(cap(a))

	//切片排序
	var a1 = [...]int{3, 7, 8, 9, 1}
	sort.Ints(a1[:]) //对切片进行排序
	fmt.Println(a1)
}