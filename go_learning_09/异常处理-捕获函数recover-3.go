package main

import "fmt"

// 如果需要保护代码段,可将代码块重构成匿名函数,
// 如此可确保后续代码被执行
func test(x, y int) {
	var z int
	func() {
		defer func() {
			if recover() != nil {
				z = 0
			}
		}()
		panic("test panic")
		z = x / y
		return
	}()

	fmt.Printf("x / y = %d\n", z)
}

func main() {
	test(2, 1)
}

/*
输出结果: x / y = 0
*/