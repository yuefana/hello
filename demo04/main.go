package main

import "fmt"

func removestr(s string, index int) (string, string) {
	sb := []byte(s)
	s1, s2 := sb[:index], sb[index:]
	return string(s1), string(s2)
}

func string_reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i <= j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
func main() {
	data := []int{3, 2, 1}
	res := mapFunc(fc, data)
	fmt.Printf("%d", res)
}
func uniq(date []byte) []byte {
	if len(date) == 0 {
		return nil
	}
	pre := date[0]
	res := []byte(nil)
	for _, v := range date {
		if v != pre {
			res = append(res, v)
		}
		pre = v
	}
	return res
}

func bubblesort(date []int) {
	for i := 0; i < len(date); i++ {
		for j := len(date) - 1; j > i; j-- {
			if date[j-1] > date[j] {
				date[j-1], date[j] = date[j], date[j-1]
			}
		}
	}
}
func mapFunc(fc func(int) int, e []int) (res []int) {
	for _, v := range e {
		res = append(res, fc(v))
	}
	return
}
func fc(num int) int {
	return 10 * num
}
