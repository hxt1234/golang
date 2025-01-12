package main

import (
	"fmt"
	"reflect"
)

/*
1、接口是一种类型，他是一种特殊的类型，规定了变量有哪些方法。
2、接口的定义

	type 接口名 interface{
		方法名1(参数1，参数2，。。。。。)(返回值1，返回值2，。。。。。。)
		方法名2(参数1，参数2，。。。。。)(返回值1，返回值2，。。。。。。)
	}

3、一个变量如果实现了接口中规定了所有的方法，那么这个变量就实现了这个接口，可以称为这个接口类型的变量
*/
type speaker interface {
	speak() //只要实现了speak方法的变量都是speaker类型
}

type Cat struct{}

type Dog struct{}

type pserson struct{}

func (c Cat) speak() {
	fmt.Print("喵喵喵~\n")
}

func (d Dog) speak() {
	fmt.Print("汪汪汪~\n")
}

func (p pserson) speak() {
	fmt.Print("啊啊啊啊啊~\n")
}

func Da(x speaker) {
	x.speak()
}

func main() {
	var c1 Cat
	var d1 Dog
	var p1 pserson
	Da(d1)
	Da(c1)
	Da(p1)
	var ss speaker
	ss = c1
	fmt.Print(ss)
	fmt.Println(reflect.TypeOf(ss))
}
