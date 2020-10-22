package main

/*
多个defer注册,按 FILO 次序执行(先进后出).哪怕函数或某个延迟调用
发送错误,这些调用依旧会被执行.
 */

func test(x int) {
	defer println("a")
	defer println("b")

	defer func() {
		println(100 / x)
	}()

	defer println("c")
}

func main() {
	test(0)
}