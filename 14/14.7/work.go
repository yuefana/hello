package main

import "sync"

type task struct {
}

// 领取任务，处理任务，提交处理结果
func worker(in, out chan *task) {
	// for {
	// 	t := <-in  //领取任务
	// 	process(t) //处理任务
	// 	out <- t   //提交结果
	// }
	for t := range in {
		process(t)
		out <- t
	}
}

func process(t *task) {
	//panic("unimplemented")
}

func main() {
	pending, done := make(chan *task), make(chan *task)
	go sendwork(pending) //分发任务
	var wg sync.WaitGroup
	N := 10
	for i := 0; i < N; i++ { //开启N个grounting去处理
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(pending, done)
		}()
	}
	go func() {
		wg.Wait()
		close(done)
	}()

	consumeWork(done) //接受并消费处理结果
}

func consumeWork(done chan *task) {
	//panic("unimplemented")
}

func sendwork(pending chan *task) {
	defer close(pending)
	task := task{}
	//for _, task := range tasks {
	pending <- &task
	//	}
}

//通道通常用于所有权交接、任务分发和异步结果；互斥锁通常用于缓存和共享状态

type Service struct {
	chcheMu sync.RWMutex //保护共享缓存
	cache   map[string]int
	jobs    chan task      //任务分发
	wg      sync.WaitGroup //等待worker结束
}
