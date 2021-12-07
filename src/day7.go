package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func getHorizontalMap(line string) map[int]int {
	horizontalMap := make(map[int]int)
	for _, horizontal := range strings.Split(line, ",") {
		value, _ := strconv.Atoi(horizontal)
		_, isPresent := horizontalMap[value]
		if isPresent {
			horizontalMap[value] += 1
		} else {
			horizontalMap[value] = 1
		}
	}
	return horizontalMap
}

func makeRangeSum(max int) int {
	sum := 0
	for i := 1; i <= max; i++ {
		sum += i
	}
	return sum
}

func MaxIntSlice(v []int) int {
	sort.Ints(v)
	return v[len(v)-1]
}

func getFuelConsumption(horizontalMap map[int]int) int {
	keys := make([]int, 0, len(horizontalMap))
	for k := range horizontalMap {
		keys = append(keys, k)
	}
	sum := 0
	sumTmp := 0
	var location int

	for i := 0; i <= MaxIntSlice(keys); i++ {
		sumTmp = 0
		location = i
		for k, v := range horizontalMap {
			rangeOfFuels := makeRangeSum(Abs(location - k))
			sumTmp += rangeOfFuels * v

			if sum != 0 {
				if sumTmp > sum {
					break
				}
			}
		}
		if sumTmp < sum || sum == 0 {
			sum = sumTmp
		}
	}
	println(sum)
	return sum
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func run7_1() {
	var input = read_lines("C:\\Users\\tothg\\Gege\\AOC2021\\res\\day_7.txt")
	horizontalMap := getHorizontalMap(input[0])
	fmt.Printf("%+v\n", horizontalMap)
	getFuelConsumption(horizontalMap)
}

func run7_2() {
	var input = read_lines("C:\\Users\\tothg\\Gege\\AOC2021\\res\\day_7.txt")
	horizontalMap := getHorizontalMap(input[0])
	fmt.Printf("%+v\n", horizontalMap)
	getFuelConsumption(horizontalMap)
}
