package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	//-n 未出现时，eachLine 指向的值为 false
	//-n 出现时，eachLine 指向的值为 true。
	eachLine := flag.Bool("n", false, "print each argument on a new  line")

	flag.Parse()
	args := flag.Args()

	separator := " "
	if *eachLine {
		separator = "\n"
	}
	fmt.Println(strings.Join(args, separator))
}
