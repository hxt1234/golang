package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
1.使用go关键字开启一个groutine
2.main函数执行完了，由main函数创建的哪些groutine就结束了
3.WaitGroup在go语言中，用于线程同步，单从字面意思理解，wait等待的意思，group组、团队的意思，
WaitGroup就是指等待一组，等待一个系列执行完成后才会继续向下执行。
正常情况下，goroutine的结束过程是不可控制的，我们可以保证的只有main goroutine的终止。
这时候可以借助sync包的WaitGroup来判断goroutine是否完成。
它有3个方法：
    Add()：每次激活想要被等待完成的goroutine之前，先调用Add()，用来设置或添加要等待完成的goroutine数量
        例如Add(2)或者两次调用Add(1)都会设置等待计数器的值为2，表示要等待2个goroutine完成
    Done()：每次需要等待的goroutine在真正完成之前，应该调用该方法来人为表示goroutine完成了，该方法会对等待计数器减1
    Wait()：在等待计数器减为0之前，Wait()会一直阻塞当前的goroutine
    也就是说，Add()用来增加要等待的goroutine的数量，Done()用来表示goroutine已经完成了，减少一次计数器，Wait()用来等待所有需要等待的goroutine完成。
4.默认CPU的逻辑核心数，默认跑满整个CPU
	runtime.GOMAXPROCS(4)
*/

func f1(i int) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)))
	fmt.Println(i)

}

var wg sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f1(i)
	}
	wg.Wait()
}
