package main

import (
	"fmt"
)

type OctopusMatrixField struct {
	i, j int
}

var adjMap map[OctopusMatrixField][]OctopusMatrixField

func getAdjacentOctopus(i, j int, iSize, jSize int) []OctopusMatrixField {
	adjacentList := make([]OctopusMatrixField, 0)

	xFrom := -1
	xTo := 1
	yFrom := -1
	yTo := 1

	if i == 0 {
		xFrom = 0
		if j == 0 {
			yFrom = 0
		} else if j == jSize-1 {
			yFrom = -1
			yTo = 0
		}
	} else if i == iSize-1 {
		xFrom = -1
		xTo = 0
		if j == 0 {
			yFrom = 0
		} else if j == jSize-1 {
			yFrom = -1
			yTo = 0
		}
	} else if j == 0 {
		yFrom = 0
		yTo = 1
	} else if j == jSize-1 {
		yFrom = -1
		yTo = 0
	}

	for x := xFrom; x <= xTo; x++ {
		for y := yFrom; y <= yTo; y++ {
			if x != 0 || y != 0 {
				adjacentList = append(adjacentList, OctopusMatrixField{i + x, j + y})
			}
		}
	}

	return adjacentList
}

var flashes = 0

func flash(octopusMatrix *[][]int, loadedList *Queue) {
	flashes++
	loaded, _ := loadedList.Get()
	fmt.Printf("Loaded:%+v\n", loaded)
	adjList := adjMap[loaded.(OctopusMatrixField)]

	for _, adj := range adjList {
		if (*octopusMatrix)[adj.i][adj.j] == 9 {
			(*octopusMatrix)[adj.i][adj.j] = 0
			loadedList.Push(adj)
		} else if (*octopusMatrix)[adj.i][adj.j] != 0 {
			(*octopusMatrix)[adj.i][adj.j] += 1
		}
	}

}

func runOctopusCycle(octopusMatrix *[][]int) {
	loadedList := Queue{}
	for i := 0; i < len(*octopusMatrix); i++ {
		for j := 0; j < len((*octopusMatrix)[0]); j++ {
			if (*octopusMatrix)[i][j] != 9 {
				(*octopusMatrix)[i][j]++
			} else {
				(*octopusMatrix)[i][j] = 0
				loadedList.Push(OctopusMatrixField{i, j})
			}
		}
	}

	for !loadedList.IsEmpty() {
		flash(octopusMatrix, &loadedList)
	}
}

func genAdjMatrix(iSize, jSize int) map[OctopusMatrixField][]OctopusMatrixField {
	adjMap := make(map[OctopusMatrixField][]OctopusMatrixField, 0)
	for i := 0; i < iSize; i++ {
		for j := 0; j < jSize; j++ {
			adjMap[OctopusMatrixField{i, j}] = getAdjacentOctopus(i, j, iSize, jSize)
		}
	}
	return adjMap
}

func run11_1() {
	var input = read_lines("C:\\Users\\tothg\\Gege\\AOC2021\\res\\day_11.txt")
	octopusMatrix := getIntMatrix(input)
	adjMap = genAdjMatrix(len(octopusMatrix), len(octopusMatrix[0]))
	fmt.Printf("Adjacent map:%+v\n", adjMap)

	for i := 1; i <= 100; i++ {
		runOctopusCycle(&octopusMatrix)
		fmt.Printf("Matrix after round: %+v\n", i)
		printIntMatrix(octopusMatrix)

	}

	fmt.Printf("All flashes: %+v\n", flashes)

}

func sumOctopusLoad(octopusMatrix *[][]int) int {
	sum := 0
	for i := 0; i < len(*octopusMatrix); i++ {
		for j := 0; j < len((*octopusMatrix)[0]); j++ {
			sum += (*octopusMatrix)[i][j]
		}
	}
	return sum
}

func run11_2() {
	var input = read_lines("C:\\Users\\tothg\\Gege\\AOC2021\\res\\day_11.txt")
	octopusMatrix := getIntMatrix(input)
	adjMap = genAdjMatrix(len(octopusMatrix), len(octopusMatrix[0]))
	fmt.Printf("Adjacent map:%+v\n", adjMap)

	i := 1
	for sumOctopusLoad(&octopusMatrix) != 0 {
		runOctopusCycle(&octopusMatrix)
		fmt.Printf("Matrix after round: %+v\n", i)
		printIntMatrix(octopusMatrix)
		i++
	}
	fmt.Printf("All flashes: %+v\n", flashes)
}
