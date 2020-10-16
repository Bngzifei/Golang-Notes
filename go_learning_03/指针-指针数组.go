package main

import "fmt"

/*
你可以定义一个指针数组来存储地址

有一种情况,我们可能需要保存数组,这样
我们就需要使用到指针.

*/

const MAX int = 3

func main() {
	a := []int{10, 100, 200}
	var i int
	// 3个指针
	var ptr [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i] // 整数地址 赋值给 指针数组
	}

	for i = 0; i < MAX; i++ {
		// 取出指针数组存储的值
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}
}

/*
a[0] = 10
a[1] = 100
a[2] = 200
 */