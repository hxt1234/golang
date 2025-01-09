package main

import (
	"fmt"
	"os"
)

type Student struct {
	id   int64
	name string
}

var Allstudents = make(map[int64]*Student, 48)

func ShowallStudent() {
	for k, v := range Allstudents {
		fmt.Printf("学号%d 姓名%s\n", k, v.name)
	}
}

func newStudent(id int64, name string) *Student {
	return &Student{
		id:   id,
		name: name,
	}
}

func AddStudent() {
	var (
		id   int64
		name string
	)

	fmt.Print("请输入id:")
	fmt.Scan(&id)
	fmt.Print("请输入名字:")
	fmt.Scan(&name)
	newStu := newStudent(id, name)
	Allstudents[id] = newStu
}
func DeleteStudent() {
	var deleteid int64
	fmt.Print("请输入你要删除的ID：")
	fmt.Scan(&deleteid)
	for k, _ := range Allstudents {
		if k == deleteid {
			delete(Allstudents, deleteid)
		} else {
			fmt.Println("用户不存在，请重新输入！")
		}
	}
}

func main() {
	for {
		fmt.Print(`
		欢迎来到学生管理系统
		1. 查询学生
		2. 增加学生
		3. 删除学生
		4. 退出	
`)
		fmt.Print("请输入你的选择：")
		var choice int64
		fmt.Scan(&choice)
		fmt.Println("你选择了:", choice)
		switch choice {
		case 1:
			ShowallStudent()
		case 2:
			AddStudent()
		case 3:
			DeleteStudent()
		case 4:
			os.Exit(1)

		}

	}
}
