package main

import "fmt"

//func main() {
//	a := [5]int{1,2,3,4,5}
//	t := a[3:4:4]
//	// 4
//	fmt.Println(t[0])
//}

/*
操作符[i:j]
假设底层数组的大小为k,截取之后获得的切片的长度和容量的计算
方法: 长度: j-i,第三个参数k用来限制新切片的容量,但不能超过
原数组(切片)的底层数组大小.截取获得的切片的长度和容量分别是
j-i, k-i
所以例子中,切片t为[4],长度和容量都是1

 */

func main() {
	a := [2]int{5,6}
	b := [3]int{5,6}
	if a == b {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
}

/*
invalid operation: a == b (mismatched types [2]int and [3]int)
Go中的数组是值类型,可比较,另外一方面,数组的长度也是数组类型的组成部分.
所以a和b是不同的类型,是不能比较的,所以编译错误.
 */