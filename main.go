package main

import "fmt"

func main() {
	Spruce()
}

func Spruce() { //ёлочка
	height := 0
	fmt.Scan(&height)
	indent := height
	weight := 1
	fmt.Printf("\u001B[32m")
	for i := height; i > 0; i-- {
		indent--
		for j := indent; j > 0; j-- { //вывести пробелы, в кол-ве равном indent
			fmt.Printf(" ")
		}
		for j := weight; j > 0; j-- { //вывести диезы, в кол-ве равном weight
			fmt.Printf("#")
		}
		fmt.Println()
		weight += 2
	}
	fmt.Printf("\u001B[0m")
	fmt.Printf("\x1b[30m")
	heightTr := height / 3
	indentTr := height - 1
	for i := heightTr; i > 0; i-- {
		for j := indentTr; j > 0; j-- { //вывести пробелы, в кол-ве равном indentTr и диез в конце
			fmt.Printf(" ")
		}
		fmt.Println("#")
	}
	fmt.Printf("\u001B[0m")
}
