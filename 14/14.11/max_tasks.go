package main

const MAXREQS = 50

var sem = make(chan int, MAXREQS)

type Request struct {
	a, b    int
	replayc chan int
}

func process(r *Request) {
	//do something
}

func handler(r *Request) {
	sem <- 1
	process(r)
	<-sem
}

func server(services chan *Request) {
	for {
		req := <-services
		go handler(req)
	}
}
func main() {
	service := make(chan *Request)
	go server(service)
}
