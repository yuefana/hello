package main

import (
	"fmt"
	"sync"
	"time"
)

type mT struct {
	mu sync.Mutex
	time.Time
}

type nT time.Time

// 可以显示转换
func (this nT) nfirst3Chars() string {
	return time.Time(this).String()[0:3]
}

func (this *mT) first3Chars() string {
	return this.String()[0:3]
	//return this.Time.String()[0:3]
}

func main() {
	m := mT{sync.Mutex{}, time.Now()}
	fmt.Println(m.first3Chars())
}
