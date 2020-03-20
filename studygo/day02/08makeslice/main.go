package main

import "fmt"

//make()函数创造切片

func main()  {
	s1 := make([]int, 5) //容量缺省，默认与长度一致
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

	s2 := make([]int, 0, 10)
	fmt.Printf("s2=%v len(s2)=%d cap(s2)=%d\n", s2, len(s2), cap(s2))

	s3 := make([]int, 5, 10)
	fmt.Printf("s3=%v len(s3)=%d cap(s3)=%d\n", s3, len(s3), cap(s3))

	//切片的赋值
	a1 := []int{1, 3, 5}
	a2 := a1 //a1, a2都指向了同一个底层数组
	fmt.Println(a1, a2)
	a1[0] = 1000
	fmt.Println(a1, a2)

	//切片的遍历
	//1、索引遍历
	for i := 0;i < len(a1); i++ {
		fmt.Println(a1[i])
	}
	//2、for range循环
	for _, v := range a1 {
		fmt.Println(v)
	}
}