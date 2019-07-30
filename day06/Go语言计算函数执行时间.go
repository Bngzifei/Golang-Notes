// 有时候,能够知道一个计算执行消耗的时间是非常有意义的.尤其是在对比和基准测试中.最简单的一个办法就是在计算开始之前设置一个起始时刻,再由计算结束时的结束时间,最后取出它们的差值,就是这个计算所消耗的时间.想要实现这样的做法,可以使用time包中的Now()和Sub()函数

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	test()
	end := time.Now()
	result := end.Sub(start)
	fmt.Printf("该函数执行完成耗时:%s\n", result)
}

func test() {
	sum := 0
	for i := 0; i < 10000000; i++ {
		sum += i
	}
}
