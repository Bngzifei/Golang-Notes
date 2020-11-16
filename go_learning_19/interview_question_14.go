package main

import "fmt"

//func main() {
//	str := "hello"
//	// cannot assign to str[0]
//	// go 语言中的字符串是只读的
//	str[0] = 'x'
//	fmt.Println(str)
//}

func incr(p *int) int {
	*p++
	return *p
}

func main() {
	p := 1
	incr(&p)
	fmt.Println(p)
}