package main

import (
	"fmt"
	"strconv"
	"sync"
)

/*
go语言中内置的map不是并发安全的.请看下面的示例:
*/

//var m = make(map[string]int)
//
//func get(key string) int {
//	return m[key]
//}
//
//func set(key string, value int) {
//	m[key] = value
//}
//
//func main() {
//	wg := sync.WaitGroup{}
//	for i := 0; i < 20; i++ {
//		wg.Add(1)
//		go func(n int) {
//			key := strconv.Itoa(n)
//			set(key, n)
//			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
//			wg.Done()
//		}(i)
//	}
//	wg.Wait()
//}
/*
上面的代码开启少量几个goroutine的时候可能没什么问题,当并发多了之后
执行上面的代码就会报 fatal error:concurrent map writes错误

像这种场景下就需要为map加锁来保证并发的安全性了,go语言的sync包中提供
了一个开箱即用的并发安全版map-sync.Map.开箱即用表示不用像内置的map一样
使用make函数初始化就能直接使用.同时sync.Map内置了诸如Store,Load,LoadOrStore
.Delete.Range等操作方法
 */

var m = sync.Map{}

func main() {
	wg := sync.WaitGroup{}
	for i :=0; i < 20; i ++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key,n)
			value, _ := m.Load(key)
			fmt.Printf("k=:%v,v:=%v\n",key,value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}


