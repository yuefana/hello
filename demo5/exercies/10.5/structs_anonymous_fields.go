package main

import "fmt"

type innerS struct {
	in1 int
	in2 int
}
type outerS struct {
	b   int
	c   float32
	int // anonymous field
	innerS
}

func main() {
	outer := new(outerS)
	outer.b = 1
	outer.c = 7.5
	outer.int = 60
	outer.innerS.in1 = 1
	outer.innerS.in2 = 2
	//因为是匿名字段
	outer.in1 = 3
	outer.in2 = 4
	fmt.Println(outer)
	outer1 := outerS{7, 1.5, 6, innerS{5, 10}}
	fmt.Println(outer1)
}
