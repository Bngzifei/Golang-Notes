package main

import "fmt"

/*
闭包 是由函数及其相关引用环境组合而成的实体(即:闭包=函数+引用环境)

官方的解释:
	所谓闭包,指的是一个拥有许多变量和绑定了这些变量的环境的表达式(通常
是一个函数),因而这些变量也是该表达式的一部分.

闭包:是引用了自由变量的函数.这个被引用的自由变量将和这个函数一起存在,即使
已经离开了创造它的环境也不例外.
所以,有另一种说法认为闭包是由函数和与其相关的引用环境组合而成的实体.
闭包在运行时可以有多个实例,不同的引用环境和相同的函数组合可以产生不同的实例

*/

/*
闭包复制的是原对象指针,这就很容易解释延迟引用现象.
*/

// 外部引用函数参数局部变量
func add(base int) func(int) int {
	return func(i int) int {
		base += 1
		return base
	}
}

func main() {
	tmp1 := add(10)
	fmt.Println(tmp1(1), tmp1(2))
	tmp2 := add(100)
	fmt.Println(tmp2(1), tmp2(2))
}

/*
递归:就是在运行的过程中调用自己.一个函数调用自己,就叫做递归函数.

构成递归需具备的条件:
1.子问题需与原始问题为同样的事,且更为简单
2.不能无限制地调用本身,需有个出口,化简为 非递归状况处理
 */