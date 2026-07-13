package main

import (
	"fmt"
	"strconv"
)

type Tint struct {
	a, b int
}

type T struct {
	a int
	b float32
	c string
}

type Celsius float64

type Day int

var week = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

const (
	MON Day = iota
	TUE
	WED
	THU
	FRI
	SAT
	SUN
)

type TZ int

const (
	UTC TZ = iota
)

var mapper = map[TZ]string{UTC: "Universal Greenwich time"}

func (t TZ) String() string {
	if re, ok := mapper[t]; ok {
		return re
	}
	return ""
}

func (d Day) String() string {
	return week[int(d)]
}

func (c Celsius) String() string {
	return strconv.FormatFloat(float64(c), 'f', -1, 64) + "°C"
}

func (t *T) String() string {
	//return strconv.Itoa(t.a) + "/" + strconv.FormatFloat(float64(t.b), 'f', 6, 32) + "/" + `"` + t.c + `"`
	return fmt.Sprintf("%d/%.6f/%q", t.a, t.b, t.c)
}

func (t *Tint) String() string {
	// return "(" + strconv.Itoa(t.a) + "/" + strconv.Itoa(t.b) + ")"
	return fmt.Sprintf("(%d/%d)", t.a, t.b)
}

func main() {
	t := Tint{1, 2}
	t1 := new(Tint)
	fmt.Println(t1)
	fmt.Println(t)

	tt := &T{7, -2.35, "abc\tdef"}
	fmt.Println(tt)
}
