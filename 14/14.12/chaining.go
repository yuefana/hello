package main

import (
	"flag"
	"fmt"
	"runtime"
)

var ngoroutine = flag.Int("n", 3, "how many goroutines")

// 把right里面的数据取出来加1放到left里面
func f(left, right chan int) {
	left <- 1 + <-right
}
func chain() {
	flag.Parse()
	leftmost := make(chan int)
	var left, right chan int = nil, leftmost
	for i := 0; i < *ngoroutine; i++ {
		left, right = right, make(chan int)
		go f(left, right)
	}
	right <- 0
	x := <-leftmost
	fmt.Println(x)
}
func chain1() {
	flag.Parse()
	r := make(chan int)
	var left, right chan int = r, nil
	for i := 0; i < *ngoroutine; i++ {
		left, right = make(chan int), left
		go f(left, right)
	}
	r <- 0
	x := <-left
	fmt.Println(x)
}
func main() {
	chain1()
}

func DoAll() {
	NCPU := runtime.GOMAXPROCS(0)
	sem := make(chan int, NCPU)
	for i := 0; i < NCPU; i++ {
		go DoPart(sem)
	}
	for i := 0; i < NCPU; i++ {
		<-sem
	}
	//所有任务都已经做完
}

func DoPart(sem chan int) {
	sem <- 1 //表示这部分已经做完
}
