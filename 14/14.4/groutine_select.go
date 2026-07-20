package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

func suck(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Println("Received on channel 1:", v)
		case v := <-ch2:
			fmt.Println("Received on channel 2:", v)
		}
	}
}
func for_loop() {
	ch := make(chan int)
	done := make(chan struct{})
	go tel(ch, done)
	// for {
	// 	v, ok := <-ch
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println(v)
	// }
	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-done:
			fmt.Println("输入结束")
			return
		}
	}
}
func tel(ch chan int, done chan struct{}) {
	defer close(ch)
	for i := 0; i < 10; i++ {
		ch <- i
	}

	done <- struct{}{}
}
func fibo() {
	ch := make(chan int)
	go comput1(ch, 10)
	for {
		if v, ok := <-ch; ok {
			fmt.Println(v)
		} else {
			return
		}
	}
}
func comput1(ch chan int, num int) {
	defer close(ch)
	for i := 1; i <= num; i++ {
		ch <- l1(i)
	}
}
func l1(num int) int {
	p, q := 1, 1
	for i := 1; i < num; i++ {
		p, q = q, p+q
	}
	return p
}

func comput2(ch chan int, num int) {
	defer close(ch)
	p, q := 1, 1

	for i := 0; i < num; i++ {
		ch <- p
		p, q = q, p+q
	}
}

func fibo2() {
	ch := make(chan int)
	go comput2(ch, 10)
	for v := range ch {
		fmt.Println(v)
	}
}

func comput3(ch chan int, done chan struct{}, num int) {
	defer close(ch)
	p, q := 1, 1
	for i := 0; i < num; i++ {
		ch <- p
		p, q = q, p+q
	}
	done <- struct{}{}
}
func fibo3() {
	ch := make(chan int)
	done := make(chan struct{})
	go comput3(ch, done, 10)
	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-done:
			return
		}
	}
}

func dup3(in <-chan int) (<-chan int, <-chan int, <-chan int) {
	a, b, c := make(chan int, 2), make(chan int, 2), make(chan int, 2)
	go func() {
		for {
			x := <-in
			a <- x
			b <- x
			c <- x
		}
	}()

	return a, b, c
}
func fib() <-chan int {
	x := make(chan int, 2)
	a, b, out := dup3(x)
	go func() {
		x <- 0
		x <- 1
		<-a
		for {
			x <- <-a + <-b
		}
	}()
	<-out
	return out
}
func fibo4() {
	x := fib()
	for range 10 {
		fmt.Println(<-x)
	}
}

func randon_bitgen() {
	ch := make(chan int)
	go func() {
		for {
			fmt.Print(<-ch, " ")
		}
	}()
	for {
		select {
		case ch <- 0:
		case ch <- 1:
		}
	}
}
func polar_to_cartesian() {
	polarCh := make(chan polar)
	cartesianch := make(chan pointer)
	var wg sync.WaitGroup
	wg.Add(3)
	go readPolar(polarCh, &wg)
	go convertToCartesian(polarCh, cartesianch, &wg)
	go printCartesian(cartesianch, &wg)
	wg.Wait()
}

type pointer struct {
	x, y float64
}
type polar struct {
	radius float64
	angle  float64
}

func readPolar(polarCh chan<- polar, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(polarCh)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("请输入极坐标下的半径和角度:(输入q退出)")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取输入失败", err)
			return
		}
		line = strings.TrimSpace(line)
		if line == "q" {
			fmt.Println("程序退出")
			return
		}
		fields := strings.Fields(line)
		if len(fields) != 2 {
			fmt.Println("请输入两个数")
			continue
		}
		radius, err := strconv.ParseFloat(fields[0], 64)
		angle, err := strconv.ParseFloat(fields[1], 64)
		if radius < 0 {
			fmt.Println("半径不能小于0")
			continue
		}
		polarCh <- polar{radius, angle}

	}

}
func convertToCartesian(polarch <-chan polar, cartesianch chan<- pointer, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(cartesianch)
	for v := range polarch {
		radian := v.angle * math.Pi / 180
		x := v.radius * math.Cos(radian)
		y := v.radius * math.Sin(radian)
		cartesianch <- pointer{x, y}
	}
}

func printCartesian(cartesianch <-chan pointer, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range cartesianch {
		x, y := v.x, v.y
		fmt.Println(x, " ", y)
	}
}

func concurrent(ch chan float64, done chan struct{}) {
	defer close(ch)
	i := 0.0
	sign := 1.0
	for {
		select {
		case <-done:
			return
		case ch <- 4 * sign / (2*i + 1):
			i++
			sign = -sign
		}
	}

}
func main() {
	ch := make(chan float64)
	done := make(chan struct{})
	res := 0.0
	stoped := false
	go concurrent(ch, done)
	for v := range ch {
		if stoped {
			continue
		}
		res += v
		if math.Abs(res-math.Pi) < 0.01 {
			//done <- struct{}{}
			close(done)
			//防止重复close() select 不一定立刻选择done那个channel
			stoped = true
		}
	}
	fmt.Print(res)
}
