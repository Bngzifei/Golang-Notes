package main

// go语言的switch中的每一个case与case之间是独立的代码块,不需要通过break语句跳出当前case代码块以避免执行到下一行.

import (
	"fmt"
)

func main() {
	var a = "world"

	switch a {

	case "hello":
		fmt.Println(1)
	case "world":
		fmt.Println(2)
	default:
		fmt.Println(0)
	}
}

// 上面的例子中,每一个case均是字符串格式,而且使用了default分支,go语言规定每个switch只能有一个default分支

// 1) 一分支多值
// 当出现多个case要放在一起的时候,可以这么写:
var a = "num"
switch a {
case "mum","daddy":
    fmt.Println("family")
}
// 不同的case表达式使用逗号分隔

// 2)分支表达式
// case后不仅仅只可以是常量,还可以和if一样添加表达式,代码如下;

var r int = 11
switch {
case r > 10 && r < 20:
    fmt.Println(r)
}

// 注意,这种情况的switch后面不再跟判断变量,连判断的目标都没有了

















