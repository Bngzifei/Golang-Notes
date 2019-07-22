// go 语言切片是对数组的抽象
// go 数组的长度不可改变,在特定场景中这样的集合就不太适用.
// go 中提供了一种灵活的,功能强悍的内置类型切片(动态数组).
// 与数组相比切片的长度不是固定的,可以追加元素,在追加时可能是切片的容量变大

package main

import (
	"fmt"
)

func main() {
	var numbers = make([]int, 3, 5)
	printSlice(numbers)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)

}
