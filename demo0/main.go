package main

import "fmt"

func getUserInfo() (string, int) {
	return "张三", 18
}

func main() {
	/*
		  变量声明的两种方式  同一个作用域不支持重复声明
			  1. var声明  var 变量名 变量类型 = 值 没有初始化值为空
				    var  a,b int = 10,20
						var  (
							a = 10
							b = 20
						)
			  2. 短变量声明  变量名 := 值
					 a,b,c:=12,13,"C"
	*/
	a, b, c := 12, 13, "C"
	//	var b int = 10
	//	a := 10
	fmt.Println("Hello, World!", a, b, c)
	fmt.Printf("a=%v a的类型是%T", a, a)
	//匿名变量	_  用于接收不需要的值
	var name, _ = getUserInfo()
	fmt.Printf("name=%v", name)
	//常量声明 const 变量名 变量类型 = 值 同时声明多个常量时，如果省略了值则表示和上面一行的值相同
	const PI = 3.14
	fmt.Printf("PI=%v\n", PI)
	const (
		n1 = iota
		_
		n2 = 100
		n3 = iota
		n4
	)
	fmt.Printf("n1=%v n2=%v n3=%v n4=%v\n", n1, n2, n3, n4)
	const (
		m1, m2 = iota, iota + 1
		m3, m4
	)
	fmt.Printf("m1=%v m2=%v m3=%v m4=%v\n", m1, m2, m3, m4)
}
