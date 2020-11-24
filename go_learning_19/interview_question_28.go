package main

import "fmt"

//func main() {
//	// go 中不同类型是不能比较的.而数组长度是数组类型的一部分,
//	// 所以[...]int{1}和[2]int{1}是两种不同的类型,不能比较.
//	// 切片是不能比较的
//	fmt.Println([...]int{1} == [2]int{1})
//	fmt.Println([]int{1} == []int{1})
//}

var p *int

func foo() (*int, error) {
	var i int = 5
	return &i, nil
}

func bar() {
	fmt.Println(*p)
}

func main() {
	p, err := foo()
	if err != nil {
		fmt.Println(err)
		return
	}
	bar()
	fmt.Println(*p)
}
