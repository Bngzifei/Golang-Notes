package main

import "fmt"

/*
range关键字用于for循环迭代数组(array),切片(slice),
通道(channel)或集合(map)的元素.在数组和切片中它
返回元素的索引和索引对应的值,在集合中返回key-value对.
*/

func main() {
	// 这是我们使用range去求一个slice的和。使用数组跟这个很类似
	nums := [] int{1, 2, 3}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// 在数组上使用range将传入index和值两个变量.
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	// range 也可以用在map的键值对上
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range 也可以用来枚举Unicode字符串。第一个参数是字符串索引,
	// 第二个是字符(Unicode的值)本身.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

/*
sum: 6
index: 2
a -> apple
b -> banana
0 103
1 111
 */