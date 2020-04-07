package main

import "fmt"

//defer

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	defer calc("1", x, calc("10", x, y))
	x = 0
	defer calc("2", x, calc("20", x, y))
	y = 1
}

//1、defer calc("1", 1, calc("10", 1, 2))
//2、calc("1", 1, 3) //"10", 1, 2, 3
//3、defer calc("1", 1, 3)
//4、a = 0
//5、defer calc("2", 0, calc("20", 0, 2))
//6、calc("2", 0, 2) //"20", 0, 2, 2
//7、defer calc("2", 0, 2)
//8、b = 1
//9、calc("2", 0, 2) //"2" 0 2 2
//10、calc("1", 1, 3) // "1" 1 3 4

//1、"10", 1, 2, 3
//2、"20", 0, 2, 2
//3、"2" 0 2 2
//4、"1" 1 3 4
