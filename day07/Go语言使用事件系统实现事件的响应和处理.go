// Go语言可以将类型的方法与普通函数视为一个概念,从而简化方法和函数混合作为回调类型时的复杂性.这个特性和C#中的代理(delegate)类似,调用者无须关心谁来支持调用,系统会自动处理是否调用普通函数或类型的方法.

// 本节中,首先将用简单的例子了解Go语言是如何将方法与函数视为一个概念,接着会实现一个事件系统,事件系统能有效地将事件触发与响应两端代码解耦.

// 方法和函数的统一调用

// 本节的例子将让一个结构体的方法(class.Do)的参数和一个普通函数(funcDo)的参数完全一致,也就是方法与函数的签名一致.然后使用它们签名一致的函数变量(delegate)分别赋值方法与函数,接着调用它们,观察实际效果.

// 详细实现请参考下面的代码:
package main

import (
	"fmt"
)

// 声明一个结构体
type class struct {
}

// 给结构体添加Do方法
func (c *class) Do(v int) {
	fmt.Println("call method do:", v)
}

func funcDo(v int) {
	fmt.Println("call function do:", v)
}

func main() {
	// 声明一个函数回调
	var delegate func(int)

	// 创建结构体实例
	c := new(class)

	// 将回调设为c的Do方法
	delegate = c.Do

	// 调用
	delegate(100)

	// 将回调设为普通函数
	delegate = funcDo

	// 调用
	delegate(100)
}

// 代码说明如下:
// 第10行,为结构体添加一个Do()方法,参数为整型.这个方法的功能是打印提示和输入的参数值

// 第16行,声明一个普通函数,参数也是整型,功能是打印提示和输入的参数值.

// 第24行,声明一个delegate的变量,类型为func(int),与funcDo和class的Do()方法的参数一致.

// 第30行,将c.Do作为值赋给delegate变量

//第33行,调用delegate()函数,传入100的参数,此时会调用c实例的Do()方法
// 第36行,将funcDo赋值给delegate
// 第39行,调用delegate()传入100的参数.此时会调用funcDo()方法.

// 运行代码,输出如下:
// call method do:100
// call function do:100

// 这段代码能运行的基础在于:无论是普通函数还是结构体的方法,只要它们的签名一致,与它们签名一致的函数变量就可以保存普通函数或是结构体方法

// 了解了Go语言的这一特性后,我们就可以将这个特性用在事件中
