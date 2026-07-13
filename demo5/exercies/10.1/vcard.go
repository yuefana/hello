package main

import (
	"fmt"
	"math"
	"time"
	"unsafe"
)

type Address struct {
	Street      string
	HouseNumber uint32
	ZipCode     string
	City        string
	Country     string
}

type VCard struct {
	Name      string
	Brithday  time.Time
	Image     string
	Addresses map[string]*Address
}

type Point struct {
	x float64
	y float64
}

type Polar struct {
	x float64
	y float64
	z float64
}

func (p Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}
func (p *Polar) Scale() {
	p.x, p.y, p.z = p.x*2, p.y*3, p.z*4
}

func main() {
	fmt.Println(unsafe.Sizeof(Polar{}))
	fmt.Println(unsafe.Sizeof(int(0)))
	fmt.Println(unsafe.Sizeof(float32(0.0)))
}
