package main

import "fmt"

/*
在指针类型前面加上 * 号（前缀）来获取指针所指向的内容。
*/

func main() {
	var a int = 20 // 声明实际变量
	var ip *int    // 声明指针变量

	ip = &a // 指针变量的存储地址 就是 a 的内存地址

	// 指针变量的存储地址
	fmt.Printf("a 变量的地址是: %x\n", &a)

	// 指针变量的存储地址
	fmt.Printf("ip 变量存储的指针地址: %x\n", ip)

	// 使用指针访问值
	fmt.Printf("ip 变量的值: %d\n", *ip)
}

/*
a 变量的地址是: c0000120b8
ip 变量存储的指针地址: c0000120b8
ip 变量的值: 20
*/
