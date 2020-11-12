package main

import "fmt"
var fg string  // 全局变量的声明
// fg1 := "sdss"  // 全局变量的声明和赋值,注意:不允许全局变量的赋值
//rule := "Short variable declarations"  // syntax error: non-declaration statement outside function body

func func1() {
	i := 0
	// := 左侧新增了一个 j 变量进行声明,这样是允许的
	i, j := 1, 2
	fmt.Printf("i = %d, j = %d\n", i, j)
}

func func2(i int) {
	// no new variables on left side of :=
	// 出现这种错误,说明是 := 的左侧没有新的声明的变量,两次声明的是同样的变量
	// 所以,下面直接使用赋值 = 号即可
	// 不能通过编译原因是形参已经声明了变量i，使用:=再次声明是不允许的
	i = 0
	fmt.Println(i)
}

// j = 1, k = 1
// i = 0, j = 0
// 这里要注意的是,block if 中声明的j,与上面的j属于不同的作用域
func func3() {
	i, j := 0, 0
	if true {
		j, k := 1, 1
		fmt.Printf("j = %d, k = %d\n", j, k)
	}
	fmt.Printf("i = %d, j = %d\n", i, j)
}

func main() {
	func3()
}

/*
当 := 左侧存在新变量声明时(如field2),那么已声明的变量如(offset)则会被重新声明,
不会有其他额外副作用.

当 := 左侧没有新变量是不允许的,编译会提示 no new variable on left side of :=

我们所说的重新声明不会引入问题要满足一个前提,变量声明要在同一个作用域中出现.如果
出现在不同的作用域,那很可能就创建了新的同名变量,同一函数不同作用域的同名变量往往
不是预期做法,很容易引入缺陷.

简短变量场景只能用于函数中,使用 := 来声明和初始化全局变量是行不通的.
比如,像下面这样:
	rule := "Short variable declarations"

这里的编译错误提示: syntax error: non-declaration statement outside function body
表示非声明语句不能出现在函数外部.可以理解成:= 实际上会拆分成两个语句,即声明和赋值.赋值语句
不能出现在函数外部的.

变量作用域的问题:
	几乎所有的工程师都了解变量作用域,但是由于 := 使用过于频繁的话,还是有可能
掉进陷阱里.

*/