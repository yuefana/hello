package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func cat(r *bufio.Reader, flag *bool, linenumber *int) {
	for {
		buf, err := r.ReadBytes('\n')
		if len(buf) > 0 {
			if *flag {
				fmt.Fprintf(os.Stdout, "%4d", *linenumber)
				*linenumber += 1
			}
			//fmt.Fprintf(os.Stdout, "%s", buf)
			os.Stdout.Write(buf)
		}
		if err == io.EOF {
			fmt.Fprintf(os.Stdout, "\n")
			break
		}
		if err != nil {
			fmt.Println("读取失败")
			break
		}
	}
}

func main() {
	b := flag.Bool("n", false, "每一行头部加入一个行号")
	flag.Parse()
	lineNumber := 1
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin), b, &lineNumber)
	}
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s error reading from %s %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat(bufio.NewReader(f), b, &lineNumber)
		f.Close()
	}
}
