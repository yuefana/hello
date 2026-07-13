package main

import (
	"fmt"
	"io"
	"strings"
	"unicode/utf8"
)

func main() {
	test11()
}

func test() {
	s := "asSASA ddd dsjkdsjs dk"
	fmt.Printf("字符串长度：%d\n", len(s))
	fmt.Printf("字符串长度：%d\n", utf8.RuneCountInString(s))
	reader := strings.NewReader(s)
	//res, _ := reader.ReadByte()
	buf := make([]byte, 5)
	for {
		n, err := reader.Read(buf)
		fmt.Printf("读取的字节数：%d\n", n)
		//读取有效字节
		fmt.Printf("读取的内容：%s\n", string(buf[:n]))
		if err == io.EOF {
			break
		}
	}
	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
	}
}
