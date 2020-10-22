package main

import "fmt"

/*
defer延迟调用:
特性:
	关键字defer用于注册延迟调用
	这些调用直到return 前才被执行.因此,可以用来做资源清理
	多个defer语句,按先进后出的方式执行
	defer语句中的变量,在defer声明时就决定了

defer 用途:
	1.关闭文件句柄
	2.锁资源释放
	3.数据库连接释放

go 语言的defer功能强大,对于资源管理非常方便

defer是先进后出
这个很自然,后面的语句会依赖前面的资源,因此如果前面的资源先释放了,后面的
语句就没法执行了.
*/

type Test struct {
	name string
}

func (t *Test) Close() {
	fmt.Println(t.name," closed")
}

//func main() {
//	ts := []Test{{"a"},{"b"},{"c"}}
//	for _, t := range ts {
//		defer t.Close()
//	}
//}

func Close(t Test) {
	t.Close()
}

func main() {
	ts := []Test{{"a"}, {"b"}, {"c"}}
	for _, t := range ts{
		// 先进后出
		defer Close(t)
	}
}
/*
c  closed
c  closed
c  closed

这个输出并不会像我们预计的输出 c b a,而是输出 c c c


defer后面的语句在执行的时候,函数调用的参数会被保存起来,但是
不执行.也就是复制了一份,但是并没有说struct这里的this指针如何处理.
go语言并没有把这个明确写出来的this指针当做参数来看待.

 */
