package main

import (
	"strconv"
	"strings"
)

func getGamma(input [][]string, defaultValue ...string) []string {
	size := len(input)
	gamma := make([]string, len(input[0]))
	for j := 0; j < len(input[0]); j++ {
		digit := 0
		for i := 0; i < len(input); i++ {
			value, _ := strconv.Atoi(input[i][j])
			digit += value
		}
		if float32(digit) > float32(size)/float32(2) {
			gamma[j] = "1"
		} else if float32(digit) < float32(size)/float32(2) {
			gamma[j] = "0"
		} else {
			gamma[j] = defaultValue[0]
		}
	}

	return gamma
}

func getMatrix(input []string) [][]string {
	var intMatrix [][]string
	for _, each_ln := range input {
		chars := []rune(each_ln)
		intArr := make([]string, len(chars))
		for i := 0; i < len(chars); i++ {
			intArr[i] = string(chars[i])
		}
		intMatrix = append(intMatrix, intArr)
	}

	printStrMatrix(intMatrix)

	return intMatrix
}

func invertBits(bitArray []string) []string {
	newArr := make([]string, len(bitArray))

	for i := 0; i < len(bitArray); i++ {
		if bitArray[i] == "0" {
			newArr[i] = "1"
		} else {
			newArr[i] = "0"
		}
	}
	return newArr
}
func run_3_1() {
	var input []string = read_lines("C:\\Users\\tothg\\Gege\\AOC2021\\res\\day_3.txt")
	intMatrix := getMatrix(input)
	gamma := getGamma(intMatrix)
	output, _ := strconv.ParseInt(strings.Join(gamma[:], ""), 2, 64)
	println("Gamma: " + strconv.FormatInt(output, 10))
	epsilon := invertBits(gamma)
	output2, _ := strconv.ParseInt(strings.Join(epsilon[:], ""), 2, 64)
	println("Epsilon: " + strconv.FormatInt(output2, 10))

	println(strconv.FormatInt(output*output2, 10))
}

func run_3_2() {
	var input = read_lines("C:\\Users\\tothg\\Gege\\AOC2021\\res\\day_3.txt")
	intMatrix := getMatrix(input)
	oxy := getOxygenAndCo2Rating(intMatrix, false)
	co2 := getOxygenAndCo2Rating(intMatrix, true)
	println("Val: " + strconv.FormatInt(oxy*co2, 10))

}

func getOxygenAndCo2Rating(intMatrix [][]string, invert bool) int64 {
	i := 0
	gamma := getGamma(intMatrix, "1")
	newMatrix := intMatrix
	for len(newMatrix) > 1 {
		printStrMatrix(newMatrix)
		var tmp [][]string
		significant := gamma[i]

		if invert {
			if significant == "0" {
				significant = "1"
			} else {
				significant = "0"
			}
		}

		for _, each_ln := range newMatrix {
			if each_ln[i] == significant {
				tmp = append(tmp, each_ln)
			}
		}

		newMatrix = tmp
		gamma = getGamma(newMatrix, "1")
		i++
	}

	output, _ := strconv.ParseInt(strings.Join(newMatrix[0][:], ""), 2, 64)
	println("Oxygen/CO2: " + strconv.FormatInt(output, 10))
	return output
}
