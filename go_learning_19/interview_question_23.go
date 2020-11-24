package main

//func main() {
//	s1 := []int{1, 2, 3}
//	s2 := s1[1:] // 2,3
//	s2[1] = 4   // 2,4  说明s2替换元素会影响到s1
//	fmt.Println(s1)
//	s2 = append(s2, 5, 6, 7)  // s2扩容不会影响到s1
//	fmt.Println(s1)
//	fmt.Println(s2)
//}

/*
golang 中切片底层的数据结构是数组.当使用s1[1:]获得切片s2,
和s1共享同一个底层数组,这会导致s2[1]=4语句影响s1,
而append操作会导致底层数据扩容,生成新的数组,因此追加数据后的
s2不会影响s1.
*/

func main() {
	if a := 1; false {
	} else if b := 2; false {

	} else {
		println(a, b)
	}
}
