package main

import "fmt"

func main(){
	s := "hello沙河"
	//遍历字符串byte
	for i :=0; i < len(s); i++ {
		//byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	//遍历字符串rune
	for _, r := range s {
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
	//字符串修改
	s2 := "白萝卜"   //'白' '萝' ‘卜’
	s3 := []rune(s2) //把字符串强制转换成了一个rune切片
	//s3[0] = "红"
	s3[0] = '红'
	fmt.Println(string(s3)) //把rune切片强制切换成字符串

	//'' ""区别
	c1 := "红"
	c2 := '红' //rune(int32)
	fmt.Printf("c1:%T c2:%T\n", c1, c2)
	c3 := "H" //string
	c4 := 'H' //int32
	fmt.Printf("c3:%T c4:%T\n", c3, c4)
	fmt.Printf("%d\n", c4)

	//类型转换
	n := 10 //int
	var f float64
	f = float64(n)
	fmt.Println(f)
	fmt.Printf("%T\n", f)
}
