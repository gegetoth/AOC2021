package main

import (
	"fmt"
	"strconv"
)

func getNumberOfIncrements(intarray []int){
	var increment int = 0
	var previous int = 0

	for _, each_ln := range intarray {
		//intVar, _ := strconv.Atoi(each_ln)
		if previous != 0 {
			if each_ln > previous{
				increment++
			}
		}
		previous=each_ln
	}
	fmt.Println(increment)
}

func run1_1() {
	var input []string = read_lines()
	var intarray []int

	for _, each_ln := range input {
		intVar, _ := strconv.Atoi(each_ln)
		intarray = append(intarray, intVar)
		//if previous != 0 {
		//	if intVar > previous{
		//		increment++
		//	}
		//}
		//previous=intVar
	}
	getNumberOfIncrements(intarray)
}

func run1_2(){
	var input []string = read_lines()

	var size = len(input)
	var windowArray []int
	for i := 0; i < size-2;  i++{
		intVar, _ := strconv.Atoi(input[i])
		intVar2, _ := strconv.Atoi(input[i+1])
		intVar3, _ := strconv.Atoi(input[i+2])

		windowArray = append(windowArray, intVar + intVar2 + intVar3)
	}

	getNumberOfIncrements(windowArray)

}
