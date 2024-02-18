package main

import (
	"bufio"
	"fmt"
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

	file, err := os.Open("data/in.txt")
	if err != nil {
		fmt.Println("os.Open('data/in.txt')", err)
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

	stringNum, countByte := 0, 0
	for scan.Scan() {
		stringNum++
		a, err := writer.WriteString(fmt.Sprintf("%v. ", stringNum) + scan.Text() + "\n")
		countByte += a

		if err != nil {
			fmt.Println("writer.WriteString", err)
			os.Exit(1)
		}
	}
	writer.Flush()

	if err != nil {
		fmt.Println("writer.WriteString", err)
		os.Exit(1)
	}
	fmt.Printf("writed: %v strings, %v bytes\n", stringNum, countByte)
	fileOut.Close()

}
