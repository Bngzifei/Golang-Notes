package main

import "fmt"

/*
如果想增加切片的容量,我们必须创建一个新的更大的切片并把
原切片的内容都拷贝过来.
下面的代码描述了从拷贝切片的copy方法和向切片追加新元素
的append方法
*/

func main() {
	var numbers []int
	printSlice(numbers)

	// 允许追加空切片
	numbers = append(numbers, 0)
	printSlice(numbers)

	// 向切片添加一个元素
	numbers = append(numbers, 1)
	printSlice(numbers)

	// 同时添加多个元素
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	// 创建切片 numbers1 是之前切片的两倍容量
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	// 拷贝numbers 的内容 到numbers1
	copy(numbers1, numbers)
	printSlice(numbers1)

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

/*
len=0 cap=0 slice=[]
len=1 cap=1 slice=[0]
len=2 cap=2 slice=[0 1]
len=5 cap=6 slice=[0 1 2 3 4]
len=5 cap=12 slice=[0 1 2 3 4]
 */