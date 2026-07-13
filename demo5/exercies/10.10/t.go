package main

// type Integer int

// func (p Integer) get() int {
//    return int(p)
// }

func f(i int) {

}

type Integer struct {
	n int
}

func (p Integer) get() int {
	return p.n
}

func main() {
	var v Integer
	//  f(int(v))
	f(v.n)
}
