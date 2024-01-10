package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func logTime() func() {
	start := time.Now()
	return func() {
		end := time.Now().Sub(start)
		fmt.Println("log:", end)
	}
}

func main() {
	logger := logTime()
	defer logger()

	file, err := os.Open("in.txt")
	if err != nil {
		fmt.Println("os.Open('in.txt')", err)
		os.Exit(1)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)

	fileOut, err := os.Create("out.txt")
	writer := bufio.NewWriter(fileOut)
	if err != nil {
		fmt.Println("os.Create('out.txt')", err)
		os.Exit(1)
	}

	stringNum := 0
	for scan.Scan() {
		stringNum++
		_, err := writer.WriteString(fmt.Sprintf("%v. ", stringNum) + scan.Text() + "\n")

		if err != nil {
			fmt.Println("writer.WriteString", err)
			os.Exit(1)
		}
	}

	writer.Flush()

	fileOutReader, err := os.Open("out.txt")
	countByte, err := io.ReadAll(fileOutReader)
	defer fileOutReader.Close()

	if err != nil {
		fmt.Println("writer.WriteString", err)
		os.Exit(1)
	}
	fmt.Printf("writed: %v strings, %v bytes\n", stringNum, len(countByte))
	fileOut.Close()

}
