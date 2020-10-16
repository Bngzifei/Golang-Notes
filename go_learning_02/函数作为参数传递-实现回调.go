package main

import "fmt"

type callbackfunc func(int) int

func main() {
	testCallBack(1, callBack)
	testCallBack(2, func(x int) int {
		fmt.Printf("我是回调, x: %d\n", x)
		return x
	})
}

func testCallBack(x int, f callbackfunc) {
	f(x)
}

func callBack(x int) int {
	fmt.Printf("我是回调, x: %d\n", x)
	return x
}
