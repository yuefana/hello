package main

import "fmt"

var resume chan int

func integers() chan int {
	yield := make(chan int)
	count := 0
	go func() {
		for {
			yield <- count
			count++
		}
	}()
	return yield
}

func generateInteger() int {
	return <-resume
}

func main() {
	eval()
}

// 1   2  3  4      5    6  7
// 14 15 16  17 18  19  20 21

type EvalFunc func(any) (any, any)

func BuildLazyEvalutor(evalFunc EvalFunc, initState any) func() any {
	retValChan := make(chan any)
	loopFunc := func() {
		var currentState any = initState
		var retVal any
		for {
			retVal, currentState = evalFunc(currentState)
			retValChan <- retVal
		}
	}
	retFunc := func() any {
		return <-retValChan
	}
	go loopFunc()
	return retFunc
}
func BuildLazyIntEvaluator(evalFunc EvalFunc, initState any) func() int {
	ef := BuildLazyEvalutor(evalFunc, initState)
	return func() int {
		return ef().(int)
	}
}
func eval() {
	evenFunc := func(state any) (any, any) {
		os := state.([]int)
		p, q := os[0], os[1]
		ns := []int{q, p + q}
		return p, ns
	}
	even := BuildLazyIntEvaluator(evenFunc, []int{1, 1})
	for i := 0; i < 10; i++ {
		fmt.Printf("%vth even: %v\n", i+1, even())
	}
}
