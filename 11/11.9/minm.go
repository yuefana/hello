package main

type Miner interface {
	Len() int
	Swap(i, j int)
	less(i, j int) bool
	ElemIx(ix int) any
}

func Min(m Miner) any {
	min := m.ElemIx(0)
	for i := 1; i < m.Len(); i++ {
		if m.less(i, i-1) {
			min = m.ElemIx(i)
		} else {
			m.Swap(i, i-1)
		}
	}
	return min
}

type IntArray []int

func (arr IntArray) Len() int {
	return len(arr)
}
func (arr IntArray) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// 找最小值
func (arr IntArray) less(i, j int) bool {
	return arr[i] < arr[j]
}
func (arr IntArray) ElemIx(ix int) any {
	return arr[ix]
}
