package main

import (
	"fmt"
)

func main() {
	var countryCapitalMap map[string]string     // 声明hash类型
	countryCapitalMap = make(map[string]string) // 初始化hash   声明+初始化=创建操作

	// hash的赋值操作
	// map 插入key -- value键值对,各个国家对应的首都
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India"] = "新德里"
	countryCapitalMap["American"] = "华盛顿特区"

	// 使用键输出地图值
	for country := range countryCapitalMap {
		// hash的读取操作
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	// 判断指定键对应的值在hash中是否存在,或者:hash中是否包含指定的值
	// 这样去判断一个hash是否包含某个值
	capital, ok := countryCapitalMap["American"]
	// ok是bool类型
	if ok {
		fmt.Println("American的首都是", capital)
		return
	}
	// 注意这种if...else...的结构,一般建议去掉else分支,直接加if进行筛选逻辑,但是要注意在if分支里面务必加return 结束程序
	fmt.Println("American的首都不存在")
	//else {
	//	fmt.Println("American的首都不存在")
	//}
}
