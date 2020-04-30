package main

import (
	"fmt"
	"os"

	//"os"
)

/*
使用“面向对象”的思维方式编写一个学生信息管理系统。
	1、学生有id、姓名、年龄、分数等信息
	2、程序提供展示学生列表、添加学生、编辑学生信息、删除学生等功能
*/

type student struct {
	id int64
	name string
	age int
	score int
}



//打印所有的学生
func showAllStudnet()  {

}

//添加学生
func addStudent()  {

}
	//向allStudent中添加一个新的学生
	//1、创建一个学生
	//1.1获取用户输入
	//1.2造学生(调用student的构造函数)
	//2、追加到allStudent这个map中
//编辑学生
func modifyStudent()  {

}
//删除学生
func deleteStudent()  {
	//1、请用户输入要删除学生的学号
	//2、去allStudent这个map中根据学号删除对应的键值对
}

func main()  {
	for {
		//主界面
		//1、打印菜单
		fmt.Println("欢迎使用学生管理系统")
		fmt.Print(`
			1、查看学生
			2、新增学生
			3、编辑学生
			4、删除学生
			5、退出
`)
		//2、等待用户选择要做什么
		var choice	int
		fmt.Print("请输入您的选择：")
		fmt.Scanln(&choice)
		//3、执行对应的选择
		switch choice {
		case 1:
			showAllStudnet()
		case 2:
			addStudent()
		case 3:
			modifyStudent()
		case 4:
			deleteStudent()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("您的选择有误，请重新输入，谢谢！")
		}
	}
}
