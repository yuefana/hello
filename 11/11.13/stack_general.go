package main

import (
	"errors"
	"fmt"
)

type Stack[T any] struct {
	data []T
}

func (s Stack[T]) Len() int {
	return len(s.data)
}
func (s Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}
func (s *Stack[T]) Push(x T) {
	s.data = append(s.data, x)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("栈为空")
	}
	lastIndex := len(s.data) - 1
	v := s.data[lastIndex]

	s.data = s.data[:lastIndex]
	return v, nil
}
func main() {
	var s Stack[int]
	s.Push(1)
	s.Push(2)
	s.Push(3)
	//栈非空才会弹出
	for !s.IsEmpty() {
		//err==nil 为正常
		if v, err := s.Pop(); err == nil {
			fmt.Printf("%v\t", v)
		}
	}
	unsafe.
}
