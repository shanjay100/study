package main

import (
	"fmt"
	"github.com/shanjay100/studygo/day05/10calc"
)

var x int8 = 10

const pi  = 3.14

func init()  {
	fmt.Println("自动执行")
	fmt.Println(x, pi)
}

func main()  {
	ret := calc.Add(1, 2)
	fmt.Println(ret)
}