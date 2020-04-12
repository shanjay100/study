package main

import "fmt"

//map

func main()  {
	var m1 map[string]int
	fmt.Println(m1 == nil) //还没有在初始化（没有在内存中开辟空间）
	m1 = make(map[string]int, 10) //初始化时，建议估算好map容量，避免在程序运行期间再动态扩容
	m1["理想"] = 18
	m1["现实"] = 35


	fmt.Println(m1)
	fmt.Println(m1["理想"])
	fmt.Println(m1["今天"]) //如果不存在这个key，返回对应类型的0值

	//一般用ok接收返回的bool值
	value, ok := m1["明天"]
	if !ok {
		fmt.Println("查无此key")
	}else {
		fmt.Println(value)
	}

	//map遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	//删除
	delete(m1, "现实")
	fmt.Println(m1)
	delete(m1, "后天") //删除不存在的key,不进行操作
}
