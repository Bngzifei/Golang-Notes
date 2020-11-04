package main

import "fmt"

/*
go 实现类似try ... catch 的异常处理
*/

func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()

	fun()
}

func main() {
	Try(func() {
		panic("test panic")
	}, func(err interface{}) {
		fmt.Println(err)
	})
}

/*
输出结果: test panic

问题: 如何区别使用 panic 和 error 两种方式?
惯例是:导致关键流程出现不可修复性错误的使用panic,其他使用error
 */