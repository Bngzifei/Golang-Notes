// 当一个函数在其函数体内调用自身,则称之为递归.最经典的例子便是计算斐波那契数列,即每个数均为钱两个数之和.

// 数列如下所示：
// 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, …

// 下面的程序可以用于生成该数列:

package main

import (
	"fmt"
)

func main() {
	result := 0
	for i := 0; i < 10; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is:%d\n", i, result)
	}
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

// 提示:在使用递归时,需要设置退出条件,否则递归将陷入无限循环中.

// 许多问题都可以使用优雅的递归来解决.比如说著名的快速排序算法

// 在使用递归函数时经常会遇到的一个问题就是栈溢出:一般出现在大量的递归调用导致的程序栈内存分配耗尽.这个问题可以通过一个名为懒惰求值的技术解决.在Go语言中,我们可以使用管道(channel)和goroutine来实现

// 构成递归需具备的条件:
// 子问题须与原始问题为同样的事,且更为简单.
// 不能无限制地调用本身,须有个出口,化简为非递归状况处理.

// Go语言中也可以使用相互调用的递归函数:多个函数之间相互调用形成闭环.因为Go语言编译器的特殊性,这些函数的声明顺序是可以任意的.下面这个简单的例子展示了函数odd和even之间的相互调用.

package main 

import (
    "fmt"
)

func main() {
    fmt.Printf("%d is even: is %t\n",16,even(16))
    fmt.Printf("%d is odd: is %t\n",17,odd(17))
    fmt.Printf("%d is odd: is %t\n",18,odd(18))
}

func even(nr int) bool {
    if nr == 0 {
        return true     
    }
    return odd(RevSign(nr)-1)
}

func odd(nr int) bool {
    if nr == 0 {
        return false
    }
    return even(RevSign(nr) -1)
}

func RevSign(nr int) int {
    if nr < 0 {
        return -nr
    }
}


