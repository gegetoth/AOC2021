package main

import (
	"fmt"
	"strconv"
	"strings"
)

func createLanterfishMap(input string) map[int]int {
	populationMap := make(map[int]int)
	populationMap[0] = 0
	populationMap[1] = 0
	populationMap[2] = 0
	populationMap[3] = 0
	populationMap[4] = 0
	populationMap[5] = 0
	populationMap[6] = 0
	populationMap[7] = 0
	populationMap[8] = 0

	for _, fish := range strings.Split(input, ",") {
		value, _ := strconv.Atoi(fish)
		populationMap[value] += 1
	}

	return populationMap
}

func modelPopulationCycle(populationMap map[int]int) map[int]int {
	mapCopy := make(map[int]int)

	// Copy from the original map to the target map
	for key, value := range populationMap {
		mapCopy[key] = value
	}

	for i := 8; i > -1; i-- {
		if i > 0 {
			mapCopy[i-1] = populationMap[i]
		} else {
			mapCopy[8] = populationMap[0]
			mapCopy[6] += populationMap[0]
		}

	}

	return mapCopy
}

func run6_1() {
	var input = read_lines("C:\\Users\\tothg\\Gege\\AOC2021\\res\\day_6.txt")
	populationMap := createLanterfishMap(input[0])
	fmt.Printf("%+v\n", populationMap)

	for i := 0; i < 256; i++ {
		populationMap = modelPopulationCycle(populationMap)
		fmt.Printf("%+v\n", populationMap)
	}

	var sum int64 = 0
	for _, value := range populationMap {
		sum += int64(value)
	}
	println(sum)
}
