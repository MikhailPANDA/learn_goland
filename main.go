package main

import "fmt"

func main() {
	//Calc()
	f()

}

func f() {

loop:
	for {
		fmt.Println("only once1")
		break loop
	}

	// boolVarCanChange := true
	// for boolVarCanChange {
	// 	// boolVarCanChange = false
	// 	fmt.Println("only once2")
	// 	break loop
	// }
}

/*
func sw() {
	a := 0
	switch a {
	case 2:
		fmt.Println("a=2")
	case 1:
		fmt.Println("a=1")
		fallthrough
	default:
		fmt.Println("def")
	}

	b := 1
	switch {
	case b == 2:
		fmt.Println("b=2")
	case b == 1:
		fmt.Println("b=1")
		break
	default:
		fmt.Println("def")
	}
}

/*
func fn() {
	m := map[string]string{"first": "1", "second": "2"}

	if _, booleanVar := m["first"]; booleanVar { // booleanVar == false
		fmt.Println("вывод", m["11"], m["first"], booleanVar)
	}
}

func Maps() {
	m := map[string]map[string]string{
		"2020": {
			"10":            "1",
			"продано копий": "11",
		},
		"2021": {
			"0":             "1",
			"продано копий": "1",
		},
		"2022": {
			"Вася":          "1",
			"продано копий": "10",
		},
		"2023": {
			"Восилий":       "1",
			"продано копий": "100",
		},
		"2024": {
			"Восилёк":       "1",
			"продано копий": "11",
		},
	}

	// _ = m
	// fmt.Println(int(m["2024"]["продано копий"])) //не работает, пушто не работает int() со строками

	fmt.Println(strconv.Atoi(m["2024"]["продано копий"])) //работает

	// fmt.Printf("%d\n", m["2020"]["продано копий"]) //не работает, пушто не работает %d со строками
	// fmt.Sprintf("%d\n", m["2020"]["продано копий"]) // -...-

	for _, v := range m {
		s, err := strconv.Atoi(v["продано копий"])
		if err != nil {
			fmt.Println(err)
			return
		}
		if s > 0 {
			fmt.Println("больше нуля")
		}
	}
}



/*
func Week() {
	// week1 := make([]int, 5, 7)
	// week1 = []int{1, 2, 3, 4, 5}
	// week1 = append(week1[0:1], week1[2:]...)

	// week := [6]bool{
	// 	4: true,
	// 	0: true,
	// }

	// weekFREE := week[:5]
	// var weekFREE1 = []string{}
	// weekFREE1 := make([]string, 5, 7)
	// weekFREE3 := copy(weekFREE1, week)

	week := []string{
		"mon",
		"tue",
		"wed",
		"thu",
		"fri",
		"sun",
		"sat",
	}

	weekWORK := week[0:5]

	week = week[5:7]

	// weekFREE := []string{}
	// for i, elem := range week {

	// 	if elem == "sun" || elem == "sat" {
	// 		weekFREE[i] = ``
	// 	}
	// }

	fmt.Println(weekWORK, week)

	week = append(weekWORK, week...)

	fmt.Println(weekWORK, week)

}

/*
func Print(str) {
	fmt.Println(str)
}

/*
func Calc() {
	var str string
	fmt.Scan(&str)

	oper := "1"
	switch oper {
	case "+":
		return str(0)+str(2)
	case "-":
		return str(0)-str(2)

	}
}
*/
// var A *int
// 	B := 123
// 	A = &B
// 	fmt.Println(*A)
/*
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
*/
/*
func test() {

	greeting := "привет"
	var convGreeting = []rune(greeting)

	convGreeting[4] = 'Е'
	fmt.Println(string(convGreeting))
	greeting[4] = 'Л'
	// fmt.Println(greeting)
	_ = greeting
}

func kratnoe7() {
	var a, b int
	fmt.Scan(&a, &b)

	for i := b; i >= a; i-- {
		if i%7 == 0 {
			fmt.Println(i)
			return
		}

	}
	fmt.Println("NO")
}
*/
