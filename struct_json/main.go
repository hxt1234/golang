package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// 定义一个人的结构体

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// 序列化：把Go语言中的结构体变量 --> json格式的字符串
// 反序列化：json格式的字符串 --> Go语言中能够识别的结构体变量

func main() {
	p := Person{
		Name: "李白",
		Age:  19,
	}
	//序列化
	b, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(reflect.TypeOf(b))
	fmt.Printf("%v\n", string(b))
	// 反序列化
	var str = `{"name":"杜甫","age":19}`
	var p2 Person
	json.Unmarshal([]byte(str), &p2)
	fmt.Printf("%#v\n", p2)
}
