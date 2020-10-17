package main

import "fmt"

/*
一个切片在未初始化之前默认为nil,长度为0
*/

func main() {
	var numbers []int

	printSlice(numbers)

	if (numbers == nil) {
		fmt.Printf("切片是空的")
	}
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

/*
len=0 cap=0 slice=[]
切片是空的
 */