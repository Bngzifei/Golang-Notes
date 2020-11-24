package main

import "fmt"

//type Direction int
//
//const (
//	North Direction = iota   // 0
//	East  // 1
//	South // 2
//	West
//)
//
//func (d Direction) String() string {
//	return [...]string{"North", "East", "South", "West"}[d]
//}
//
//func main() {
//	fmt.Println(South)
//}

type Math struct {
	x, y int
}

var m = map[string]Math{
	"foo": &Math{2, 3},
}

func main() {
	//tmp := m["foo"]
	//tmp.x = 4
	//m["foo"] = tmp
	//fmt.Println(m["foo"].x)
	m["foo"].x = 4
	fmt.Println(m["foo"].x)
	fmt.Printf("%#v", m["foo"])
}
