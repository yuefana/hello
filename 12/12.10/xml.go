package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

type Person struct {
	FirstName string
	LastName  string
}

func xmlte() {
	p := Person{}
	input := "<Person><FirstName>Laura</FirstName><LastName>Lynn</LastName></Person>"
	inputreader := strings.NewReader(input)
	decoder := xml.NewDecoder(inputreader)
	decoder.Decode(&p)
	fmt.Println(p)
}

func token() {
	input := "<Person id='11' name='yue'><FirstName>Laura</FirstName><LastName>Lynn</LastName></Person>"
	inputreader := strings.NewReader(input)
	decoder := xml.NewDecoder(inputreader)
	for {
		t, err := decoder.Token()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("解析失败：", err)
			return
		}

		switch token := t.(type) {
		case xml.StartElement:
			fmt.Println("开始标签：", token.Name.Local)
			for _, attr := range token.Attr {
				attrName := attr.Name.Local
				attrValue := attr.Value
				fmt.Printf("An attribute is: %s %s\n", attrName, attrValue)
			}

		case xml.EndElement:
			fmt.Println("结束标签：", token.Name.Local)

		case xml.CharData:
			content := strings.TrimSpace(string(token))
			if content != "" {
				fmt.Println("文本内容：", content)
			}
		}
	}
}
func main() {
	token()
}
