package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// defer logTime()
	getFile()

}

func getFile() {
	f, _ := os.Open("in2.txt")
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
}

func logTime() {

}
