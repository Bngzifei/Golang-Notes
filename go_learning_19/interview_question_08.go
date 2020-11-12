package main

/*
一个包中,可以包含多个init函数
程序编译时,先执行依赖包的init函数,再执行main包内的init函数

关于init()函数有几个需要注意的地方:
init()函数是用于程序执行前做包的初始化的函数,比如初始化包里
的变量等.
一个包可以出现多个init()函数,一个源文件也可以包含多个init()
函数.
同一个包中多个init()函数的执行顺序没有明确定义,但是不同包的init
函数是根据导入的依赖关系决定的.
init()函数在代码中不能被显示调用,不能被引用(赋值给函数变量),否则
出现编译错误.
一个包被引用多次,如 A import B,C import B, A import C, B被
引用多次,但B包只会初始化一次
引入包,不可以出现死循环.即 A import B, B import A, 这种情况编译失败

*/

//func hello() []string {
//	return nil
//}
//
//func main() {
//	h := hello
//	if h == nil {
//		fmt.Println("nil")
//	} else {
//		fmt.Println("not nil")
//	}
//}

func GetValue() int {
	return 1
}

func main() {
	i := GetValue()
	// cannot type switch on non-interface value i (type int)
	// 类型选择的语法: i.(type),其中i是接口,type是固定关键字,
	// 需要注意的是,只有接口类型才可以使用类型选择.
	switch i.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}
}


/*
编译失败.
 */