package sort

import (
	"math/rand"
	"strconv"
)

type Float64Array []float64

func (f Float64Array) Len() int {
	return len(f)
}

func (f Float64Array) Less(i, j int) bool {
	return f[i] < f[j]
}
func (f Float64Array) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func NewFloat64Array() Float64Array {
	return make(Float64Array, 25)
}
func (f Float64Array) List() string {
	s := "{"
	for _, v := range f {
		if v == 0 {
			continue
		}
		s += strconv.FormatFloat(v, 'f', 2, 64) + " "
	}
	s += "}"
	return s
}
func (f Float64Array) String() string {
	return f.List()
}
func (f Float64Array) Fill(n int) {
	if n < 0 {
		n = 0
	}
	if n > f.Len() {
		n = f.Len()
	}
	for i := 0; i < n; i++ {
		f[i] = rand.Float64() * 100
	}
}
