package main

import (
	"fmt"
	"runtime"
	"time"
)

const LIMIT = 4

type Resource struct {
	name string
}

func NewResource(name string) *Resource {
	r := &Resource{name: name}

	runtime.SetFinalizer(r, func(obj *Resource) {
		fmt.Println("执行终结函数", obj.name)
	})
	return r
}

func main() {
	// var m runtime.MemStats
	// runtime.ReadMemStats(&m)
	// fmt.Printf("%d Kb\n", m.Alloc/1024)
	// runtime.SetFinalizer()
	r := NewResource("DB CONNECTION")
	fmt.Println("正在使用", r.name)
	runtime.KeepAlive(r)
	r = nil
	runtime.GC()
	time.Sleep(time.Second)
}
