package main

import (
	"fmt"
	"sync"
)

type Info struct {
	//互斥锁
	mu sync.Mutex
}

func Update(info *Info) {
	info.mu.Lock()
	fmt.Printf("修改内容")
	info.mu.Unlock()
}

func main() {

}
