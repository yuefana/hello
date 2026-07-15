package main

import "fmt"

type Car struct {
	Model        string
	Manufacturer string
	BuildYear    int
}
type Cars []*Car

// 处理切片的每一个元素，具体的处理看传入的函数
func (c Cars) Process(f func(car *Car)) {
	for i := range c {
		f(c[i])
	}
}

// 传入的函数决定是否选取这个car
func (c Cars) FindAll(f func(car *Car) bool) Cars {
	cars := make([]*Car, 0)
	c.Process(func(c *Car) {
		if f(c) {
			cars = append(cars, c)
		}
	})
	return cars
}

func (c Cars) Map(f func(c *Car) any) []any {
	result := make([]any, len(c))
	ix := 0
	c.Process(func(c *Car) {
		result[ix] = f(c)
		ix++
	})
	return result
}
func MakeSortedAppender(manufacturers []string) (func(car *Car), map[string]Cars) {
	sortedCars := make(map[string]Cars)
	for _, m := range manufacturers {
		sortedCars[m] = make([]*Car, 0)
	}
	sortedCars["Default"] = make([]*Car, 0)

	appender := func(c *Car) {
		if _, ok := sortedCars[c.Manufacturer]; ok {
			sortedCars[c.Manufacturer] = append(sortedCars[c.Manufacturer], c)
		} else {
			sortedCars["Default"] = append(sortedCars["Default"], c)
		}
	}
	return appender, sortedCars
}
func c() {
	// make some cars:
	ford := &Car{"Fiesta", "Ford", 2008}
	bmw := &Car{"XL 450", "BMW", 2011}
	merc := &Car{"D600", "Mercedes", 2009}
	bmw2 := &Car{"X 800", "BMW", 2008}
	// query:
	allCars := Cars([]*Car{ford, bmw, merc, bmw2})
	allNewBMWs := allCars.FindAll(func(car *Car) bool {
		return (car.Manufacturer == "BMW") && (car.BuildYear > 2010)
	})
	fmt.Println("AllCars: ", allCars)
	fmt.Println("New BMWs: ", allNewBMWs)
	//
	manufacturers := []string{"Ford", "Aston Martin", "Land Rover", "BMW", "Jaguar"}
	sortedAppender, sortedCars := MakeSortedAppender(manufacturers)
	allCars.Process(sortedAppender)
	fmt.Println("Map sortedCars: ", sortedCars)
	BMWCount := len(sortedCars["BMW"])
	fmt.Println("We have ", BMWCount, " BMWs")
}

// 多个匿名函数如果捕获的是同一个外部变量，那么它们共享这个变量。外部变量被重新赋值后，之后所有通过该变量访问数据的闭包，都会看到变量的新值。
func makeMap() (func(), func(), map[string]int, *int) {
	m := map[string]int{"count": 0}
	n := 0
	add := func() {
		m["count"]++
		n++
	}

	reset := func() {
		m["count"] = 2
		n = 3
		n++
		m = make(map[string]int)
		m["count"] = 100

	}

	return add, reset, m, &n
}

func t() {
	add, reset, returnedMap, n := makeMap()

	add()
	fmt.Println(returnedMap, *n) // map[count:1]

	reset()
	add()

	fmt.Println(returnedMap, *n) // map[count:2]
}
func main() {
	t()
}
