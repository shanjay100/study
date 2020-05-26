package main

import (
	"fmt"
	"os"
)

//学生管理系统 --- 结构体

var smr studentMgr //声明一个全局变量学生管理对象：smr

//菜单函数
func showMenu()  {
	fmt.Println("--------welcome sms!--------")
	fmt.Println(`
	1、查看所有学生
	2、添加学生
	3、修改学生
	4、删除学生
	5、退出系统
	`)
}

func main()  {
	smr = studentMgr{
		allStudent: make(map[int64]student, 100),
	} //修改的全局的那个变量
	for {
		//显示菜单
		showMenu()
		//等待用户输入选项
		fmt.Print("请输入序号：")
		var choice int
		fmt.Scanln(&choice)
		fmt.Println("你输入的是：", choice)
		// 执行方法，谁的？
		switch choice {
		case 1:
			smr.showStudents()
		case 2:
			smr.addStudents()
		case 3:
			smr.editStudents()
		case 4:
			smr.deleteStudents()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("滚！！！")
		}
	}
}