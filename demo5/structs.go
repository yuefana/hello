package main

import "fmt"

// 结构体
type struct1 struct {
	i1  int
	f1  float32
	str string
}

// 单向链表
type Node struct {
	data int
	next *Node
}

// 双向链表 or 树节点
type TNode struct {
	data int
	//lchild
	pre *TNode
	//rchild
	next *TNode
}

// 时间间隔
type Interval struct {
	start int
	end   int
}

func modify(s struct1) {
	s.i1 = 10
}
func modify1(s *struct1) {
	s.i1 = 10
}

func main() {
	//指向结构体的指针
	//ms := new(struct1)
	ms := &struct1{i1: 20, f1: 15.5, str: "hello"}
	modify1(ms)

	//指向这个结构体，传入函数的是结构体的副本
	//var ms struct1
	//modify(ms)
	fmt.Printf("The int is: %d\n", ms.i1)
	fmt.Printf("The float is: %f\n", ms.f1)
	fmt.Printf("The string is: %s\n", ms.str)
	fmt.Println(ms)

	intr := &Interval{0, 3}
	fmt.Println(intr)
}
