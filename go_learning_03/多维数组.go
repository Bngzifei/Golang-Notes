package main

import "fmt"

/*
支持多维数组，以下为常用的多维数组声明方式：
例如:   var threedim [5][10][4]int

初始化二维数组:
多维数组初始化:

*/

func main() {
	var a = [5][2] int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	var i, j int

	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}
