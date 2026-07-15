package main

import "fmt"

func main() {
	t := []any{12, "1234"}
	t = mapFunc(mf, t)
	fmt.Println(t)
}

func mf(d any) any {
	switch n := d.(type) {
	case int:
		return n * 2
	case string:
		return n + n
	default:
		return nil
	}
}

func mapFunc(mf func(any) any, list ...any) []any {
	result := make([]any, len(list))
	for ix, v := range list {
		result[ix] = mf(v)
	}
	return result
}
