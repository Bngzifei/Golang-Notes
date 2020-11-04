package main

import (
	"errors"
	"fmt"
)

/*
除用panic引发中断性错误外,还可以返回error类型错误对象来
表示函数调用的状态.

type error interface {
	Error() string

	标准库errors.New 和 fmt.Errorf函数用于创建实现error接口的错误
对象.通过判断错误对象实例来确定具体的错误类型.
*/
var ErrDivByZero = errors.New("division by zero")

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, ErrDivByZero
	}
	return x / y, nil
}

func main() {
	defer func() {
		fmt.Println(recover())
	}()
	switch z, err := div(10, 0)
	err {
		case nil:
			println(z)
		case ErrDivByZero:
			panic(err)
	}

}

/*
输出结果:  division by zero
 */