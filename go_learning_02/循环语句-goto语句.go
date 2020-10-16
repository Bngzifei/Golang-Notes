package main

import "fmt"

/*
goto语句可以无条件地转移到过程中指定的行
goto语句通常与条件语句配合使用.可用来实现条件转移,
构成循环,跳出循环体等功能.
但是,在结构化程序设计中一般不主张使用goto语句,
以免造成程序流程的混乱,使理解和调试程序都产生困难.
*/

func main() {
	var a int = 10

LOOP:
	for a < 20 {
		if a == 15 {
			/* 跳过迭代 */
			a = a + 1
			goto LOOP
		}
		fmt.Printf("a的值为 : %d\n", a)
		a++
	}
}
