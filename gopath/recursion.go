package main

// 递归,就是在运行的过程中调用自己
// go支持递归,但是需要开发者在使用的时候设置退出条件,否则递归将陷入无限循环中

// 对于解决数学上的问题是非常有用的,就像计算阶乘,生成斐波那契数列
func recursion() {
	recursion() //函数调用自身
}

func main() {
	recursion()
}
