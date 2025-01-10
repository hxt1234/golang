package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// 结构体复习
// 结构体可以用来存储不同的数据类型的一个集合

// 定义一个person结构体
type Person struct {
	// 名字 string
	name string
	// 年龄 int
	age int
	// 匿名嵌套结构体，使用类型作为字段名称。
	student
}

// 构造函数，一般使用new关键字+对象名称

func newPerson(name string, age int) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

// 方法与接收者
// 方法是有接收者的函数，接收者是指哪个类型的变量可以调用对应的方法
// 接收者应该是对应类型的首字母的小写
func (p Person) eat() {
	fmt.Printf("%s 吃饭", p.name)
}

func main() {
	//结构体实例化方式一，键值对
	// p := Person{
	// 	name: "李白",
	// 	age:  19,
	// }
	// fmt.Println(p)
	// 结构体实例化方式二，通过new函数实例化
	// p := new(Person)
	// p.name = "杜甫"
	// p.age = 100
	// fmt.Println(p)
	// 结构体实例化三，通过变量复制
	// var p Person
	// p.name = "白居易"
	// p.age = 120
	// fmt.Println(p)
	// 结构体实例化四，通过值列表
	// p := Person{
	// 	"王维",
	// 	20,
	// }
	// fmt.Println(p)

	// p1 := newPerson("李白", 190)
	// fmt.Println(p1)
	//  p1属于person类，所以可以调用eat方法
	// p1.eat

	// 序列化：把Go语言中的结构体变量 --> json格式的字符串
	// 注意问题：1.如果序列化之后得到一个空值，需要注意结构体中的字段名首字母是否大写，因为使用序列化是使用json中的包，如果首字母小写，
	//			则表示对外不可见，不能访问。
	stu := student{100, "张三"}
	b, err := json.Marshal(stu)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", string(b))
	// 反序列化：json格式的字符串 --> Go语言中能够识别的结构体变量
	// 注意问题：1.如果反序列化之后得到一个空值，需要注意结构体中的字段名首字母是否大写，因为使用反序列化是使用json中的包，如果首字母小写，
	//			则表示对外不可见，不能访问。
	// 			2.反序列化要传递指针
	var P = `{"id":10,"addr":"湖北"}`
	type info struct {
		Id   int64
		Addr string
	}
	var p1 info
	err = json.Unmarshal([]byte(P), &p1)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("%v\n", p1)

	// 自定义类型与类型别名
	// 1. type Myint int  自定义类型
	// 2. type newInt = int   类型别名
	// 类型别名只在代码编写过程中有效，编译完之后就不存在，内置的byte和rune都属于类型别名。
}
