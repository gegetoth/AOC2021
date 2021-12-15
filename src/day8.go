package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func getSlice(inStr string) []string {
	d := make([]string, 0)
	for _, l := range []rune(inStr) {
		d = append(d, string(l))
	}
	return d
}

func getSlices(inStr string) [][]string {
	in := make([][]string, 0)
	for _, digit := range strings.Split(inStr, " ") {
		in = append(in, getSlice(digit))
	}

	return in
}

func getInputOutPutArrays(line string) ([]string, []string) {
	inStr := strings.TrimSpace(strings.Split(line, "|")[0])
	outStr := strings.TrimSpace(strings.Split(line, "|")[1])

	return strings.Split(inStr, " "), strings.Split(outStr, " ")
}

type Number struct {
	upper     string
	bottom    string
	middle    string
	topRight  string
	topLeft   string
	loweRight string
	lowerLeft string
}

func run8_1() {
	var input = read_lines("C:\\Users\\tothg\\Gege\\AOC2021\\res\\day_8_example.txt")
	sum := 0
	for _, line := range input {
		_, out := getInputOutPutArrays(line)

		for _, digit := range out {
			fmt.Printf("%+v\n", digit)
			if len(digit) == 2 || len(digit) == 3 || len(digit) == 4 || len(digit) == 7 {
				sum++
			}
		}

	}
	println(sum)
}

func difference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

func isSubset(a, b []string) bool {
	for _, i := range b {
		if !contains(a, i) {
			return false
		}
	}
	return true
}

//func delete(lst []string, elem string) []string {
//	e := -1
//	for i := 0; i < len(lst); i++ {
//		if lst[i] == elem {
//			e = i
//		}
//	}
//	lst[e] = lst[len(lst)-1]
//	return lst[:len(lst)-1]
//}

func deleteP(lst *([]string), elem string) {
	e := -1
	for i := 0; i < len(*lst); i++ {
		if (*lst)[i] == elem {
			e = i
		}
	}
	(*lst)[e] = (*lst)[len(*lst)-1]
	*lst = (*lst)[:len(*lst)-1]
}

func getDigit9(rest *[]string, numMap map[int][]string, number *Number) {
	for _, d := range *rest {
		if isSubset(getSlice(d), append(numMap[4], number.upper)) {
			diff := difference(getSlice(d), append(numMap[4], number.upper))
			number.bottom = diff[0]
			numMap[9] = getSlice(d)
			deleteP(rest, d)
		}
	}
}

func getDigit5And3(rest *[]string, numMap map[int][]string, number *Number) {
	for _, d := range *rest {
		if isSubset(numMap[9], getSlice(d)) {
			diff := difference(numMap[9], getSlice(d))
			if contains(numMap[1], diff[0]) {
				number.topRight = diff[0]
				numMap[5] = getSlice(d)
			} else {
				number.topLeft = diff[0]
				numMap[3] = getSlice(d)
			}
			deleteP(rest, d)
		}
	}
}

func getDigit6And0And2(rest *[]string, numMap map[int][]string, number *Number) {
	for _, d := range *rest {
		if !contains(getSlice(d), number.topRight) {
			numMap[6] = getSlice(d)

		} else {
			diff := difference(numMap[8], getSlice(d))
			if len(diff) == 1 {
				number.middle = diff[0]
				numMap[0] = getSlice(d)
			} else {
				numMap[2] = getSlice(d)
			}
		}
		deleteP(rest, d)
	}
}

func getOutputNumber(line string, c chan int) {
	numMap := map[int][]string{0: nil, 1: nil, 2: nil, 3: nil, 4: nil, 5: nil, 6: nil, 7: nil, 8: nil, 9: nil}
	in, out := getInputOutPutArrays(line)
	rest := make([]string, 0)
	for _, digitStr := range in {
		digit := getSlice(digitStr)
		//fmt.Printf("%+v\n", digit)
		if len(digit) == 2 {
			numMap[1] = digit
		} else if len(digit) == 3 {
			numMap[7] = digit
		} else if len(digit) == 4 {
			numMap[4] = digit
		} else if len(digit) == 7 {
			numMap[8] = digit
		} else {
			rest = append(rest, digitStr)
		}
	}

	number := Number{}

	number.upper = difference(numMap[7], numMap[1])[0]

	getDigit9(&rest, numMap, &number)
	getDigit5And3(&rest, numMap, &number)
	getDigit6And0And2(&rest, numMap, &number)

	mapCopy := make(map[string]int)

	// Copy from the original map to the target map
	for key, value := range numMap {
		sort.Strings(value)
		mapCopy[strings.Join(value, "")] = key
	}

	//fmt.Printf("num  %+v\n",number)
	//fmt.Printf("%+v\n", numMap)
	//fmt.Printf("%+v\n", mapCopy)

	outputNum := make([]string, 0)

	for _, o := range out {
		oo := getSlice(o)
		sort.Strings(oo)
		v := mapCopy[strings.Join(oo, "")]
		outputNum = append(outputNum, strconv.Itoa(v))
	}
	value, _ := strconv.Atoi(strings.Join(outputNum, ""))
	println(value)
	c <- value
}

func run8_2() {
	var input = read_lines("C:\\Users\\tothg\\Gege\\AOC2021\\res\\day_8.txt")
	sum := 0
	c := make(chan int, len(input))
	for _, line := range input {
		getOutputNumber(line, c)
		sum += <-c
	}
	println(sum)
}
