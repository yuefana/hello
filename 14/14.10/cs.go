package main

import "fmt"

type Reply struct {
	res int
}

type Request struct {
	a, b   int
	replyc chan *Reply
}

const MAXREQS = 2

var sem = make(chan int, MAXREQS)

func (r *Request) String() string {
	return fmt.Sprintf("%d+%d =%d", r.a, r.b, (<-r.replyc).res)
}

type binOp func(a, b int) int

func run(op binOp, req *Request) {
	req.replyc <- &Reply{op(req.a, req.b)}
}

func server(op binOp, service chan *Request, done chan struct{}) {
	for {
		select {
		case req := <-service:
			sem <- 1
			go run(op, req)
			<-sem
		case <-done:
			return
		}
	}
}

func startService(op binOp) (chan *Request, chan struct{}) {
	reqChan := make(chan *Request)
	done := make(chan struct{})
	go server(op, reqChan, done)
	return reqChan, done
}

func main() {
	addFunc := func(a, b int) int {
		return a + b
	}
	adder, done := startService(addFunc)
	// const N = 100
	// var reqs [N]Request
	// for i := 0; i < N; i++ {
	// 	req := &reqs[i]
	// 	req.a = i
	// 	req.b = i + 100
	// 	req.replyc = make(chan *Reply)
	// 	adder <- req
	// }
	req1 := &Request{3, 4, make(chan *Reply)}
	req2 := &Request{150, 250, make(chan *Reply)}
	adder <- req1
	adder <- req2

	// for i := N - 1; i >= 0; i-- {
	// 	if (<-reqs[i].replyc).res != N+2*i {
	// 		fmt.Println("fail at", i)
	// 	} else {
	// 		fmt.Println("Request", i, "is ok")
	// 	}
	// }
	done <- struct{}{}
	fmt.Println(req1)
	fmt.Println(req2)
}
