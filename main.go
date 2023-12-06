// to do!
//  exponents(!)

//  2-3% - проценты
//  2%2 - остаток деления
//  битовый оператор (што ето...)

//  деление 			DONE
//  дроби (5.2) 		DONE
//  удалять все пробелы DONE
//  скобки 				DONE

//  использовать struct, чтобы не использовать цепочку конвертаций string-float-string?

//  Ошибки:
//  передано менее трех аргументов
//  нет операндов в выражении			DONE
//  нет чисел в выражении				DONE
//  некорректная последовательность операндов "---2", "2+2-" "5..2"
//  ошибка на запятую

// прим: кол-во сегментов всегда должно быть 3+ и нечётным

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var expr string
	for i := 1; i < len(os.Args); i++ {
		expr = expr + os.Args[i]
	}
	Calculator(expr)
}

func Calculator(expr string) {

	expr = checksExpr(expr)
	exprSlice := separator(parseParntheses(expr))

	if len(exprSlice) == 0 {
		fmt.Println(expr)
		return
	}

	exprSlice = multAndDiv(exprSlice)
	exprSlice = sumAndSub(exprSlice)

	fmt.Println(exprSlice[0])
}

func checksExpr(expr string) string {
	exprSl := strings.Split(expr, " ")
	var exprUpd string
	for _, elem := range exprSl {
		exprUpd += elem
	}

	if exprUpd[0] == '.' ||
		exprUpd[0] == ',' ||
		exprUpd[0] == '*' ||
		exprUpd[0] == '/' ||
		exprUpd[0] == '+' {
		fmt.Println("Incorrect element on first position", exprUpd[0])
		os.Exit(1)
	}
	return exprUpd
}

func parseParntheses(expr string) string {
	if countOp, countCl := strings.Count(expr, "("), strings.Count(expr, ")"); countOp != countCl {
		fmt.Println("Error func parseParntheses(): Unequal number of parentheses")
		os.Exit(1)
	} else if countOp == 0 {
		return expr
	}
	// также обработать случаи ")(2+2" "(2+)2" "2)+2("

	for strings.Contains(expr, "(") {
		posCl := strings.IndexByte(expr, ')')
		posOp := strings.LastIndexByte(expr[:posCl], '(')

		var exprSlicePart []string
		exprSlicePart = separator(expr[posOp+1 : posCl]) // распарсить выражение между скобок
		exprSlicePart = multAndDiv(exprSlicePart)
		exprSlicePart = sumAndSub(exprSlicePart)
		expr = expr[:posOp] + exprSlicePart[0] + expr[posCl+1:]
	}
	return expr
}

func multAndDiv(exprSlice []string) []string {
	if len(exprSlice) == 1 {
		return exprSlice
	}
	for i := 1; i < len(exprSlice); i++ {
		var expSlCopyR []string
		switch exprSlice[i] {
		case "*":
			a, b := cnvToFloat(exprSlice[i-1], exprSlice[i+1])
			expSlCopyR = exprSlice[i+2:]                           // копируем часть слайса, исключая все пройденные элементы, текущий и следующий
			exprSlice = exprSlice[:i]                              // перезаписываем целый слайс, только неполной левой его частью (исключая текущий и предыдущий)
			exprSlice[i-1] = strconv.FormatFloat(a*b, 'g', -1, 64) // перезаписываем использованное число результатом
			exprSlice = append(exprSlice, expSlCopyR[:]...)        // добавляем правую часть, таким образом, оригинальный срез потерял 2 элемента
			i -= 2                                                 // уменьшаем счетчик итераций на 2
		case "/":
			a, b := cnvToFloat(exprSlice[i-1], exprSlice[i+1])
			expSlCopyR = exprSlice[i+2:]
			exprSlice = exprSlice[:i]
			exprSlice[i-1] = strconv.FormatFloat(a/b, 'g', -1, 64)
			exprSlice = append(exprSlice, expSlCopyR[:]...)
			i -= 2
		}
	}
	return exprSlice
}
func sumAndSub(exprSlice []string) []string {
	if len(exprSlice) == 1 {
		return exprSlice
	}
	for i := 1; i < len(exprSlice); i++ {
		var expSlCopyR []string
		switch exprSlice[i] {
		case "+":
			a, b := cnvToFloat(exprSlice[i-1], exprSlice[i+1])
			expSlCopyR = exprSlice[i+2:]
			exprSlice = exprSlice[:i]
			exprSlice[i-1] = strconv.FormatFloat(a+b, 'g', -1, 64)
			exprSlice = append(exprSlice, expSlCopyR[:]...)
			i -= 2
		case "-":
			a, b := cnvToFloat(exprSlice[i-1], exprSlice[i+1])
			expSlCopyR = exprSlice[i+2:]
			exprSlice = exprSlice[:i]
			exprSlice[i-1] = strconv.FormatFloat(a-b, 'g', -1, 64)
			exprSlice = append(exprSlice, expSlCopyR[:]...)
			i -= 2
		}
	}
	return exprSlice
}
func separator(expr string) []string {
	exprSlice := make([]string, 0, 15)
	segmtSt, countNum, countOper := 0, 0, 0
	for i := 1; i < len(expr); i++ {
		if expr[i] == '*' ||
			expr[i] == '/' ||
			expr[i] == '+' ||
			expr[i] == '-' {
			exprSlice = append(exprSlice, expr[segmtSt:i], string(expr[i]))
			segmtSt = i + 1
			countOper++
			countNum++
			if i+2 < len(expr) { //если это ещё не предпоследняя итерация
				i++ //пропустить возможный знак отриц. числа
			}
		}
		if i+2 == len(expr) { // если это последняя итерация - добавить последнее значение
			exprSlice = append(exprSlice, expr[segmtSt:i+2])
			countNum++
		}
	}
	if countNum == 0 {
		fmt.Println("Incorrect number of numbers received")
		os.Exit(1)
	} else if countOper == 0 {
		fmt.Println("Incorrect number of operators received")
		os.Exit(1)
	}
	return exprSlice
}
func cnvToFloat(a, b string) (float64, float64) {
	l, err := strconv.ParseFloat(a, 64)
	if err != nil {
		fmt.Println("Error func cnvToFloat(), 'a'.", err)
		os.Exit(1)
	}
	r, err2 := strconv.ParseFloat(b, 64)
	if err2 != nil {
		fmt.Println("Error func cnvToFloat(), 'b'.", err2)
		os.Exit(1)
	}
	return l, r
}
