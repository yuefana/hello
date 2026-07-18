package main

import (
	"fmt"
	"math"
	"reflect"
)

func ConvertInt64ToInt(i int64) (res int) {
	value := reflect.ValueOf(res)
	flag := value.OverflowInt(i)
	if flag {
		fmt.Println("err ocured")
		panic("超出32位可表示范围")
	}
	res = int(i)
	return
}

func ConvertInt64ToInt1(i int64) (res int) {
	if math.MinInt32 <= i && i <= math.MaxInt32 {
		res = int(i)
		return
	}
	panic(fmt.Sprintf("%d is out the int32 range", i))
}

func IntFromInt64() {
	defer func() {
		if err, ok := recover().(string); ok {
			fmt.Println(err)
		}
	}()

	num := ConvertInt64ToInt(111111111111)
	fmt.Println(num)
}

func main() {
	IntFromInt64()
}
