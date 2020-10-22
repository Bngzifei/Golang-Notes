package main

/*
golang 没有结构化异常,使用panic(应急,惊慌)抛出错误,
recover(恢复,弥补)捕获错误.

异常的使用场景简单描述:go中可以抛出一个panic的异常,然后
在defer中通过recover捕获这个异常,然后正常处理

panic:
	1.内置函数
	2.假如函数F中书写了panic语句,会终止其后要执行的代码,
在panic所在的函数F内如果存在要执行的defer函数列表,按照defer的
逆序执行
	3.返回函数F的调用者G,在G中,调用函数F语句之后的代码不会执行,
假如函数G中存在要执行的defer函数列表,按照defer的逆序执行
	4.直到goroutine整个退出,并报告错误

recover:
	1.内置函数
	2.用来控制一个goroutine的panicking行为,从而影响应用的行为
	3.一般的调用建议:
		a.在defer函数中,通过recever来终止一个goroutine的panicking
过程,从而恢复正常代码的执行
		b.可以获取通过panic传递的error

注意:
	1.利用 recover处理panic指令,defer必须放在panic之前定义,另外 recover
	只有在 defer 调用的函数中才有效.否则当 panic 时, recover 无法捕获
到 panic ,无法防止 panic 扩散
	2. recover 处理异常后,逻辑并不会恢复到 panic 那个点去,函数跑到 defer
之后的那个点.
	3. 多个defer 会形成 defer 栈,后定义的defer 语句会被最先调用.
*/

func main() {
	test()
}

func test() {
	// 延迟调用 函数
	defer func() {
		if err := recover(); err != nil {
			println(err.(string)) // 将interface{} 转型为具体类型
		}
	}() // 写成这样就是 函数调用
	panic("panic error!")
}

/*
panic error!

由于 panic,recover 参数类型为interface{},因此 可抛出任何类型对象


 */