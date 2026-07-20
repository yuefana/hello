package main

// func ParalleProcessData(in <-chan *int, out chan<- *int) {
// 	preOut := make(chan *int, 100)
// 	stepAOut := make(chan *int, 100)
// 	stepBOut := make(chan *int, 100)
// 	stepCOut := make(chan *int, 100)
// 	go PreprocessData(in, preOut)
// 	go ProcessStepA(preOut, stepAOut)
// 	go ProcessStepB(stepAOut, stepBOut)
// 	go ProcessStep3(stepBOut, stepCOut)
// 	go PostProcessData(stepCOut, out)
// }

type Buffer struct {
}

var freeList = make(chan *Buffer, 100)

var serverChan = make(chan *Buffer)

func client() {
	for {
		var b *Buffer
		select {
		case b = <-freeList:
		default:
			b = new(Buffer)
		}
		serverChan <- b
	}
}

func server() {
	for {
		b := <-serverChan
		//process(b)
		select {
		case freeList <- b:
		default:
		}
	}
}

func main() {

}
