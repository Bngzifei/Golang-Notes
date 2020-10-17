package main

import "fmt"

/*
go语言切片是对数组的抽象
go语言数组的长度不可改变,在特定场景中这样的集合就不太适用,
go中提供了一种灵活,功能强悍的内置类型切片(动态数组),
与数组相比切片的长度是不固定的,可以追加元素,在追加时可能使
切片的容量增大.

可以声明一个未指定大小的数组来定义切片:
var identifier []type

切片不需要说明长度
或使用make()函数来创建切片:
var slice1 []type = make([]type, len)
也可以简写为
slice1 := make([]type, len)

也可以指定容量,其中capacity为可选参数:
make([]T, length, capacity)
这里len是数组的长度并且也是切片的初始长度

切片初始化
	s :=[] int {1,2,3}
直接初始化切片,[] 表示是切片类型,{1,2,3}初始化值依次是1,2,3,
其cap=len=3
	s := arr[:]

通过内置函数make()初始化切片s,[]int标识为其元素类型为int的切片

len() 和 cap() 函数
切片是可索引(就是可以通过下标取值)的,并且可以由len()方法获取长度
切片提供了计算容量的方法cap(),可以测量切片最长可以达到多少
*/

func main() {
	//var numbers = make([]int, 3, 5)
	numbers := []int{2, 5, 7, 8}
	printSlice(numbers)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
