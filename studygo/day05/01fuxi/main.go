package main

import (
	"encoding/json"
	"fmt"
)

//复习结构体

type person struct {
	a int
	b int
	c int
}

//构造（结构体变量的）函数，返回值是对应的结构体类型
func newPerson(a int, b int, c int,) person {
	return person{
		a:	a,
		b:	b,
		c:	c,
	}
}

//方法
//接收者使用对应类型的首字母小写
//指定了接收者之后，只有接收者这个类型的变量才能调用这个方法
func (p person) dream (str string)  {
	fmt.Printf("%v的梦想是%v。\n", p.a, str)
}

func (p person) guoNian()  {
	p.b++ //此处的P是p1的副本，修改的是副本
}

//指针接收者
//1、需要修改结构体变量的值时要试用指针接收者
//2、结构体本身比较大，拷贝的内存开销比较大时也要试用指针接收者
//3、如果有一个方法试用了指针接收者，其他的方法为了统一，也要试用指针接收者
func (p *person) fangJia() {
	p.c++ //此处修改的P是p1的副本，修改的是副本
}

func main()  {
	//匿名结构体
	var a = struct {
		x int
		y int
	}{10, 20} //初始值
	fmt.Println(a)

	p1 := person{
		a: 1,
		b: 2,
		c: 3,
	}

	//调用构造函数生成person类型
	p2 := newPerson(1, 2, 3)
	fmt.Println(p1, p2)
	p1.dream("做个咸鱼")
	p2.dream("学习")

	fmt.Println(p1.b)
	p1.guoNian()
	fmt.Println(p1.b)

	fmt.Println(p1.c)
	p1.fangJia()
	fmt.Println(p1.c)

	//结构体嵌套
	type addr struct {
		province string
		city string
	}
	type student struct {
		name string
		address addr //嵌套别的结构体
		addr //匿名嵌套，就使用类型做名称
	}

	type point struct {
		X int `json:"zhoulin"`
		Y int `json:"baodelu"`
	}

	p10 := point{100, 200}
	//序列化
	b, err := json.Marshal(p10)
	//如果出错了
	if err != nil {
		fmt.Printf("marshal failed, err:%v\n", err)
	}
	fmt.Println(string(b))
	//反序列化：由字符串 --> Go中的结构体变量
	str1 := `{"zhoulin":10010,"baodelu":10086}`
	var p20 point
	err = json.Unmarshal([]byte(str1), &p20)
	if err != nil {
		fmt.Printf("unmarshal failed, err:%v\n", err)
	}
	fmt.Println(p20)
}