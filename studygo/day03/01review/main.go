package main

import (
	"fmt"
)

//day03复习

func f1(a [3]int)  {
	//Go中函数传递的都是值（ctrl+c ctrl+v）
	a[1] =100 //此处修改的是副本的值
}

func main()  {
	var names string
	names = "理想"
	var ages [30]int //声明了一个ages变量，类型int
	ages = [30]int{1, 3, 4}
	fmt.Println(names)
	fmt.Println(ages)
	var ages1 = [...]int{1, 2, 3}
	fmt.Println(ages1)
	var ages2 = [...]int{1: 100, 99: 200}
	fmt.Println(ages2)

	//二维数组
	//多维数组只有最外层可以使用...
	var a1 [3][2]int
	a1 = [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(a1)

	//数组是值类型
	x := [3]int{1, 2, 3}
	y := x //将x赋值给y
	y[1] = 200 //修改y，不影响x
	fmt.Println(x, y) //[1 2 3] [1 200 3]
	f1(x)
	fmt.Println(x) //[1 2 3]

	//切片slice
	var s1 []int //没有分配内存， == nil
	fmt.Println(s1)
	fmt.Println(s1 == nil)
	s1 = []int{1, 2, 3}
	fmt.Println(s1)
	//make初始化，分配内存
	s2 := make([]bool, 2, 4) //[false false]
	fmt.Println(s2)
	s3 := make([]int, 0 ,4)
	fmt.Println(s3 == nil) //false

	b1 := []int{1, 2, 3} //[1 2 3]
	b2 := b1
	fmt.Println(b2) //[1 2 3]
	b2[1] = 200
	fmt.Println(b2) //[1 200 3]
	fmt.Println(b1) //[1 200 3]

	var c1 []int //nil
	//c1 = make([]int, 1)
	//c1[0] = 100
	//fmt.Println(c1)
	c1 = append(c1, 1)
	fmt.Println(c1)

	d1 := []int{1, 2, 3} //[1 2 3]
	d2 := d1
	var d3 = make([]int, 3, 3)
	copy(d3, d1)
	fmt.Println(d2) //[1 2 3]
	d2[1] = 200
	fmt.Println(d2) //[1 200 3]
	fmt.Println(d1) //[1 200 3]
	fmt.Println(d3) //?

	//指针
	//Go里面的指针只能读不能修改，不能修改指针变量指向的地址
	addr := "沙河"
	addrP := &addr
	fmt.Println(addrP) //内存地址
	fmt.Printf("%T\n", addrP)
	addrV := *addrP //根据内存地址找值
	fmt.Println(addrV)

	//map
	var m1 map[string]int
	fmt.Println(m1 == nil)
	m1 = make(map[string]int, 10)
	m1["理想"] = 100
	fmt.Println(m1)
	fmt.Println(m1["现实"]) //如果key值不存在，返回的是value对应类型的零值
	score, ok := m1["现实"]
	if !ok {
		fmt.Println("wu")
	} else {
		fmt.Println("现实是:", score)
	}
	delete(m1, "明天") //删除的key不存在，什么都不做
	delete(m1, "理想")
	fmt.Println(m1)
	fmt.Println(m1 == nil) //已经申请了内存，不为空
}
