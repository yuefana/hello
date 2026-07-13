package main

import "fmt"

type Engine interface {
	Start()
	Stop()
}

type GasEngine struct {
}

func (_ GasEngine) Start() {
	fmt.Println("汽油引擎启动")
}

func (_ GasEngine) Stop() {
	fmt.Println("汽油引擎停止")
}

type Car struct {
	Engine
	wheelCount int
}

func (c *Car) Start() {
	fmt.Println("car start run")
}
func (c *Car) Stop() {
	fmt.Println("car stop run")
}

type Mercedes struct {
	Car
}

func (this Mercedes) sayHiToMerkel() {
	fmt.Println("HiToMerkel", this.wheelCount)
}

func (this Car) numberOfWheels() int {
	return this.wheelCount
}

func (c *Car) GoToWorkIn() {
	c.Engine.Start()
	// get in car
	c.Start()
	// drive to work
	fmt.Printf("轮子的数量 %d\n", c.wheelCount)
	c.Stop()
	// get out of car
	c.Engine.Stop()
}

func main() {
	m := Mercedes{Car{Engine: GasEngine{}, wheelCount: 1}}
	m.GoToWorkIn()
	m.sayHiToMerkel()
}
