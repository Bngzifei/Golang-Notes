package main

import "fmt"

/*
如果你想向函数传递数组参数,你需要在函数定义时,声明形参为
数组

实例中函数接收整型数组参数,另一个参数指定了数组元素的个数,
并返回平均值.
*/

func main() {
	// 数组长度为5
	var balance = [5]int{1000, 2, 3, 17, 50}
	var avg float32

	// 数组作为参数传递给函数
	avg = getAverage(balance, 5)

	fmt.Printf("平均值为: %f ", avg)
}

func getAverage(arr [5]int, size int) float32 {
	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float32(sum) / float32(size)

	return avg
}
