// 在go语言中,函数也是一种类型,可以和其他类型一样被保存在变量中. 下面的代码定义了一个安徽省农户变量f,并将一个函数名fire()赋值给函数变量f,这样调用函数变量f时,实际调用的就是fire()函数,代码如下:

package main

import (
	"fmt"
)

func fire() {
	fmt.Println("fire")
}

func main() {
	var f func() //将变量f声明为func()类型,此时f就被俗称为"回调函数".此时f的值为nil

	f = fire

	f()
}

// 代码说明:
// 第7行,定义了一个fire()函数
// 第13行,将变量f声明为func()类型,此时f就被俗称为"回调函数".此时f的值为nil
// 第15行,将fire()函数名作为值,赋给f变量,此时f的值为fire()函数
// 第17行,使用f变量进行函数调用,实际调用的是fire()函数
