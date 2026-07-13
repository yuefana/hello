package main

import "strconv"

type Structnew struct {
	//初始化为0 表示栈的下一个元素位置 先赋值后++
	//初始化为-1 表示栈顶元素  先++后赋值
	index int
	//存储数据
	data [LIMIT]int
}

func (s *Structnew) push(data int) bool {
	if s.index >= len(s.data) {
		return false
	}
	s.data[s.index] = data
	s.index++
	return true
}
func (s *Structnew) pop() (int, bool) {
	if s.index == 0 {
		return 0, false
	}
	s.index--
	return s.data[s.index], true
}
func (s Structnew) String() string {
	res := ""
	for i := 0; i < s.index; i++ {
		res += "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
	}
	return res
}
