package main

import "fmt"

// 使用延迟匿名函数或下面这样都是有效的

func except() {
	// recover 就是 Python 中的except
	fmt.Println(recover())
}

func test() {
	defer except()
	// panic 就是Python 中的raise
	panic(" test panic")
}

func main() {
	test()
}

// 输出结果: test panic
