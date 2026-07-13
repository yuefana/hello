package main

import (
	"fmt"
)

// 常量和变量的声明
const b = 10

var c = 20

type T struct {
	// 结构体的字段
	name string
	age  int
}

/*
GO的基本数据类型

	  1int类型 int8 int16 int32 int64  uint8 uint16 uint32 uint64 uintptr
		2float类型 float32 float64
		3bool类型 true false
		4string类型 字符串
		5byte类型 uint8的别名
		6rune类型 int32的别名 表示一个Unicode码点

		& | ^ 翻转某一位 x^=1<<n
		&^  把右边是 1 的位，在左边全部清零
		<< >>
*/
func main() {
	var a int8 = 20
	fmt.Printf("num=%v num的类型是%T\n", a, a)
}
