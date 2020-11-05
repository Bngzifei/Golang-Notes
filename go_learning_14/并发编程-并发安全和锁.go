package main

import (
	"fmt"
	"sync"
)

/*
有时候在go代码中可能会存在多个goroutine同时操作一个资源
(临界区),这种情况会发生竞态问题(数据竞态).类比现实生活中
的例子,有十字路口被各个方向的汽车竞争,还有火车上的卫生间被
车厢里的人竞争.
*/

var x int64
var wg sync.WaitGroup
var lock sync.Mutex


func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock()   // 加锁
		x = x + 1
		lock.Unlock()  // 解锁
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}

/*
上面的代码中我们开启了两个goroutine去累加变量x的值,这两个goroutine
在访问和修改x变量的时候就会存在数据竞争,导致最后的结果与期待的不符.

互斥锁是一种常用的控制共享资源访问的方法,它能够保证同时只有一个goroutine
可以访问共享资源.go语言中使用sync包的Mutex类型来实现互斥锁.
使用互斥锁来修复上面代码的问题.

使用互斥锁能够保证同一时间有且只有一个goroutine进入临界区,其他的goroutine
则在等待锁;当互斥锁释放后,等待的goroutine才可以获取锁进入临界区,多个goroutine
同时等待一个锁时,唤醒的策略是随机的.
 */
