package main

import "fmt"

/*
函数可以返回多个值
*/

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("Google", "Taobao")
	fmt.Println(a, b)
}
