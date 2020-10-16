package main

import "fmt"

/*
当一个指针被定义后没有分配到任何变量时,它的值为nil
nil指针也称为空指针
nil在概念上和其他语言的null,None,nil,NULL一样,都指代零值
或空值
一个指针变量通常缩写为ptr
*/

func main() {
	var ptr *int
	fmt.Printf("ptr 的值为: %x\n", ptr)
}

/*
空指针判断:
if (ptr != nil)  // ptr 不是空指针
if (ptr == nil)  // ptr 是空指针
 */