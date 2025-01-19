package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
文件操作
1. os.Open打开文件，返回一个*file和一个err
2. bufio是在file的基础上封装了一层API，支持更多的功能

一、文件的写入
os.OpenFile()函数能够以指定模式打开文件，从而实现文件写入相关功能。
name：要打开的文件名 flag：打开文件的模式。 模式有以下几种：

模式		 含义
os.O_WRONLY	只写
os.O_CREATE	创建文件
os.O_RDONLY	只读
os.O_RDWR	读写
os.O_TRUNC	清空
os.O_APPEND	追加
perm：文件权限，一个八进制数。r（读）04，w（写）02，x（执行）01。
*/

/*
使用bufio获取用户输入
*/
func usebufioscan() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入内容：")
	s, _ = reader.ReadString('\n')
	fmt.Printf("输入的内容是：%s\n", s)
}

func main() {
	// 第一种方式，通过os.Open读取文件
	// f, err := os.Open("./main.go")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer f.Close()
	// var tmp = make([]byte, 512)
	// n, err := f.Read(tmp[:])
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("读了%d个字节\n", n)
	// fmt.Println(string(tmp[:]))

	// 第二种，使用bufio读取文件
	// file, err := os.Open("./main.go")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer file.Close()
	// 创建一个用来从文件中读的对象
	// reader := bufio.NewReader(file)
	// for {
	// 	line, err := reader.ReadString('\n')
	// 	if err == io.EOF {
	// 		fmt.Println("文件读完了")
	// 		return
	// 	}
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	fmt.Print(line)
	// }
	usebufioscan()

}
