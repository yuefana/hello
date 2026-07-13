// package main

// import "fmt"

// type Base struct {
// 	id int
// }

// type Employee struct {
// 	Persion
// 	salary float64
// }

// type Persion struct {
// 	Base
// 	FirstName string
// 	LastName  string
// }

// func (b Base) Id() int {
// 	return b.id
// }

// func (b *Base) SetId(id int) {
// 	b.id = id
// }

//	func main() {
//		e := Employee{Persion{Base{10086}, "yue", "fan"}, 100.0}
//		fmt.Println(e.id)
//	}
package main

import (
	"fmt"
)

type Base struct{}

func (Base) Magic() {
	fmt.Println("base magic")
}

func (self Base) MoreMagic() {
	self.Magic()
	self.Magic()
}

type Voodoo struct {
	Base
}

func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}

func main() {
	v := new(Voodoo)
	v.Magic()
	v.MoreMagic()
}
