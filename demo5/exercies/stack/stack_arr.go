package main

import "strconv"

type Stack [LIMIT]int

func (s *Stack) push(num int) bool {
	for i := range s {
		if s[i] == 0 {
			s[i] = num
			return true
		}
	}
	//添加失败
	return false
}

func (s *Stack) pop() int {
	v := 0
	for i := len(s) - 1; i >= 0; i-- {
		if v = s[i]; v != 0 {
			s[i] = 0
			return v
		}
	}
	//没有可以弹出来的了
	return -1
}

func (s Stack) String() string {
	res := ""
	for i, v := range s {
		if s[i] != 0 {
			res += "[" + strconv.Itoa(i) + ":" + strconv.Itoa(v) + "]"
		}
	}
	return res
}
