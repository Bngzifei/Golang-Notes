package main

import (
	"errors"
	"fmt"
)

func getCircleArea(radius float32) (area float32, err error) {
	if radius < 0 {
		// 构建一个异常对象
		err = errors.New("半径不能为负")
		return
	}

	area = 3.14 * radius * radius
	return
}

func main() {
	area, err := getCircleArea(-9)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(area)
	}

}
