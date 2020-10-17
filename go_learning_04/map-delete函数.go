package main

import "fmt"

/*
delete()函数用于删除集合的元素,参数为map和其对应的key.
*/

func main() {
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}

	fmt.Println("原始地图")

	// 打印地图
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	// 删除元素
	delete(countryCapitalMap, "France")
	fmt.Println("法国条目被删除")
	fmt.Println("删除后的地图")

	// 打印地图
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

}

/*
原始地图
France 首都是 Paris
Italy 首都是 Rome
Japan 首都是 Tokyo
India 首都是 New delhi
法国条目被删除
删除后的地图
Italy 首都是 Rome
Japan 首都是 Tokyo
India 首都是 New delhi
 */