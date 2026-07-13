package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"strings"
)

func test01() {
	for i := 0; i < 15; i++ {
		fmt.Printf("%d", i)
	}
	fmt.Println()
	i := 0
START:
	fmt.Printf("%d", i)
	i++
	if i < 15 {
		goto START
	}
}

func test02() {
	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("G")
		}
		fmt.Println()
	}
	s := "G"
	for i := 0; i < 5; i++ {
		println(s)
		s += "G"
	}
}

func test03() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%b", i)
		fmt.Println()
	}
}
func test04() {
	for i := 0; i < 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

func test05() {
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
}

func compute(x1 int, x2 int) (sum int, acc int, diff int) {
	sum = x1 + x2
	acc = x1 * x2
	diff = x1 - x2
	return
}

func compute2(x1 int, x2 int) (int, int, int) {
	return x1 + x2, x1 * x2, x1 - x2
}

func MySqrt(x float64) (float64, error) {
	if x < 0 {
		return math.NaN(), fmt.Errorf("cannot Sqrt negative number: %v", x)
	}
	return math.Sqrt(x), nil
}

func MySqrt2(x float64) (res float64, err error) {
	if x < 0 {
		res = math.NaN()
		err = fmt.Errorf("cannot Sqrt negative number: %v", x)
		return
	}
	res = math.Sqrt(x)
	return
}

func Multiply(x1, x2 int, res *int) {
	*res = x1 * x2
}
func test06() {
	m := 0
	Multiply(3, 4, &m)
	fmt.Println(m)
}
func test07(s ...string) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func test08() {
	s := make([]string, 5, 10)
	test07(s...)
}

func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("defer func1(%q)=%d.%v", s, n, err)
	}()
	return 7, io.EOF
}

func trace(s string) string {
	fmt.Printf("entering func %q\n", s)
	return s
}
func un(s string) {
	fmt.Printf("leaving func %q\n", s)
}
func test09() {
	defer un(trace("test08"))
	fmt.Println("in test08")
}

func even(x int) bool {
	if x == 0 {
		return true
	}
	return even(x - 1)
}
func odd(x int) bool {
	if x == 0 {
		return false
	}
	return even(x - 1)
}

func revsign(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func trace1(n int) {
	if n == 0 {
		return
	}
	fmt.Printf("%d\n", n)
	trace1(n - 1)

}

/*
包 strings 中的 Map() 函数和 strings.IndexFunc() 一样都是非常好的使用例子。请学习它的源代码并基于该函数书写一个程序，要求将指定文本内的所有非 ASCII 字符替换成问号 '?' 或空格 ' '。您需要怎么做才能删除这些字符呢？
*/
func test10() {
	s := "hello中国world"
	for {
		index := strings.IndexFunc(s, isNonASCII)
		if index == -1 {
			break
		}
		r := []rune(s)
		for i, ch := range r {
			if ch > 127 {
				r[i] = ' '
				//r = append(r[:i], r[i+1:]...)
				break
			}
		}
		s = string(r)
	}
	fmt.Println(s)
	s = strings.Map(func(r rune) rune {
		if r == ' ' || r == '?' {
			return -1
		}
		return r
	}, s)
	fmt.Println(s)
}

func isNonASCII(r rune) bool {
	if r > 127 {
		return true
	}
	return false
}

func test11() {
	//fmt.Printf("%d", fibonacci()())
	var ar [3]int
	f(ar)   // passes a copy of ar
	fp(&ar) // passes a pointer to ar
}
func fibonacci() func() int {
	p, q := 0, 1
	return func() int {
		p, q = q, p+q
		return p
	}
}
func f(a [3]int)   { fmt.Println(a) }
func fp(a *[3]int) { fmt.Println(a) }
