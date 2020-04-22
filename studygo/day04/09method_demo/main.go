package main

import "fmt"

//方法

//标识符：变量名、函数名、类项名、方法名
//Go语言中如果标识符首字母是大写的，就表示对外部包可见（暴露的、公有的）

// Dog 这是一个狗的结构体
type dog struct {
	name string
}

type person struct {
	name string
	age int
}

//构造函数
func newDog(name string) dog {
	return dog{
		name,
	}
}

func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

//方法是作用于特定类型变量的函数
//接收者表示的是调用该方法的具体类型变量，多用类型名首字母小写表示
func (d dog) wang()  {
	fmt.Printf("%s:wangwangwang\n", d.name)
}

//使用值接收者：传拷贝进去
func (p person) guoNian()  {
	p.age++
}

//使用指针接收者：传内存地址进去
func (p *person) guoDaNian()  {
	p.age++
}

func main()  {
	d1 := newDog("ahuang")
	d1.wang()
	p1 := newPerson("lisi", 18)
	fmt.Println(p1.age) //18
	p1.guoNian()
	fmt.Println(p1.age) //18
	p1.guoDaNian()
	fmt.Println(p1.age) //19
}
