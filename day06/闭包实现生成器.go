// 闭包的记忆效应进程被用于实现类似于设计模式中工厂模式的生成器.
// 下面的例子展示了创建一个玩家生成器的过程

// 玩家生成器的实现:
package main

import (
	"fmt"
)

// 创建一个玩家生成器,输入名称,输出生成器
func playerGen(name string) func() (string, int) {
	// 血量一直为150
	hp := 150

	// 创建返回的闭包
	return func() (string, int) {
		// 将变量引用到闭包中
		return name, hp
	}
}

func main() {

	// 创建一个玩家生成器
	generator := playerGen("high noon")

	// 返回玩家的名字和血量
	name, hp := generator()

	// 打印值
	fmt.Println(name, hp)
}

// 代码输出如下：
// high noon 150

// 代码说明如下：
// 第 8 行，playerGen() 需要提供一个名字来创建一个玩家的生成函数。
// 第 11 行，声明并设定 hp 变量为 150。
// 第 14～18 行，将 hp 和 name 变量引用到匿名函数中形成闭包。
// 第 24 行中，通过 playerGen 传入参数调用后获得玩家生成器。
// 第 27 行，调用这个玩家生成器函数，可以获得玩家的名称和血量。

// 闭包还具有一定的封装性,第11行的变量是playerGen的局部变量,
// playerGen的外部无法直接访问及修改这个变量,这种特性也和面向对象中强调的封装性类似
