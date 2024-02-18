package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fileIn, err := os.Open("data/in2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fileIn.Close()

	fileOut, err := os.Create("data/data_out.txt")
	writer := bufio.NewWriter(fileOut)
	if err != nil {
		log.Fatal(err)
	}
	defer fileOut.Close()

	defer func() {
		if panicErr := recover(); panicErr != nil {

			fileOut, err := os.Open("data/data_out.txt")
			if err != nil {
				log.Fatal(err)
			}
			defer fileOut.Close()

			dataOut := make([]byte, 1000)
			fileOut.Read(dataOut)

			fmt.Println("here:", string(dataOut))
		}
	}()

	scan := bufio.NewScanner(fileIn)
	stringNum := 0
	for scan.Scan() {
		stringNum++
		strArr := strings.Split(scan.Text(), "|")
		for _, v := range strArr {
			if v == "" {
				log.Panic(fmt.Sprintf("parse error: empty field on string %d", stringNum))
			}
		}

		_, err := writer.WriteString(fmt.Sprintf("Row: %d\nName: %s\nAddress: %s\nCity: %s\n\n\n", stringNum, strArr[0], strArr[1], strArr[2]))
		if err != nil {
			log.Fatal(err)
		}
		writer.Flush()
	}
}
