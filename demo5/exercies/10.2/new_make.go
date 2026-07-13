package main

import "fmt"

type Foo map[string]string

type Bar struct {
	thingOne string
	thingTwo string
}

func main() {
	y := new(Bar)
	//y.thingOne = "hello"
	(*y).thingOne = "1"
	//y.thingTwo = "world"
	(*y).thingTwo = "2"
	fmt.Println(y)

	//make 只能用于slice map chan
	//需要创建切片描述符和底层数组
	//z := make(Bar)
	//fmt.Println(z)

	x := make(Foo)
	x["x"] = "hi"
	fmt.Println(x)

	u := new(Foo)    //u指向nil 尚未被分配内存
	(*u) = make(Foo) //初始化
	(*u)["x"] = "i"
	fmt.Println(u)
}
