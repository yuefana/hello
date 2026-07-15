package main

import (
	"fmt"
	"strconv"

	sort "example.com/hello/11/11.7"
)

type day struct {
	num       int
	shortName string
	longName  string
}
type dayArray struct {
	data []*day
}

func (p dayArray) Len() int {
	return len(p.data)
}

// Less的含义
// 第 i 个元素是否应该排在第 j 个元素前面
// 第i个元素小就交换
func (p dayArray) Less(i, j int) bool {
	return p.data[i].num < p.data[j].num
}

func (p dayArray) Swap(i, j int) {
	p.data[i], p.data[j] = p.data[j], p.data[i]
}

func (p dayArray) String() string {
	str := ""
	for _, v := range p.data {
		str += strconv.Itoa(v.num) + v.shortName
	}
	return str
}

func sai() {
	/*
		i 是 data 的切片描述符副本 len cap 底层数组
		刚创建时两者共享底层数组；执行 i.Put(999) 后，append 分配了新数组，才导致两者分离。
	*/
	data := []int{1, 8, 3, 4}
	//排序只修改底层数组元素，不改变切片长度、容量和指针
	//sort.SortInts(&data)
	fmt.Println(data)
	//IntArray
	i := sort.IntArray(data)
	//i.Put(999)
	//想让追加后的结果也赋给 data
	data = []int(i)
	sort.Sort(i)
	fmt.Println("i", i)
	fmt.Println("data", data)
	fmt.Println(sort.IsSorted(i))

	data1 := []string{"monday", "friday", "tuesday", "wednesday", "sunday", "thursday", "", "saturday"}
	j := sort.StringArray(data1)
	sort.Sort(j)
	fmt.Println("data1", data1)
}
func main() {
	//sai()
	//Days()
	Persions()
}

func Days() {
	Sunday := day{0, "SUN", "Sunday"}
	Monday := day{1, "MON", "Monday"}
	Tuesday := day{2, "TUE", "Tuesday"}
	Wednesday := day{3, "WED", "Wednesday"}
	Thursday := day{4, "THU", "Thursday"}
	Friday := day{5, "FRI", "Friday"}
	Saturday := day{6, "SAT", "Saturday"}
	data := []*day{&Tuesday, &Thursday, &Wednesday, &Sunday, &Monday, &Friday, &Saturday}
	a := dayArray{data}
	fmt.Println(a)
	sort.Sort(a)
	for _, d := range data {
		fmt.Println(d.shortName)
	}
	fmt.Println(a)
	//panic("fail")
}

type Persion struct {
	firstName string
	lastName  string
}

type Persons []Persion

func (p Persons) String() string {
	str := ""
	for _, v := range p {
		str += v.firstName
	}
	return str
}
func (p Persons) Len() int {
	return len(p)
}

// Less的含义
// 第 i 个元素是否应该排在第 j 个元素前面
// 第i个元素小就交换
func (p Persons) Less(i, j int) bool {
	return p[i].firstName < p[j].firstName
}

func (p Persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func Persions() {
	p1 := Persion{"b", "b"}
	p2 := Persion{"a", "a"}
	data := Persons{p1, p2}
	fmt.Println(data)
	sort.Sort(data)
	fmt.Println(data)
}
