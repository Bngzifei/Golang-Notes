package main

import "fmt"
// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
/*
示例1:
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 */
func lengthOfLongestSubstring(s string) int {
	// 哈希集合,记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针,初始值为-1,相当于我们在字符串的左边界的左侧,还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 右指针往右移动一格,移除一个字符
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第i到rk个字符是一个很长的无重复字符子串
		ans = max(ans, rk-i+1)
	}
	fmt.Println(ans)
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	s := "abcdrttuyua"
	lengthOfLongestSubstring(s)
}
