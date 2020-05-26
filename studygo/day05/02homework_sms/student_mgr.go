package main

import (
	"fmt"
)

//有一个物件：
//1、它保存了一些数据 ---> 结构体的字段
//2、它有三个功能 	 ---> 结构体的方法


type student struct {
	id int64
	name string
}

//造一个学生的管理者
type studentMgr struct {
	allStudent map[int64]student
}

//查看学生
func (s studentMgr) showStudents()	 {
	//从s.allStudent这个map中把所有的学生逐个拎出来
	//stu是具体每一个学生
	for _, stu := range s.allStudent {
		fmt.Printf("学号：%d 姓名：%s\n", stu.id, stu.name)
	}
}

//增加学生
func (s studentMgr) addStudents()	 {
	//1、根据用户输入的内容，创建一个新的学生
	var  (
		stuId int64
		stuName string
	)
	//获取用户输入
	fmt.Print("请输入学号：")
	fmt.Scanln(&stuId)
	fmt.Print("请输入姓名：")
	fmt.Scanln(&stuName)
	//根据用户输入，创建结构体对象
	newStu := student{
		id:   stuId,
		name: stuName,
	}
	//2、把新的学生，放到s.allStudent这个map中
	s.allStudent[newStu.id] = newStu
	fmt.Println("添加学生成功")
}

//修改学生
func (s studentMgr) editStudents()	 {
	//1、获取用户输入的学号
	var stuId int64
	fmt.Print("请输入修改学生的学号：")
	fmt.Scanln(&stuId)
	//2 展示该学号对应的学生信息，如果没有提示查无此人
	stuObj, ok := s.allStudent[stuId]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	fmt.Printf("你要修改的学生信息如下：学号：%d 姓名：%s\n", stuObj.id, stuObj.name)
	//3、请输入与修改后的学生名
	fmt.Print("请输入学生的新姓名：")
	var newName string
	fmt.Scanln(&newName)
	//4、更新学生的姓名
	stuObj.name = newName //未更新map中的学生
	s.allStudent[stuId] = stuObj //更新map中的学生
}

//删除学生
func (s studentMgr) deleteStudents()	 {
	//1、请用户输入要删除学生的id
	var stuId int64
	fmt.Print("请输入删除学生的学号：")
	fmt.Scanln(&stuId)
	//2、去map找有没有这个学生，如果没有打印查无此人的提示
	_, ok := s.allStudent[stuId]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	//3、有的话就删除
	delete(s.allStudent, stuId)
	fmt.Println("删除成功！")
}
