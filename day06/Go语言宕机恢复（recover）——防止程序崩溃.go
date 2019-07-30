// Recover是一个Go语言的内建函数,可以让进入宕机流程中的goroutine恢复过来.recover仅在延迟函数defer中有效.在正常的执行过程中,调用recover会返回nil并且没有其他任何效果.如果当前的goroutine陷入恐慌(就是宕机了),调用recover可以捕获到panic的输入值,并且恢复正常的执行.

// recover函数用于终止错误处理流程.一般情况下,recover应该在一个使用defer关键字的函数中执行以有效截取错误处理流程.如果没有在发生异常的goroutine中明确调用恢复过程(使用recover关键字),会导致该goroutine所属的进程打印异常信息后直接退出.

// 通常来说,不应该对进入panic宕机的程序做任何处理,但有时,也许我们可以从宕机中恢复,至少我们可以在程序崩溃前,做一些操作.举个例子,当web服务器遇到不可预料的严重问题时,在崩溃前应该将所有的连接关闭,如果不做任何处理,会使得客户端一直处于等待状态.如果web服务器还在开发阶段,服务器甚至可以将异常信息反馈到客户端,帮助调试.

// 如果在deferred函数中调用了内置函数recover,并且定义该defer语句的函数发生了panic异常,recover会使程序从panic中恢复,并返回panic value.导致panic异常的函数不会继续运行,但能正常返回.在未发生panic时调用recover,recover会返回nil

// 提示:
// 在其他语言里,宕机往往以异常的形式存在.底层抛出异常,上层逻辑通过try/catch机制捕获异常,没有被吧捕获的严重异常会导致宕机.捕获的异常可以被忽略,让代码继续运行.

// Go没有异常系统,其使用panic触发宕机类似于其他语言的抛出异常,那么recover的宕机恢复机制就对应try/catch机制.

// 让程序在崩溃时继续执行
// 下面的代码实现了ProtectRun()函数,该函数传入一个匿名函数或闭包后的执行函数,当传入函数以任何形式发生panic崩溃后,可以将崩溃发生的错误打印出来,同时允许后面的代码继续执行,不会造成整个进程的崩溃.

// 保护运行函数:

package main

import (
	"fmt"
	"runtime"
)

// 崩溃时需要传递的上下文信息
type panicContext struct {
	function string //所在函数
}

// 保护方式允许一个函数
func ProtectRun(entry func()) {

	// 延迟处理的函数
	defer func() {
		// 发生宕机时获取panic传递的上下文并打印
		err := recover()

		switch err.(type) {
		case runtime.Error: //运行时错误
			fmt.Println("runtime error:", err)
		default: //非运行时的错误
			fmt.Println("error:", err)
		}

	}()

	entry()

}

func main() {
	fmt.Println("运行前")

	// 允许一段手动触发的错误
	ProtectRun(func() {

		fmt.Println("手动宕机前")

		// 使用panic传递上下文
		panic(&panicContext{
			"手动触发panic",
		})

		fmt.Println("手动宕机后")
	})

	// 故意造成空指针访问错误
	ProtectRun(func() {
		fmt.Println("赋值宕机前")

		var a *int
		*a = 1

		fmt.Println("赋值宕机后")

	})

	fmt.Println("运行后")

}

// 代码输出结果：
// 运行前
// 手动宕机前
// error: &{手动触发panic}
// 赋值宕机前
// runtime error: runtime error: invalid memory address or nil pointer
// dereference
// 运行后

// 对代码的说明：
// 第 9 行声明描述错误的结构体，成员保存错误的执行函数。
// 第 17 行使用 defer 将闭包延迟执行，当 panic 触发崩溃时，ProtectRun() 函数将结束运行，此时 defer 后的闭包将会发生调用。
// 第 20 行，recover() 获取到 panic 传入的参数。
// 第 22 行，使用 switch 对 err 变量进行类型断言。
// 第 23 行，如果错误是有 Runtime 层抛出的运行时错误，如空指针访问、除数为 0 等情况，打印运行时错误。
// 第 25 行，其他错误，打印传递过来的错误数据。
// 第 44 行，使用 panic 手动触发一个错误，并将一个结构体附带信息传递过去，此时，recover 就会获取到这个结构体信息，并打印出来。
// 第 57 行，模拟代码中空指针赋值造成的错误，此时会由 Runtime 层抛出错误，被 ProtectRun() 函数的 recover() 函数捕获到。

// panic和recover的关系
// panic和defer的组合有如下特性:
// 1)有panic没recover,程序宕机.
// 2)有panic也有recover捕获,程序不会宕机.执行完对应的defer后,从宕机点退出当前函数后继续执行.

// 提示:
// 虽然panic/recover能模拟其他语言的异常机制,但并不建议编写普通函数也经常性使用这种特性.
// 如果想在捕获错误时设置当前函数的返回值,可以对返回值使用命名返回值方式直接进行设置
