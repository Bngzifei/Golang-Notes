package main

import "fmt"

/*
每个类型都有与之关联的方法集,这会影响到接口实现规则.

用实例value和pointer调用方法(含匿名字段不受方法集约束,编译器
总是查找全部方法,并自动转换receiver实参)

go语言中内部类型方法集提升的规则:

类型T方法集包含全部receiver T方法

*/

type T struct {
	int
}

func (t T) test() {
	fmt.Println("类型 T 方法集包含全部 receiver T 方法.")
}

func main() {
	t1 := T{1}
	fmt.Printf("t1 is : %v\n", t1)
	t1.test()
}

/*
输出结果:
t1 is : {1}
类型 T 方法集包含全部 receiver T 方法.

给定一个结构体类型S和一个命名为T的类型,方法提升像下面规定的这样被包含在
结构体方法集 中:
	如类型S包含匿名字段T,则S和*S方法集包含T方法
这条规则说的是当我们嵌入一个类型,嵌入类型的接受者为值类型的方法将被提升,
可以被外部类型的值和指针调用.
 */
