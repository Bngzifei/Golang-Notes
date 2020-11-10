package main

import (
	"fmt"
	"time"
)

/*
time包提供了时间的显示和测量用的函数.日历的计算采用的是公历
time.Time类型表示时间.我们可以通过time.Now()函数获取当前的
时间对象,然后获取时间对象的年月日时分秒等信息.示例代码如下:

*/

//func timeDemo() {
//	now := time.Now() // 获取当前时间
//	fmt.Printf("current time:%v\n", now)
//
//	year := now.Year()     // 年
//	month := now.Month()   // 月
//	day := now.Day()       // 日
//	hour := now.Hour()     // 小时
//	minute := now.Minute() // 分钟
//	second := now.Second() // 秒
//	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
//}

/*
时间戳
时间戳是自1970年1月1日至当前时间的总毫秒数.它也被称为Unix时间戳

使用time.Unix()函数可以将时间戳转为时间格式

时间间隔
time.Duration是time包定义的一个类型,它代表两个时间点之间经过的时间.
以纳秒为单位,time.Duration表示一段时间间隔,可表示的最长时间段大约290年.


*/

func timestampDemo() {
	now := time.Now()            // 获取当前时间
	timestamp1 := now.Unix()     // 时间戳
	timestamp2 := now.UnixNano() // 纳秒时间戳

	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)
}

func main() {
	timestampDemo()
}
