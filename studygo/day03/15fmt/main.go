package main

import "fmt"

//fmt

func main()  {
	//fmt.Print("沙河")
	//fmt.Print("na")
	//fmt.Println("........")
	//fmt.Println("沙河")
	//fmt.Println("na")
	//
	//var m1 = make(map[string]int, 1)
	//m1["理想"] = 100
	//fmt.Printf("%v\n", m1)
	//fmt.Printf("%#v\n", m1)

	//获取用户输入
	//var s string
	//fmt.Scan(&s)
	//fmt.Println("用户输入的内容是：", s)
	var (
		name string
		age int
		class string
	)

	//fmt.Scanf("%s %d %s\n", &name, &age, &class)
	//fmt.Println(name, age, class)
	fmt.Scanln(&name, &age, &class)
	fmt.Println(name, age, class)
}
