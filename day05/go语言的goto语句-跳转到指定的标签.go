// goto语句通过标签进行代码间的无条件跳转.goto语句可以在快速跳出循环,避免重复退出上有一定的帮助.go语言中使用goto语句能简化一些代码的实现过程

// 使用goto退出多层循环
// 下面这段代码在满足条件时，需要连续退出两层循环，使用传统的编码方式如下：

package main

import (
	"fmt"
)

func main() {
	var breakAgain bool

	// 外循环
	for x := 0; x < 10; x++ {
		// 内循环
		for y := 0; y < 10; y++ {
			// 满足某个条件时,退出循环
			if y == 2 {
				breakAgain = true
				// 退出本次循环
				break
			}
		}
		//  根据标记,还需要退出一次循环
		if breakAgain {
			break
		}

	}
	fmt.Println("done")

}

// 上述逻辑使用goto语句进行优化

package main

import "fmt"

func main() {
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if y == 2 {
				// 跳转到标签
				goto breakHere

			}
		}
	}
	// 手动返回,避免执行进入标签
	return
	// 标签
breakHere:
	fmt.Println("done")
}
// 使用goto语句后,无须额外的变量就可以快速退出所有的循环.

// 统一错误处理
// 多处错误处理存在代码重复时非常棘手的,例如:
err := firstCheckError()
if err != nil {
    fmt.Println(err)
    exitProcess()
    return
}

err = secondCheckError()

if err != nil {
    fmt.Println(err)
    exitProcess()
    return
}

fmt.Println("done")
// 这样的逻辑很繁琐,不好.


// 如果使用goto语句来实现同样的逻辑:
    err := firstCheckError()
    if err != nil {
        goto onExit
    }

    err = secondCheckError()

    if err != nil {
        goto onExit
    }

    fmt.Println("done")

    return

onExit:
    fmt.Println(err)
    exitProcess()


// 可以看到,在遇到错误逻辑的时候,goto执行跳转到错误处理的标签处