package main

import (
	"fmt"
	"os"
	"strconv"
)

type Day int

var dayName = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func (day Day) String() string {
	return dayName[day]
}

type Celsius float64

func (c Celsius) String() string {
	//return strconv.FormatFloat(float64(c),'f',1,64)+"°C"
	return fmt.Sprintf("%06.1f °C", float64(c))
}

type Stringer interface {
	String() string
}

func print(args ...any) {
	for i, arg := range args {
		if i > 0 {
			os.Stdout.WriteString(" ")
			switch a := arg.(type) {
			case Stringer:
				os.Stdout.WriteString(a.String())
			case int:
				os.Stdout.WriteString(strconv.Itoa(a))
			case string:
				os.Stdout.WriteString(a)
			default:
				os.Stdout.WriteString("???")
			}
		}
	}
}
