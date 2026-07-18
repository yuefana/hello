package main

import (
	"fmt"

	"example.com/hello/13/parse"
)

func main() {

	e := []string{"1 2 3 4 5", "100 6.5,12.5", "2 + 2 = 4", "1st class", ""}
	for _, v := range e {
		fmt.Printf("Parsing %q:\n", v)
		nums, err := parse.Parse(v)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(nums)
	}
}
