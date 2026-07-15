package sort

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// 切片 值包括 len cap 底层数组
type IntArray []int

func (p IntArray) Len() int {
	return len(p)
}

// Less的含义
// 第 i 个元素是否应该排在第 j 个元素前面
// 第i个元素小就交换
func (p IntArray) Less(i, j int) bool {
	return p[i] < p[j]
}

// 底层数组是共享的
func (p IntArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *IntArray) Put(i int) {
	*p = append(*p, i)
}

func SortInts(a []int) {
	Sort(IntArray(a))
}

func IntsAreSorted(a []int) bool {
	return IsSorted(IntArray(a))
}

type StringArray []string

func (p StringArray) Len() int {
	return len(p)
}

// Less的含义
// 第 i 个元素是否应该排在第 j 个元素前面
// 第i个元素小就交换
func (p StringArray) Less(i, j int) bool {
	return p[i] < p[j]
}

// 底层数组是共享的
func (p StringArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func SortStrings(a []string) {
	Sort(StringArray(a))
}

func IsSorted(data Sorter) bool {
	for i := data.Len() - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}

// 冒泡排序
func Sort(data Sorter) {
	for i := 0; i < data.Len(); i++ {
		for j := data.Len() - 1; j > 0; j-- {
			if data.Less(j, j-1) {
				data.Swap(j-1, j)
			}
		}
	}
}
