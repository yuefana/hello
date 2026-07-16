package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"example.com/hello/11/Stack"
)

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input1                 = "56.12/ 52 12/ Go"
	format                 = "%f/%d/%s"
)

func readinput1() {
	//fmt.Println("Please enter your full name:")
	//Scanln() 扫描来自标准输入的文本，将空格分隔的值依次存放到后续的参数内，直到碰到换行
	//fmt.Scanln(&firstName, &lastName)
	//fmt.Scanf("%s %s", &firstName, &lastName)

	//fmt.Printf("Hi %s %s!\n", firstName, lastName)
	//Sscanf从字符串读取
	fmt.Sscanf(input1, format, &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s)
}

// bufio
func readinput2() {
	//*bufio.Reader
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入:")
	if input, err := inputReader.ReadString('\n'); err == nil {
		fmt.Printf("The input is %s \n", input)
	}
}

func switch_input() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入你的名字")
	if input, err := inputReader.ReadString('\n'); err == nil {
		switch input {
		case "Philip\r\n", "Ivo\r\n":
			fmt.Printf("Welcome %s\n", input)
		case "Chris\r\n":
			fmt.Println("Welcome Chris!")
			//fallthrough
		default:
			fmt.Printf("You are not welcome here! Goodbye!")
		}
	}
}

func word_letter_count() {
	inputReader := bufio.NewReader(os.Stdin)
	line, wordnum, charnum := 1, 0, 0
	preisspace := true
	if input, err := inputReader.ReadString('S'); err == nil {
		//按照rune遍历
		//for i,v:=range input

		//按照字节遍历
		for i := 0; i < len(input); i++ {
			if input[i] == 'S' {
				break
			}

			switch input[i] {
			case '\r':
				preisspace = true
			case '\n':
				line++
				preisspace = true
			case ' ', '\t':
				charnum++
				preisspace = true
			default:
				if preisspace {
					wordnum++
				}
				charnum++
				preisspace = false
			}
		}
	}

	fmt.Printf("字符个数%d 单词个数 %d 行数 %d", charnum, wordnum, line)
}

func add(x1 float64, x2 float64) float64 {
	return x1 + x2
}
func diff(x1 float64, x2 float64) float64 {
	return x1 - x2
}
func mul(x1 float64, x2 float64) float64 {
	return x1 * x2
}
func div(x1 float64, x2 float64) float64 {

	return x1 / x2
}

func calculator() {
	inputReader := bufio.NewReader(os.Stdin)
	var inputarr []string
	if input, err := inputReader.ReadString('q'); err == nil {
		inputarr = strings.Split(input, "\r\n")
	}
	var dataStack Stack.Stack[float64]
	m := map[string]func(float64, float64) float64{
		"+": add, "-": diff, "*": mul, "/": div,
	}

	for _, v := range inputarr {
		if i, err := strconv.ParseFloat(v, 64); err == nil {
			dataStack.Push(i)
		} else if v == "q" {
			re, _ := dataStack.Pop()
			fmt.Println("最终结果为", re)
		} else {
			op := v
			x2, _ := dataStack.Pop()
			x1, _ := dataStack.Pop()
			dataStack.Push(m[op](x1, x2))
		}
	}
}

func calculator1() {
	inputReader := bufio.NewReader(os.Stdin)
	var dataStack Stack.Stack[float64]
	operations := map[string]func(float64, float64) float64{
		"+": add, "-": diff, "*": mul, "/": div,
	}

	fmt.Println("请输入：number1、number2、operator")
	fmt.Println("每输入一项后按 Enter，输入 q 结束")

	for {
		fmt.Print("> ")
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("读取错误：", err)
			return
		}
		// 去掉 Windows 的 \r\n 或其他首尾空白。
		v := strings.TrimSpace(input)
		if v == "" {
			continue
		}

		if v == "q" {
			fmt.Println("程序结束")
			return
		}

		// 题目要求输入整型数，所以使用 Atoi。
		number, err := strconv.Atoi(v)
		if err == nil {
			dataStack.Push(float64(number))
			continue
		}

		// 不是整数，判断是否是合法运算符。
		operation, _ := operations[v]
		x2, err := dataStack.Pop()
		if v == "/" && x2 == 0 {
			fmt.Println("除数不能为 0")
			dataStack.Push(x2)
			continue
		}
		x1, err := dataStack.Pop()
		result := operation(x1, x2)
		fmt.Println("结果为：", result)
		dataStack = Stack.Stack[float64]{}
	}
}
func main() {
	readinput1()
}
