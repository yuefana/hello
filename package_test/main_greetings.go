package main

import (
	"fmt"

	"example.com/hello/package_test/even"
	"example.com/hello/package_test/greetings"
)

func main() {
	greetings.Say()
	fmt.Printf("是否为AM,%t", greetings.IsAm())
	for i := range 100 {
		fmt.Printf("%d 是否为偶数,%t", i, even.IsEven(i))
	}
}
