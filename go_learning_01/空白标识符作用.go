package main

import "fmt"

func main() {
	// 实际调用的时候只使用两个
	_, numb, strs := numbers() //只获取函数返回值的后两个
	fmt.Println(numb, strs)
}

//一个可以返回多个值的函数
func numbers() (int, int, string) {
	a, b, c := 1, 2, "str"
	// 返回三个值
	return a, b, c
}
