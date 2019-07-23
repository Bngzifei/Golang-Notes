// 被捕获到闭包中的变量让闭包本身拥有了记忆效应,闭包中的逻辑可以修改闭包捕获的变量,变量会跟随闭包生命期一直存在,闭包本身就如同变量一样拥有了记忆效应

// 累加器的实现:

package main

import (
	"fmt"
)

// 提供一个值,每次调用函数会指定对值进行累加
func Accumulate(value int) func() int {
	// 返回一个闭包
	return func() int {

		// 累加
		value++

		// 返回一个累加值
		return value
	}

}

func main() {

	// 创建一个累加器,初始值为1
	accumulator := Accumulate(1)

	// 累加1并打印
	fmt.Println(accumulator())

	fmt.Println(accumulator())

	// 打印累加器的函数地址
	fmt.Printf("%p\n", accumulator)

	// 创建一个累加器,初始值为1
	accumulator2 := Accumulate(100000000)

	// 累加1并打印
	fmt.Println(accumulator2())

	//打印累加器的函数地址
	fmt.Printf("%p\n", accumulator2)

}

// 代码说明如下：
// 第 8 行，累加器生成函数，这个函数输出一个初始值，调用时返回一个为初始值创建的闭包函数。
// 第 11 行，返回一个闭包函数，每次返回会创建一个新的函数实例。
// 第 14 行，对引用的 Accumulate 参数变量进行累加，注意 value 不是第 11 行匿名函数定义的，但是被这个匿名函数引用，所以形成闭包。
// 第 17 行，将修改后的值通过闭包的返回值返回。
// 第 24 行，创建一个累加器，初始值为 1，返回的 accumulator 是类型为 func()int 的函数变量。
// 第 27 行，调用 accumulator() 时，代码从 11 行开始执行匿名函数逻辑，直到第 17 行返回。
// 第 32 行，打印累加器的函数地址。

// 对比输出的日志发现 accumulator 与 accumulator2 输出的函数地址不同，因此它们是两个不同的闭包实例。

// 每调用一次 accumulator 都会自动对引用的变量进行累加。
