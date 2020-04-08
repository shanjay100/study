package main

import "fmt"

//fmt

func main()  {
	fmt.Print("沙河")
	fmt.Print("na")
	fmt.Println("........")
	fmt.Println("沙河")
	fmt.Println("na")

	var m1 = make(map[string]int, 1)
	m1["理想"] = 100
	fmt.Printf("%v\n", m1)
	fmt.Printf("%#v\n", m1)
}
