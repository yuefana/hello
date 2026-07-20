package main

import (
	"fmt"
)

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

// 生产者
func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
}

func getData(ch chan string) {
	var input string
	for {
		input = <-ch
		fmt.Printf("%s", input)
		if input == "Tokyo" {
			break
		}
	}
}

type semaphore chan struct{}

func (s semaphore) P(n int) {
	for i := 0; i < n; i++ {
		s <- struct{}{}
	}
}

func (s semaphore) V(n int) {
	for range n {
		<-s
	}
}

/* mutexes */
func (s semaphore) Lock() {
	s.P(1)
}

func (s semaphore) Unlock() {
	s.V(1)
}

/* signal-wait */
func (s semaphore) Wait(n int) {
	s.P(n)
}

func (s semaphore) Signal() {
	s.V(1)
}
func gosum() {
	ch := make(chan int)
	go func(x1 int, x2 int) {
		ch <- x1 + x2
	}(1, 2)
	res := <-ch
	fmt.Println(res)
}

func producer_consumer() {
	//var wg sync.WaitGroup
	ch := make(chan int)
	done := make(chan struct{})
	//wg.Add(2)
	go func(ch chan int) {
		//	defer wg.Done()
		defer close(ch)
		for i := range 10 {
			ch <- i * 10
		}
	}(ch)
	go func(ch chan int) {
		//	defer wg.Done()
		for range 10 {
			fmt.Println(<-ch)
		}
		done <- struct{}{}
	}(ch)
	//wg.Wait()
	<-done
}

func pump1() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck1[T any](ch chan T) {
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
}

// func (c *container) Iter() <-chan T {
// 	ch := make(chan T)
// 	go func() {
// 		defer close(ch)
// 		for i := range 10 {
// 			ch <- c.items[i]
// 		}
// 	}()
// 	return ch
// }

func generate() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func filter(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			i := <-in
			if i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func sieve() chan int {
	out := make(chan int)
	ch := generate()
	go func() {
		for {
			prime := <-ch
			out <- prime
			ch = filter(ch, prime)
		}
	}()
	return out
}

func main() {
	primes := sieve()
	for {
		fmt.Println(<-primes)
	}
}
