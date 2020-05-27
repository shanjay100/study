package main

import "fmt"

//接口的实现

type animal interface {
	move()
	eat(string) //带参数的方法，与不带参数的方法即使明知一样也是不同类型！！！
}

type cat struct {
	name string
	feet int8
}

func (c cat) move()  {
	fmt.Println("走猫步")
}

func (c cat) eat(food string)  {
	fmt.Printf("猫吃%s\n", food)
}

type chicken struct {
	feet int8
}

func (c chicken) move()  {
	fmt.Println("鸡冻！")
}

func (c chicken) eat(food string)  {
	fmt.Printf("鸡吃%s\n", food)
}

func main()  {
	var a1 animal //定义一个接口类型的变量,动态类型、动态值
	fmt.Printf("%T\n", a1) //未赋值未空类型，空值

	bc := cat{ //定义一个cat类型的变量bc
		name: "淘气",
		feet: 4,
	}

	a1 = bc
	a1.eat("小黄鱼")
	fmt.Println(a1)
	fmt.Printf("%T\n", a1)

	kfc := chicken{
		feet: 2,
	}
	a1 = kfc
	a1.eat("虫子")
	fmt.Println(a1)
	fmt.Printf("%T\n", a1)
}