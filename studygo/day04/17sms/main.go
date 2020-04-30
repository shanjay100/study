package main

import (
	"fmt"
	"os"
)

/*
	函数版学生管理系统——可以查看、新增、删除学生
*/

type student struct {
	id int64
	name string
}

var allStudent map[int64]*student //变量声明

//newStudent是student类型的构造函数
func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

//打印所有的学生
func showAllStudent()  {
	for k, v := range allStudent {
		fmt.Printf("学号：%d 姓名：%s\n", k, v.name)
	}
}

//添加学生
func addStudent()  {
	//向allStudent中添加一个新的学生
	//1、创建一个学生
	//1.1获取用户输入
	var (
		id int64
		name string
	)
	fmt.Print("请输入要添加学生的工号：")
	fmt.Scanln(&id)
	fmt.Print("请输入要添加学生的姓名：")
	fmt.Scanln(&name)
	//1.2造学生(调用student的构造函数)
	new := newStudent(id, name)
	//2、追加到allStudent这个map中
	allStudent[id] = new
}

//删除学生
func deleteStudent()  {
	//1、请用户输入要删除学生的学号
	var deleteId int64
	fmt.Print("请输入要删除学生的工号：")
	fmt.Scanln(&deleteId)
	//2、去allStudent这个map中根据学号删除对应的键值对
	delete(allStudent, deleteId)
}

func main()  {
	//初始化，开辟内存空间
	allStudent = make(map[int64]*student, 48)
	//主界面
	for {
		//1、打印菜单
		fmt.Println("欢迎光临学生管理系统")
		fmt.Println(`
		"1、查看学生"
		"2、新增学生"
		"3、删除学生"
		"4、退出"
				`)
		//2、等待用户选择要做什么
		fmt.Print("请输入您好选择的功能:")
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("您选择了%v功能\n", choice)
		//3、执行对应的选择
		switch choice {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("您的输入有误，请重新输入")
		}
	}
}
