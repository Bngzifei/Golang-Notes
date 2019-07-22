// Go语言中continue语句可以结束当前循环,开始下一次的循环迭代过程,仅限在for循环内使用.在continue语句后添加标签时,表示开始标签对应的循环.例如:

package main

import "fmt"

func main() {
OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				continue OuterLoop
			}
		}
	}
}
