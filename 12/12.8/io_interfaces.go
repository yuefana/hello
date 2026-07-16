package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, _ := os.Open("goprogram")
	outputFile, _ := os.OpenFile("goprogramT", os.O_WRONLY|os.O_CREATE, 0666)
	defer inputFile.Close()
	defer outputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)
	for {
		inputString, _, readerError := inputReader.ReadLine()
		if len(inputString) > 0 {
			outputString := string(inputString[2:5]) + "\r\n"
			_, err := outputWriter.WriteString(outputString)
			if err != nil {
				fmt.Println("写入失败", err.Error())
				return
			}
		}
		if readerError == io.EOF {
			outputWriter.Flush()
			fmt.Println("EOF")
			return
		}
		if readerError != nil {
			fmt.Println("读取失败")
			return
		}
	}
	//fmt.Println("Conversion done")
}
