package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var expr string
	for i := 1; i < len(os.Args); i++ {
		expr = expr + os.Args[i]
	}

	Calculator(expr)
}

func log(exprSlice []string) {
	for _, elem := range exprSlice {
		fmt.Print(elem, " ")
	}
	fmt.Println()
}

func conv(l string, r string) (int, int) {
	a, err := strconv.Atoi(l)
	if err != nil {
		print(err)
	}

	b, err2 := strconv.Atoi(r)
	if err2 != nil {
		print(err)
	}
	return a, b
}

func Calculator(expr string) {
	exprSlice := separator(expr)

	exprSlice = multAndDiv(exprSlice)
	exprSlice = sumAndSub(exprSlice)

	log(exprSlice)
}

func multAndDiv(exprSlice []string) []string {
	for i := 1; i < len(exprSlice); i++ {
		var res, a, b int
		var expSlCopyR []string
		switch exprSlice[i] {
		case "*":
			a, b = conv(exprSlice[i-1], exprSlice[i+1])
			res = a * b
			expSlCopyR = exprSlice[i+1:]
			exprSlice = exprSlice[:i-1]
			exprSlice = append(exprSlice, expSlCopyR[:]...)
			exprSlice[i-1] = strconv.Itoa(res)
			i -= 2
		case "/":
		}
	}
	return exprSlice
}

func sumAndSub(exprSlice []string) []string {
	for i := 1; i < len(exprSlice); i++ {
		var res, a, b int
		switch exprSlice[i] {
		case "+":
			a, b = conv(exprSlice[i-1], exprSlice[i+1])
			res = a + b
			exprSlice = exprSlice[i+1:]
			exprSlice[i-1] = strconv.Itoa(res)
			i -= 2
		case "-":
			a, b = conv(exprSlice[i-1], exprSlice[i+1])
			res = a - b
			exprSlice = exprSlice[i+1:]
			exprSlice[i-1] = strconv.Itoa(res)
			i -= 2
		}
	}
	return exprSlice
}

func separator(expr string) []string {
	exprSlice := make([]string, 0, 15)
	segmtSt := 0

	if expr[0] == ' ' {
		fmt.Println("Unknow element")
		os.Exit(1)
	}
	for i := 1; i < len(expr); i++ {
		if expr[i] == '*' ||
			expr[i] == '/' ||
			expr[i] == '+' ||
			expr[i] == '-' {

			exprSlice = append(exprSlice, expr[segmtSt:i], string(expr[i]))
			segmtSt = i + 1
			if i+2 < len(expr) { //если это ещё не предпоследняя итерация
				i++ //пропустить возможный знак отриц. числа
			}
		}
		if i+2 == len(expr) { // если это последняя итерация
			exprSlice = append(exprSlice, expr[segmtSt:i+2]) // добавить последнее значение
		}
	}
	return exprSlice
}
