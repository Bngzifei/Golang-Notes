package main

import "fmt"

/*
switch 语句还可以被用于 type-switch 来判断某个
interface 变量中实际存储的变量类型。
*/

func main() {
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf("x 的类型 : %T", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或者 string 型")
	default:
		fmt.Printf("未知类型")
	}
}
