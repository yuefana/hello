package greetings

import (
	"fmt"
	"time"
)

func Say() {
	switch {
	case IsAm():
		fmt.Println("Good Day")
	case IsAfternoon():
		fmt.Println("Good Day")
	case IsEvening():
		fmt.Println("Good Night")
	}
}

func IsAm() bool {
	hour := time.Now().Hour()
	return hour < 12
}

func IsAfternoon() bool {
	hour := time.Now().Hour()
	return hour >= 12 && hour < 18
}
func IsEvening() bool {
	hour := time.Now().Hour()
	return hour >= 18
}
