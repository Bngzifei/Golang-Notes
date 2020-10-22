package main

import (
	"errors"
	"fmt"
)

/*
defer 与 closure
*/

func foo(a, b int) (i int, err error) {
	defer fmt.Printf("first defer err %v\n", err)
	defer func(err error) { fmt.Printf("second defer err %v\n", err) }(err)
	defer func() { fmt.Printf("third defer err %v\n", err) }()

	if b == 0 {
		err = errors.New("divided by zero!")
		return
	}

	i = a / b
	return
}

func main() {
	foo(2, 0)
}

/*
解释:
如果 defer 后面跟的不是一个closure最后执行的时候我们得到的并不是最新的值
 */