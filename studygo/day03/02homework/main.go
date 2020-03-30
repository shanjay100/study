package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main()  {
	//1、判断字符串中汉字的数量
	//难点是判断字符是汉字
	s1 := "hello沙河"
	//1.1依次拿到字符串中的字符
	//1.2判断当前这个字符是不是汉字
	//1.3把汉字出现的次数累积得到最终结果
	var count int
	for _, c := range s1{
		if unicode.Is(unicode.Han, c) {
			count++
		}
	}
	fmt.Println(count)

	//2.判断how do you do单词出现的次数
	s2 := "how do you do"
	//2.1把字符串按照空格切割得到切片
	s3 := strings.Split(s2, " ")
	//2.2遍历切片存储到map中
	m1 := make(map[string]int, 10)
	for _, w := range s3 {
		//fmt.Println(w)
		//如果原来map中不存在w这个次数，那么出现次数=1
		//如果map中存在w这个key，那么出现次数+1
		if _, ok := m1[w]; !ok{
			m1[w] = 1
		} else {
			m1[w]++
		}
	}
	//2.3累加出现的次数
	for key, value := range m1 {
		fmt.Println(key, value)
	}

	//3、回文判断：字符串从左往右读和从右往左读都是一样
	//上海自来水来自海上
	//山西运煤车煤运西山
	ss := "a山西运煤车煤运西山a"
	//解题思路
	//把字符串的字符拿出来放到一个[]rune中
	r := make([]rune, 0, len(ss))
	for _, c := range ss {
		r = append(r, c)
	}
	fmt.Println("[]rune:", r)
	for i := 0; i < len(r)/2; i++{
		// 山 ss[0] s[len(ss)-1]
		// 西 ss[0] s[len(ss)-1-1]
		// 运 ss[0] s[len(ss)-1-2]
		// 煤 ss[0] s[len(ss)-1-3]
		// ...
		// n ss[i] s[len(ss)-1-i]
		if r[i] != r[len(r)-1-i] {
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")
}
