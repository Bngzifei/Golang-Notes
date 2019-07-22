package main

import (
	"fmt"
)

// 在go语言中case是一个独立的代码块,执行结束之后不会像c语言那样紧接着下一个case执行.但是为了兼容一些移植代码,依然加入了fallthrough关键字来实现这一功能(像c语言那样紧接着下一个case执行)

func main() {
	var s = "hello"
	switch {
	case s == "hello":
		fmt.Println("hello")
		fallthrough
	case s != "world":
		fmt.Println("world")
	}
}

// 这样就实现了在前面的case执行结束之后继续执行下一个case的分支情况
