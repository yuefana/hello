package main

import (
	"fmt"
)

type T struct {
	value int
}

func (_ T) ValueMethod() {
	fmt.Printf("ValueMethod")
}
func (_ *T) PointerMethod() {
	fmt.Println("PointerMethod")
}
func (t T) String() string {
	return fmt.Sprint(t.value)
}

type Valuer interface {
	ValueMethod()
}

type Pointerer interface {
	PointerMethod()
}

type Qp struct {
	*T
}

type Qv struct {
	T
}

func m() {
	//T
	p1 := T{0}
	if _, ok := any(p1).(Valuer); ok {
		fmt.Println("T实现了Valuer接口", p1)
	}
	p1 = T{1}
	if _, ok := any(p1).(Pointerer); ok {
		fmt.Println("T实现了Pointerer接口", p1)
	} else {
		fmt.Println("T没有实现Pointerer接口")
		// p1 可寻址，编译器自动转换为 (&p1).PointerMethod()
		p1.PointerMethod()
	}
	//*T
	p2 := &T{2}
	if _, ok := any(p2).(Valuer); ok {
		fmt.Println("*T实现了Valuer接口", p2)
	}

	p2 = &T{3}
	if _, ok := any(p2).(Pointerer); ok {
		fmt.Println("*T实现了Pointerer接口", p2)
	}
	//Qp
	p3 := Qp{&T{4}}
	if _, ok := any(p3).(Valuer); ok {
		fmt.Println("Qp实现了Valuer接口", p3)
	}

	p3 = Qp{&T{5}}
	if _, ok := any(p3).(Pointerer); ok {
		fmt.Println("Qp实现了Pointerer接口", p3)
	}
	//*Qp
	p4 := &Qp{&T{5}}
	if _, ok := any(p4).(Valuer); ok {
		fmt.Println("*Qp实现了Valuer接口", p4)
	}
	p4 = &Qp{&T{6}}
	if _, ok := any(p4).(Pointerer); ok {
		fmt.Println("*Qp实现了Pointerer接口", p4)
	}

	//Qv
	p5 := Qv{T{7}}
	if _, ok := any(p5).(Valuer); ok {
		fmt.Println("Qv实现了Valuer接口", p5)
	}
	p5 = Qv{T{8}}
	if _, ok := any(p5).(Pointerer); ok {
		fmt.Println("Qv实现了Pointerer接口", p5)
	} else {
		fmt.Println("Qv未实现了Pointerer接口")
	}
	//*Qv
	p6 := &Qv{T{8}}
	if _, ok := any(p6).(Valuer); ok {
		fmt.Println("*Qv实现了Valuer接口", p6)
	}
	p6 = &Qv{T{9}}
	if _, ok := any(p6).(Pointerer); ok {
		fmt.Println("*Qv实现了Pointerer接口", p6)
	}

}

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}
type car struct {
	make  string
	model string
	price float32
}
type valuable interface {
	getValue() float32
}

func (s stockPosition) getValue() float32 {
	return s.count * s.sharePrice
}

func (c car) getValue() float32 {
	return c.price
}

func showValue(v valuable) {
	fmt.Printf("%f\n", v.getValue())
}
func mn() {
	s := stockPosition{"", 1, 2}
	c := car{price: 12}
	showValue(s)
	showValue(c)
}

type Simpler interface {
	Get() int
	Set(num int)
}
type Simple struct {
	n int
}
type RSimple struct {
	n int
}

func (s RSimple) Get() int {
	return s.n
}
func (s *RSimple) Set(num int) {
	s.n = num
}

func (s Simple) Get() int {
	return s.n
}
func (s *Simple) Set(num int) {
	s.n = num
}

func fi(s Simpler) int {
	switch v := s.(type) {
	case *Simple:
		fmt.Println("type is *Simple")
		return v.n
	case *RSimple:
		fmt.Println("type is *RSimple")
		return v.n
	default:
		fmt.Println("other")
		return 0
	}
}
func gI(a any) int {
	if v, ok := a.(Simpler); ok {
		return v.Get()
	}
	return 0
}
func main() {
	m()

}
