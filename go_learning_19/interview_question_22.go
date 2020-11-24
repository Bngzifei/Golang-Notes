package main

import "fmt"

//func main() {
//	// golang 的字符串类型是不能赋值nil的
//	var x string = nil
//	// 也不能跟nil比较
//	if x == nil {
//		x = "default"
//	}
//	fmt.Println()
//}


var a bool = true

func main() {
	defer func() {
		fmt.Println("1")
	}()
	if a == true {
		fmt.Println("2")
		return
	}
	defer func() {
		fmt.Println("3")
	}()
}

/*
defer 关键字后面的函数或者方法想要执行必须先注册,
return之后的defer是不能注册的,也就不能执行后面的
函数或方法
 */