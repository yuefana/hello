package main

import (
	"fmt"
	"reflect"

	"example.com/hello/demo5/exercies/structPack"
)

func main() {
	// struct1 := new(structPack.ExpStruct)
	// struct1.Mi1 = 10
	// struct1.Mi2 = 16.0
	// fmt.Println(struct1)
	// r1 := structPack.NewNode(10)
	// var re any
	// re = 10
	// switch value := re.(type) {
	// case int:
	// 	fmt.Println("int", value)
	// case string:
	// 	fmt.Println("str", value)
	// }
	// //类型断言
	// if value, ok := re.(int); ok {
	// 	fmt.Println(value)
	// }

	// fmt.Println(r1)
	test()
}

func test() {
	t := structPack.NewNode(100)
	for i := range 3 {
		//返回*Node类型 ptr
		tType := reflect.TypeOf(t)
		fmt.Println(tType.Kind())
		//Node类型 struct
		tType = tType.Elem()
		fmt.Println(tType.Kind())
		fileds := tType.Field(i)
		fmt.Println(fileds.Name)
		fmt.Println(fileds.Tag.Get("info"))
	}
}
