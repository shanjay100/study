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
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() (left int){
	//1、一次拿到每个人的名字
	for _, name := range users {
		//2、拿到一个人名根据分金币的规则取分金币
		//distribution[name] = 0
		//var number = 0
		for _, u := range name {
			if u == 'e' || u == 'E'{
				//	2.1、每个人分得金币的金币数应保存到distribution中
				distribution[name] += 1
				//  2.2、记录剩下的金币数
				coins -= 1
			} else if u == 'i' || u == 'I'{
				distribution[name] += 2
				coins -= 2
			} else if u == 'o' || u == 'O'{
				distribution[name] += 3
				coins -= 3
			} else if u == 'u' || u == 'U'{
				distribution[name] += 4
				coins -= 4
			}
		}
		//3、整个第二步执行完就能得到每个人最终分得的金币数和剩余金币数
		//distribution[name] = number
		left = coins
		//fmt.Println(name, number)
	}
	//fmt.Println(distribution)
	return
}

func main() {
	left := dispatchCoin()
	fmt.Printf("分完剩下：%v个金币\n", left)
	for u, v := range distribution{
		fmt.Printf("%s分得：%v个金币\n", u, v)
	}
}