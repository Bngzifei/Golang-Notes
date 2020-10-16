package main

import (
	"fmt"
	"math"
)

/*
函数定义后可作为另外一个函数的实参数传入

go可以很灵活地创建函数,并作为另外一个函数的实参.
以下实例中我们在定义的函数中初始化一个变量,该
函数仅仅是为了使用内置函数math.sqrt()
*/

func main() {
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getSquareRoot(169))
}
