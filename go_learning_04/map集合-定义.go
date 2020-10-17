package main

import "fmt"

/*
Map是一种无序的键值对的集合, Map最重要的一点
是通过key来快速检索数据,key类似于索引,指向
数据的值.

Map是一种集合,所以我们可以像迭代数组和切片那样迭代它.
不过,Map是无序的,我们无法决定它的返回顺序,这是
因为Map是使用Hash表来实现的.

可以使用内建函数 make 也可以使用map关键字来定义 Map:

如果不初始化map,那么就会创建一个nil map.nil map
不能用来存放键值对
*/

func main() {
	// 创建集合
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	// map 插入key-value对,各个国家对应的首都
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India"] = "新德里"

	// 使用键输出地图值
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	// 查看元素在集合中是否存在
	capital, ok := countryCapitalMap["American"]
	// 如果确定是真实的,则存在,否则不存在
	//fmt.Println(capital)  // 空
	//fmt.Println(ok)  // false
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}

}

/*
France 首都是 巴黎
Italy 首都是 罗马
Japan 首都是 东京
India 首都是 新德里
American 的首都不存在
 */