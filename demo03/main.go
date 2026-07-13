package main

import (
	"fmt"
	"math"
)

func testSlice() {
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	b1 := b[1:4] //'o', 'l', 'a'
	b2 := b[:2]  //'g', 'o'
	b3 := b[2:]  //'l', 'a', 'n', 'g'
	b4 := b[:]   //'g', 'o', 'l', 'a', 'n', 'g'
	fmt.Printf("b1 is%c\n", b1)
	fmt.Printf("b2 is %c\n", b2)
	fmt.Printf("b3 is %c\n", b3)
	fmt.Printf("b4 is %c\n", b4)
}

func tAppend(slice, data []byte) []byte {
	oldlen := len(slice)
	newlen := oldlen + len(data)
	if newlen > cap(slice) {
		newcap := newlen * 2
		newSlice := make([]byte, oldlen, newcap)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[:newlen]
	copy(slice[oldlen:], data)
	return slice
}
func sum(arrF ...float32) (res float32) {
	for _, v := range arrF {
		res += v
	}
	return
}

func minSlice(arrF []int) int {
	res := math.MaxInt
	for _, v := range arrF {
		if v < res {
			res = v
		}
	}
	return res
}

/*
用顺序函数过滤容器：s 是前 10 个整型的切片。构造一个函数 Filter，第一个参数是 s，第二个参数是一个 fn func(int) bool，返回满足函数 fn 的元素切片。通过 fn 测试方法测试当整型值是偶数时的情况
*/

func Filter(s []int, fn func(int) bool) []int {
	res := make([]int, 0, len(s))
	for _, v := range s {
		if fn(v) {
			res = append(res, v)
		}
	}
	return res
}
func fn(v int) bool {
	return v%2 == 0
}

func main() {
}

func InsertStringSlice(s1, s2 []int, index int) []int {
	res := make([]int, len(s1)+len(s2))
	res = append(res, s1[:index]...)
	res = append(res, s2...)
	res = append(res, s1[index:]...)
	return res
}

func RemoveStringSlice(s1 []int, star, end int) []int {
	s1 = append(s1[:star], s1[end+1:]...)
	return s1
}
