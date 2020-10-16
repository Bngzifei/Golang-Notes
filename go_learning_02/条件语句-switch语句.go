package main

/*
注意：Go 没有三目运算符，所以不支持 ?: 形式的条件判断。

switch 语句用于基于不同条件执行不同动作,每一个case分支都是
唯一的,从上至下逐一测试,直到匹配为止.

switch语句执行的过程从上至下,直到找到匹配项,匹配项后面也不
需要再加break.

switch默认情况下case最后自带break语句,匹配成功后就不会执行
其他case,如果我们需要执行后面的case,可以使用fallthrough.

*/

import "fmt"

func main() {
	/* 定义局部变量 */
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
	fmt.Printf("你的等级是 %s\n", grade)
}
