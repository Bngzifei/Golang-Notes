package main

import "fmt"

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	o := 0
	// 双指针,指针就是这里数组的索引位置.
	i, j := 0, len(height)-1
	// i != j 就是说左右两个指针没有碰到之前的时候
	// 死循环 for ...
	for i != j {
		// height 是一个数组
		hi, hj := height[i], height[j]
		// j - i 就是底  求面积的时候,需要知道最小的高, 底 * 高 = 面积
		s := (j - i) * min(hi, hj)
		fmt.Println("对应的两个高是:", hi, hj, "此时面积是:", s)
		// 判断最小面积是否大于0
		if s > o {
			o = s
		}
		// 如果左侧的高大于右侧
		if hi > hj {
			// 右指针向左移
			j--
		} else {
			// 否则 左指针向右移
			i++
		}
	}
	return 0
}

func main() {
	height := []int{1, 2, 3, 6, 4, 5, 9}
	maxArea(height)
}
