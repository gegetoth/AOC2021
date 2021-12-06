package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type VentCoordinates struct {
	x int
	y int
}

type VentLine struct {
	v1 VentCoordinates
	v2 VentCoordinates
}

type VentMap struct {
	size   int
	fields [][]int
}

func (m *VentMap) Init(size int) {
	m.size = size
	fields := make([][]int, size)
	for i := 0; i < size; i++ {
		fields[i] = make([]int, size)
	}
	m.fields = fields
}

func getVentMapAndLines(input []string, onlyHorizontalVertical bool) ([]VentLine, VentMap) {
	size := 0
	ventLines := make([]VentLine, 0)
	for _, line := range input {
		split := strings.Split(line, " -> ")
		vert1 := strings.Split(split[0], ",")
		vert2 := strings.Split(split[1], ",")
		x1, _ := strconv.Atoi(vert1[0])
		y1, _ := strconv.Atoi(vert1[1])
		x2, _ := strconv.Atoi(vert2[0])
		y2, _ := strconv.Atoi(vert2[1])

		max := getMax(x1, x2, y1, y2, size)
		if size < max {
			size = max
		}

		if onlyHorizontalVertical {
			if x1 == x2 || y1 == y2 {
				ventCoord1 := VentCoordinates{x1, y1}
				ventCoord2 := VentCoordinates{x2, y2}

				ventLine := VentLine{ventCoord1, ventCoord2}
				ventLines = append(ventLines, ventLine)
			}
		} else {
			ventCoord1 := VentCoordinates{x1, y1}
			ventCoord2 := VentCoordinates{x2, y2}

			ventLine := VentLine{ventCoord1, ventCoord2}
			ventLines = append(ventLines, ventLine)
		}
	}

	ventMap := VentMap{}
	ventMap.Init(size + 1)
	return ventLines, ventMap
}

func getMax(arr ...int) int {
	sort.Ints(arr[:])
	return arr[len(arr)-1]
}

func printVentMap(ventMap VentMap) {
	for i := 0; i < len(ventMap.fields); i++ {
		row := ventMap.fields[i]
		for j := 0; j < len(row); j++ {
			if ventMap.fields[i][j] == 0 {
				print(".")
			} else {
				print(ventMap.fields[i][j])
			}
		}
		println("")
	}
}

func step(v1 int, v2 int) int {
	if v1 < v2 {
		return 1
	} else if v1 == v2 {
		return 0
	}
	return -1
}

func getCoordsFromLine(ventLine VentLine) []VentCoordinates {
	x1 := ventLine.v1.x
	y1 := ventLine.v1.y
	x2 := ventLine.v2.x
	y2 := ventLine.v2.y

	fields := make([]VentCoordinates, 0)
	xStep := step(x1, x2)
	yStep := step(y1, y2)
	xPrev := x1
	yPrev := y1

	fields = append(fields, VentCoordinates{x1, y1})
	for xPrev != x2 || yPrev != y2 {
		fields = append(fields, VentCoordinates{xPrev + xStep, yPrev + yStep})
		xPrev = xPrev + xStep
		yPrev = yPrev + yStep
	}
	return fields
}

func getOverlaps(ventLines []VentLine, ventMap VentMap, isPrintMap bool) {
	overlapCounter := 0
	for _, line := range ventLines {
		fmt.Printf("%+v\n", line)
		coords := getCoordsFromLine(line)
		fmt.Printf("%+v\n", coords)

		for _, coord := range coords {
			if ventMap.fields[coord.y][coord.x] == 1 {
				overlapCounter++
			}
			ventMap.fields[coord.y][coord.x]++
		}
	}

	if isPrintMap {
		printVentMap(ventMap)
	}
	println(overlapCounter)
}

func run5_1() {
	var input = read_lines("C:\\Users\\tothg\\Gege\\AOC2021\\res\\day_5_example.txt")
	ventLines, ventMap := getVentMapAndLines(input, true)
	getOverlaps(ventLines, ventMap, true)
}

func run5_2() {
	var input = read_lines("C:\\Users\\tothg\\Gege\\AOC2021\\res\\day_5.txt")
	ventLines, ventMap := getVentMapAndLines(input, false)
	getOverlaps(ventLines, ventMap, false)

}
