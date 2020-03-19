package main

import (
	"fmt"
)

func main()  {
	//求数组[1, 3, 5, 7, 8]所有元素的和
	a :=  [5]int{1, 3, 5, 7, 8}
	b := 0

	for i := 0; i < len(a); i++ { //遍历数组a的所有元素
			b += a[i]			  //b = a[0] + ... + a[4]
		}
	fmt.Println(b)

	//找出数组中和为指定值的两个元素的下标
	//比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
	for i, _ := range a{
		for l := i+1; l < len(a); l++{
			if a[i] + a[l] == 8{
				fmt.Printf("(%d, %d)", i, l)
			}
		}
	}
}


