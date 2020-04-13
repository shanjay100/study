package main

import "fmt"

//panic和recover

func f1()  {
	defer func() {
		erro := recover() //收集当前的错误信息
		fmt.Println("放手去爱")
		fmt.Println(erro)
	}()
	panic("犯了不可原谅的错误")
	fmt.Println("f1")
}

func f2()  {
	fmt.Println("f2")
}

func main()  {
	f1()
	f2()
}