package main

import "fmt"

func getType(s Shape) {
	/*
		方法一:
			instance,ok := 接口对象.(实际类型)
				如果该接口对象是对应的实际类型,那么instance就是转型
				之后的对象,ok的值为true
	*/

	if ins, ok := S.(Triangle); ok {
		fmt.Println("是三角形,三边是: ", ins.a, ins.b, ins.c)
	} else if ins, ok := s.(Circle); ok {
		fmt.Println("是圆形,半径是: ", ins.radius)
	} else {
		fmt.Println("我也不知道了...")
	}
}

func getType2(s Shape) {
	/*
		方法二:接口对象.(type),配合switch 和 case 语句使用
	*/
	switch ins := s.(type) {
	case Triangle:
		fmt.Println("三角形...", ins.a, ins.b, ins.c)
	case Circle:
		fmt.Println("圆形...", ins.radius)

	}
}

func main() {
	// 接口类型的对象 --> 对应实现类 类型
	getType(c1)
	getType2(t1)
}
