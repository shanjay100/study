package main

import "fmt"

//递归：函数自己调用自己
//递归适合处理那种问题相同、问题的规模越来越小的场景
//递归一定要有一个明确的退出条件

//上台阶的面试题
//n个台阶，一次可以走一步，也可以走两步，有多少种走法

func taijie(n uint64) uint64 {
	if n == 1{
		//如果只有一个台阶，只有一种走法
		return 1
	}
	if n== 2{
		//如果剩余两个台阶，有两种走法
		return 2
	}
	return taijie(n - 1) + taijie(n - 2)
}

func f(n uint64)  uint64{
	if n <= 1 {
		return 1
	}
	return n * f( n-1)
}

func main()  {
	//ret := f(1)
	//fmt.Println(ret)
	ret := taijie(4)
	fmt.Println(ret)
}