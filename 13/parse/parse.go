// 输入的字符串解析为整数切片；这个包有自己特殊的 ParseError
// 当没有东西需要转换或者转换成整数失败时，这个包会 panic()
// 可导出的 Parse() 函数会从 panic() 中 recover() 并用所有这些信息返回一个错误给调用者
package parse

import (
	"fmt"
	"strconv"
	"strings"
)

type ParseError struct {
	Index int
	Word  string
	Err   error
}

func (e *ParseError) String() string {
	return fmt.Sprintf("pkg parse :error parsing %q as int", e.Word)
}

func (e *ParseError) Error() string {
	return fmt.Sprintf(
		"pkg parse: error parsing %q as int",
		e.Word,
	)
}

func Parse(input string) (number []int, err error) {

	defer func() {
		if r := recover(); r != nil {
			var ok bool
			if err, ok = r.(error); !ok {
				err = fmt.Errorf("pkg:%v [%T]", r, r)
			}
		}
	}()
	fields := strings.Fields(input)
	number = fields2numbers(fields)
	return
}
func fields2numbers(fields []string) (numbers []int) {
	if len(fields) == 0 {
		panic("no words to parse")
	}
	for idx, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			panic(&ParseError{idx, field, err})
		}
		numbers = append(numbers, num)
	}
	return
}
