package main

import "fmt"

// 定义一个car类型的接口,用来实现run方法
type car interface {
	run()
}

type falali struct {
	brand string
}

type baoshijie struct {
	brand string
}

// 定义法拉利的run方法
func (f falali) run() {
	fmt.Printf("%s法拉利速度100迈~\n", f.brand)
}

// 定义保时捷的run方法
func (b baoshijie) run() {
	fmt.Printf("%s保时捷速度200迈~\n", b.brand)
}

func drive(x car) {
	x.run()
}

func main() {
	f1 := falali{
		brand: "法拉利",
	}
	b1 := baoshijie{
		brand: "保时捷",
	}
	drive(f1)
	drive(b1)
}
