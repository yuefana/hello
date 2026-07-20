package main

import (
	"fmt"
	"time"
)

func main() {
	timer_grountine()
}

// func runService(r <-chan Request, client *Client) {
// 	const (
// 		ratePerSec = 10
// 		burst      = 5
// 	)
// 	interval := time.Second / time.Duration(ratePerSec)

// 	tokens := make(chan struct{}, burst)

// 	for i := 0; i < burst; i++ {
// 		tokens <- struct{}{}
// 	}
// 	ticker := time.NewTicker(interval)
// 	defer ticker.Stop()

// 	go func() {
// 		for range ticker.C {
// 			select {
// 			case tokens <- struct{}{}:
// 				//令牌桶未满
// 			default:
// 				//令牌桶已满
// 			}
// 		}
// 	}()

// 	for req := range r {
// 		<-tokens //取得令牌，没有令牌时阻塞

// 		go client.Call("service", req)
// 	}

// }

func timer_grountine() {
	tick := time.Tick(time.Second)
	boom := time.After(5 * time.Second)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM")
			return
		default:
			fmt.Println("   .")
			time.Sleep(5 * time.Second)
		}
	}
}
