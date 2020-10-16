package main

/*
使用fallthrough会强制执行后面的case语句,fallthrough不会判断
下一条case的表达式结果是否为true

从以上代码输出的结果可以看出：switch 从第一个判断表达式为 true 的
case 开始执行，如果 case 带有 fallthrough，
程序会继续执行下一条 case，且它不会去判断下一个 case
的表达式是否为 true。
*/

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("1、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	case false:
		fmt.Println("3、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("4、case 条件语句为 true")
	case false:
		fmt.Println("5、case 条件语句为 false")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}
}
