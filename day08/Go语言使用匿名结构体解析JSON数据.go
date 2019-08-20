// Go语言使用匿名结构体解析JSON数据

// JavaScript对象表示法(JSON)是一种用于发送和接收结构化信息的标准协议.在类似的协议中,JSON并不是唯一的一个标准协议,XML,asn.1和Goole的Protocol Buffers都是类似的协议,并且有各自的特色,但是由于简洁性,可读性和流行程度等原因,JSON是应用最广泛的一个.

// Go语言对于这些标准格式的编码和解码都要良好的支持,由标准库中的encoding/
// json,encoding/xml,encoding/asn1 等包提供支持,并且这类包都有着相似的API接口.

// 基本的JSON类型有数字(十进制或科学计数法),布尔值(true或false),字符串,其中字符串是以双引号包含的Unicode字符序列,支持和Go语言类似的反斜杠转义特性,不过JSON使用的是\Uhhh转义数字类表示一个UTF-16编码,而不是Go语言的rune类型

// 手机拥有屏幕,电池,指纹识别等信息,将这些信息填充为JSON格式的数据.如果需要选择性地分离JSON中的数据则比较麻烦.Go语言中的匿名结构体可以方便地完成这个操作.

package main

import (
	"encoding/json"
	"fmt"
)

// 定义手机屏幕
type Screen struct {
	Size       float32 //屏幕尺寸
	ResX, ResY int     //屏幕水平和垂直分辨率
}

// 定义电池
type Battery struct {
	Capacity int //容量
}

// 生成json数据
func genJsonData() []byte {
	// 完整数据结构
	raw := &struct {
		Screen
		Battery
		HasTouchID bool //序列化时添加的字段:是否有指纹识别
	}{
		// 屏幕参数
		Screen: Screen{
			Size: 5.5,
			ResX: 1920,
			ResY: 1080,
		},

		// 电池参数
		Battery: Battery{
			4100,
		},

		// 是否有指纹识别
		HasTouchID: true,
	}

	// 将数据序列化为json
	jsonData, _ := json.Marshal(raw)

	return jsonData
}

func main() {

	// 生成一段json数据
	jsonData := genJsonData()

	fmt.Println(string(jsonData))

	// 只需要屏幕和指纹识别信息的结构和实例
	screenAndTouch := struct {
		Screen
		HasTouchID bool
	}{}

	// 反序列化到screenAndTouch
	json.Unmarshal(jsonData, &screenAndTouch)

	// 输出screenAndTouch详细结构
	fmt.Printf("%+v\n", screenAndTouch)

	// 只需要电池和指纹识别信息的结构和实例
	batteryAndTouch := struct {
		Battery
		HasTouchID bool
	}{}

	// 反序列化到batteryAndTouch
	json.Unmarshal(jsonData, &batteryAndTouch)

	// 输出screenAndTouch的详细结构
	fmt.Printf("%+v\n", batteryAndTouch)

}

// 定义数据结构

// 首先,定义手机的各种数据结构体,如屏幕和电池,参考如下代码:
// 定义手机屏幕
type Screen struct {
	Size       float32 //屏幕尺寸
	ResX, RexY int     //屏幕水平和垂直分辨率
}

// 定义电池
type Battery struct {
	Capacity int //容量
}

// 上面代码定义了屏幕结构体和电池结构体，它们分别描述屏幕和电池的各种细节参数。

// 准备 JSON 数据

// 生成JSON数据
func genJsonData() []byte {
	// 完整数据结构
	raw := &struct {
		Screen
		Battery
		HasTouchID bool //序列化时添加的字段:是否有指纹识别
	}{
		// 屏幕参数
		Screen: Screen{
			Size: 5.5,
			ResX: 1920,
			ResY: 1080,
		},

		// 电池参数
		Battery: Battery{
			4100,
		},

		// 是否有指纹识别
		HasTouchID: true,
	}

	// 将数据序列化为JSON
	jsonData, _ := json.Marshal(raw)
}

// 代码说明如下：
// 第 4 行定义了一个匿名结构体。这个结构体内嵌了 Screen 和 Battery 结构体，同时临时加入了 HasTouchID 字段。
// 第 10 行，为刚声明的匿名结构体填充屏幕数据。
// 第 17 行，填充电池数据。
// 第 22 行，填充指纹识别字段。
// 第 26 行，使用 json.Marshal 进行 JSON 序列化，将 raw 变量序列化为 []byte 格式的 JSON 数据。

// 分离JSON数据
// 调用genJsonData获得JSON数据,将需要的字段填充到匿名结构体实例中,通过json.Unmarshal反序列化JSON数据达成分离JSON数据效果.代码如下:

func main() {
	// 生成一段JSON数据
	jsonData := genJsonData()
	fmt.Println(string(jsonData))
	// 只需要屏幕和指纹识别信息的结构和实例
	screenAndTouch := struct {
		Screen
		HasTouchID bool
	}{}
	// 反序列化到screenAndTouch中
	json.Unmarshal(jsonData, &screenAndTouch)
	// 输出screenAndTouch的详细结构
	fmt.Printf("%+v\n", screenAndTouch)
	// 只需要电池和指纹识别信息的结构和实例
	batteryAndTouch := struct {
		Battery
		HasTouchID bool
	}{}
	// 反序列化到batteryAndTouch
	json.Unmarshal(jsonData, &batteryAndTouch)
	// 输出screenAndTouch的详细结构
	fmt.Printf("%+v\n", batteryAndTouch)
}

// 代码说明如下：
// 第 4 行，调用 genJsonData() 函数，获得 []byte 类型的 JSON 数据。
// 第 6 行，将 jsonData 的 []byte 类型的 JSON 数据转换为字符串格式并打印输出。
// 第 9 行，构造匿名结构体，填充 Screen 结构和 HasTouchID 字段，第 12 行中的 {} 表示将结构体实例化。
// 第 15 行，调用 json.Unmarshal，输入完整的 JSON 数据（jsonData），将数据按第 9 行定义的结构体格式序列化到 screenAndTouch 中。
// 第 18 行，打印输出 screenAndTouch 中的详细数据信息。
// 第 21 行，构造匿名结构体，填充 Battery 结构和 HasTouchID 字段。
// 第 27 行，调用 json.Unmarshal，输入完整的 JSON 数据（jsonData），将数据按第 21 行定义的结构体格式序列化到 batteryAndTouch 中。
// 第 30 行，打印输出 batteryAndTouch 的详细数据信息。
