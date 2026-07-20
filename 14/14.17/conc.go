package main

import "fmt"

type Person struct {
	Name   string
	salary float64
	chF    chan func()
}

func NewPerson(name string, salary float64) *Person {
	p := &Person{name, salary, make(chan func())}
	go p.backend()
	return p
}

// 后台协程
func (p *Person) backend() {
	for f := range p.chF {
		f()
	}
}

func (p *Person) SetSalary(sal float64) {
	p.chF <- func() { p.salary = sal }
}
func (p *Person) Salary() float64 {
	fchan := make(chan float64)
	p.chF <- func() {
		fchan <- p.salary
	}
	return <-fchan
}
func (p *Person) String() string {
	fchan := make(chan string)
	p.chF <- func() {
		fchan <- fmt.Sprintf("NAME %s SALARY %f", p.Name, p.salary)
	}
	return <-fchan
}

func main() {
	bs := NewPerson("mike", 100)
	bs.SetSalary(190)
}
