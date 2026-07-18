package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func searchPart(values []int, target int, offset int, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i, v := range values {
		if v == target {
			result <- offset + i
		}
	}
}

func main() {
	values := make([]int, 1_000_000)
	for i := range values {
		values[i] = i
	}

	target := 876543
	workers := 4
	//向上取整
	partSize := (len(values) + workers - 1) / workers
	result := make(chan int, workers)
	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		//第一个要包含的元素的下标
		start := w * partSize
		//结束边界，指向最后一个元素的后一个位置
		end := start + partSize
		if start >= len(values) {
			break
		}
		if end > len(values) {
			end = len(values)
		}
		wg.Add(1)
		go searchPart(values[start:end], target, start, result, &wg)
	}

	go func() {
		wg.Wait()
		close(result)
	}()
	foundIndex := -1
	for index := range result {
		foundIndex = index
		break
	}
	fmt.Println("找到位置", foundIndex)
}

func t() {
	fmt.Println("逻辑CPU数", runtime.NumCPU())
	//GOMAXPROCS 限制同一时刻能够执行用户级 Go 代码的操作系统线程数量，也就是 Go 代码可使用的最大并行度。
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
	fmt.Println("IN MAIN()")
	var wg sync.WaitGroup
	wg.Add(2)
	go longWait(&wg)
	go shortWait(&wg)
	fmt.Println("Waiting in main()")
	wg.Wait()
	fmt.Println("At the end of main()")
}

func longWait(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Beginning longWait()")
	time.Sleep(5 * time.Second)
	fmt.Println("END OF LongWait()")
}

func shortWait(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Beginning shortWait()")
	time.Sleep(2 * time.Second)
	fmt.Println("END OF shortWait()")
}
