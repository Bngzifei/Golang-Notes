package main

import "fmt"

//func twoSum(nums []int, target int) []int {
//	res := []int{}
//	for index1, value1 := range nums {
//		for index2, value2 := range nums {
//			if (value1 + value2) == target {
//				res = []int{index1, index2}
//				break
//			}
//		}
//	}
//	fmt.Println(res)
//	return res
//}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	res := []int{}
	for index, element := range nums {
		if _, val2 := m[target-element]; val2 {
			fmt.Println(m)
			res = []int{index, m[target-element]}
			break
		}
		// 是在这里,把数组中的元素作为map的key,把数组的index作为map的value.
		// 先进行存储
		// 格式:{2:0,7:1,8:2,9:3...}
		m[element] = index
	}
	fmt.Println(res)
	return res
}

func main() {
	//nums := []int{2, 7, 8, 9, 15, 17}
	//twoSum(nums, 15)

	// 测试
	maps := map[int]int{1:1,2:2,3:3}
	fmt.Println(maps)
	if _, ok := maps[1];ok {
		fmt.Println("exist")
	}

}
