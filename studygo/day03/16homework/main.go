package main

import "fmt"

/*
你有50个金币，需要分配给以下几个人：Matthew, Sara, Augustus, Heidi, Emilie, Peter, Giana, Adriano, Aaron, Elizabeth.
分配如下：
1、名字中包含1个e或者E分1个金币
2、名字中包含1个i或者I分2个金币
3、名字中包含1个o或者O分3个金币
4、名字中包含1个u或者U分4个金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现‘dispatchCoin’ 函数
 */

var (
	coins = 500
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
	name string
)

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
}

func dispatchCoin() (left int){
	for u, v := range distribution{
	fmt.Println(u, v)
	}

	//3、整个第二步执行完就能得到每个人最终分得的金币数和剩余金币数
	//1、一次拿到每个人的名字
	for _, name := range users {
		//2、拿到一个人名根据分金币的规则取分金币
		//	2.1、每个人分得金币的金币数应保存到distribution中
		//  2.2、记录剩下的金币数
		var name = 0
		for _, u := range name {
			var number = 0
			if u == 'e'{
				number = number + 1
			} else if u == 'E' {
				number = number + 1
			} else if u == 'i' {
				number = number + 2
			} else if u == 'I' {
				number = number + 2
			} else if u == 'o' {
				number = number + 3
			} else if u == 'O' {
				number = number + 3
			} else if u == 'u'{
				number = number + 4
			} else if u == 'U' {
				number = number + 4
			}
			fmt.Println(name, number)
		}
		//fmt.Println(name, number)
		//fmt.Printf("%v:")
	}
	//left = coins - number
	//distribution[name] = number
	//fmt.Println(distribution)
	//fmt.Println(left)
	return
}