package main

import "fmt"

//使用值接收者和指针接收者的区别

type animal interface {
	move()
	eat(string) //带参数的方法，与不带参数的方法即使明知一样也是不同类型！！！
}

type cat struct {
	name string
	feet int8
}


//使用值接收者实现了接口的所有方法
//func (c cat) move()  {
//	fmt.Println("走猫步")
//}
//
//func (c cat) eat(food string)  {
//	fmt.Printf("猫吃%s\n", food)
//}

func (c *cat) move()  {
	fmt.Println("走猫步")
}

func (c *cat) eat(food string)  {
	fmt.Printf("猫吃%s\n", food)
}

func main()  {
	var a1 animal
	c1 := &cat{"tom", 4} //cat
	c2 := &cat{"假老练",4} //*cat
	a1 = c1
	fmt.Println(a1)
	a1 = c2
	fmt.Println(a1)
}
