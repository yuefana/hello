package main

import (
	"fmt"
	"reflect"
)

/*
	   reflect
		 反射包的Type表示一个go类型
		 Field获取结构体的第i个字段
*/
type Persion struct {
	name string
	age  int
}

func tryReflect() {
	x := 5.3
	y := []int{1, 3, 2}
	z := Persion{"hi", 1}
	p := Persion{}
	q := []Persion{z, p}
	r := &z
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.TypeOf(x).Kind())
	fmt.Println(reflect.ValueOf(x))
	fmt.Println(reflect.ValueOf(x).Kind())
	fmt.Println(reflect.ValueOf(x).Interface())

	fmt.Println(reflect.TypeOf(y))
	fmt.Println(reflect.TypeOf(y).Kind())
	fmt.Println(reflect.ValueOf(y))

	fmt.Println(reflect.TypeOf(z))
	fmt.Println(reflect.TypeOf(z).Kind())
	fmt.Println(reflect.TypeOf(z).Field(0))
	// {name main string  0 [0] false}
	// 字段名 未导出字段所属的包 字段类型 没有结构体标签 Offset 字段索引路径 不是匿名嵌入字段
	fmt.Println(reflect.ValueOf(z))
	fmt.Println(reflect.ValueOf(z).Field(0))

	fmt.Println(reflect.TypeOf(q))
	//Type.Elem()表示类型
	fmt.Println(reflect.TypeOf(q).Elem())
	fmt.Println(reflect.ValueOf(q))
	//Value.Elem()表示解引用  r是一个指针
	fmt.Println(reflect.ValueOf(r).Elem())
	//Value.Index(i)  切片或数组的第 i 个元素值
	//Value.Field(i)  结构体的第 i 个字段值
	fmt.Println(reflect.ValueOf(q).Index(0).Field(0))
}
func reflect1() {
	var x float64 = 3.4
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	fmt.Println("type", t)
	fmt.Println("value", v)
	fmt.Println("type", v.Type())
	fmt.Println("kind", v.Kind())
	fmt.Println("value", v.Float())
	fmt.Println(v.Interface())
	fmt.Printf("value is %5.2e\n", v.Interface())
	y := v.Interface().(float64)
	fmt.Println(y)
}

func change() {
	var x float64 = 3.4
	v := reflect.ValueOf(&x)
	v = v.Elem()
	fmt.Println(v.CanSet())
	v.SetFloat(5.0)
	fmt.Println(x)
	fmt.Println(v)

}

type NotknownType struct {
	S1, s2, s3 string
}

func (n NotknownType) String() string {
	return n.S1 + " - " + n.s2 + " - " + n.s3
}

var secret NotknownType = NotknownType{"A", "B", "C"}

func main() {
	t := reflect.TypeOf(secret)
	v := reflect.ValueOf(secret)
	vn := reflect.ValueOf(&secret)
	fmt.Println(t, v.Type(), v.Kind())
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Field  i %d is %v\n", i, v.Field(i))
	}
	fmt.Println(secret)
	vn = vn.Elem()
	for i := 0; i < vn.NumField(); i++ {
		fmt.Printf("Field  vn i %d is %v\n", i, vn.Field(i))
	}
	vn.Field(0).SetString("D")
	fmt.Println(secret)
	res := v.Method(0).Call(nil)
	res = vn.Method(0).Call(nil)
	fmt.Println(res)
}
