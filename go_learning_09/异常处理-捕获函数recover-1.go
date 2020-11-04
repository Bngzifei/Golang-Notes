package main

import "fmt"

func test() {
	defer func() {
		fmt.Println(recover())  // 有效
	}()

	defer recover()  // 无效
	defer fmt.Println(recover())

	defer func() {
		func() {
			println("defer inner")
			recover()
		}()
	}()

	panic("test panic")
}

func main() {
	test()
}

/*
defer inner
<nil>
test panic
 */