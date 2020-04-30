package main

import (
	"encoding/json"
	"fmt"
)

//结构体与json

//1、序列化：把Go语言中的结构体变量 --> json格式的字符串
//2、反序列化：json格式的字符串 --> Go语言中能够识别的结构体变量

type person struct {
	Name string // 首字母大写，转换是才能识别数据
	Age int `json:"age"` //反引号
}

func main()  {
	p1 := person{
		Name: "zhangsan",
		Age:  10,
	}
	//序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed, err:%v", err)
		return
	}
	fmt.Println(string(b))
	fmt.Printf("%#v\n", string(b))

	//反序列化
	str := `{"name":"理想","age":18}`
	var p2 person
	json.Unmarshal([]byte(str), &p2) //传指针是为了能在json.Unmarshal内部修改p2的值
	fmt.Printf("%#v\n", p2)
}