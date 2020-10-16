// goto语句通过标签进行代码间的无条件跳转,同时goto语句在快速跳出循环,
// 避免重复退出上也有一定的帮助,使用goto语句能简化一些代码的实现过程.

package main

import "fmt"

func main() {
	var breakAgain bool

	// 外循环
	for x := 0; x < 10; x++ {
		// 内循环
		for y := 0; y < 10; y++ {
			// 满足某个条件时,退出循环
			if y == 2 {
				//设置退出标记
				breakAgain = true

				// 退出本次循环
				break

			}
		}

		// 根据标记,还需要退出一次循环
		if breakAgain {
			break
		}
	}
	fmt.Printf("done")
}
