package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func getVolcanoMatrix(input []string) [][]int {
	volcanoMatrix := make([][]int, 0)
	for _, line := range input {
		elemStr := strings.Split(line, "")
		row := make([]int, len(elemStr))
		for i := 0; i < len(elemStr); i++ {
			value, _ := strconv.Atoi(elemStr[i])
			row[i] = value
		}
		volcanoMatrix = append(volcanoMatrix, row)
	}
	return volcanoMatrix
}

func getAdjacentFields(i, j int, volcanoMatrix [][]int) []int {
	adjacent := make([]int, 0)
	if i > 0 && j > 0 && i < len(volcanoMatrix)-1 && j < len(volcanoMatrix[0])-1 {
		adjacent = append(adjacent, volcanoMatrix[i-1][j], volcanoMatrix[i][j+1], volcanoMatrix[i+1][j], volcanoMatrix[i][j-1])
	} else if i == 0 {
		if j == 0 {
			adjacent = append(adjacent, volcanoMatrix[i][j+1], volcanoMatrix[i+1][j])
		} else if j == len(volcanoMatrix[0])-1 {
			adjacent = append(adjacent, volcanoMatrix[i][j-1], volcanoMatrix[i+1][j])
		} else {
			adjacent = append(adjacent, volcanoMatrix[i][j+1], volcanoMatrix[i+1][j], volcanoMatrix[i][j-1])
		}
	} else if i == len(volcanoMatrix)-1 {
		if j == 0 {
			adjacent = append(adjacent, volcanoMatrix[i-1][j], volcanoMatrix[i][j+1])
		} else if j == len(volcanoMatrix[0])-1 {
			adjacent = append(adjacent, volcanoMatrix[i][j-1], volcanoMatrix[i-1][j])
		} else {
			adjacent = append(adjacent, volcanoMatrix[i-1][j], volcanoMatrix[i][j+1], volcanoMatrix[i][j-1])
		}
	} else if j == 0 {
		adjacent = append(adjacent, volcanoMatrix[i-1][j], volcanoMatrix[i][j+1], volcanoMatrix[i+1][j])
	} else {
		adjacent = append(adjacent, volcanoMatrix[i][j-1], volcanoMatrix[i-1][j], volcanoMatrix[i+1][j])
	}

	return adjacent
}

type VolcanoField struct {
	i, j int
}

func getAdjacentCoords(i, j int, volcanoMatrix [][]int) []VolcanoField {
	adjacent := make([]VolcanoField, 0)
	if i > 0 && j > 0 && i < len(volcanoMatrix)-1 && j < len(volcanoMatrix[0])-1 {
		adjacent = append(adjacent, VolcanoField{i - 1, j}, VolcanoField{i, j + 1}, VolcanoField{i + 1, j}, VolcanoField{i, j - 1})
	} else if i == 0 {
		if j == 0 {
			adjacent = append(adjacent, VolcanoField{i, j + 1}, VolcanoField{i + 1, j})
		} else if j == len(volcanoMatrix[0])-1 {
			adjacent = append(adjacent, VolcanoField{i, j - 1}, VolcanoField{i + 1, j})
		} else {
			adjacent = append(adjacent, VolcanoField{i, j + 1}, VolcanoField{i + 1, j}, VolcanoField{i, j - 1})
		}
	} else if i == len(volcanoMatrix)-1 {
		if j == 0 {
			adjacent = append(adjacent, VolcanoField{i - 1, j}, VolcanoField{i, j + 1})
		} else if j == len(volcanoMatrix[0])-1 {
			adjacent = append(adjacent, VolcanoField{i, j - 1}, VolcanoField{i - 1, j})
		} else {
			adjacent = append(adjacent, VolcanoField{i - 1, j}, VolcanoField{i, j + 1}, VolcanoField{i, j - 1})
		}
	} else if j == 0 {
		adjacent = append(adjacent, VolcanoField{i - 1, j}, VolcanoField{i, j + 1}, VolcanoField{i + 1, j})
	} else {
		adjacent = append(adjacent, VolcanoField{i, j - 1}, VolcanoField{i - 1, j}, VolcanoField{i + 1, j})
	}

	return adjacent
}

func getIfLowPoint(value int, adjacent []int) int {
	for i := 0; i < len(adjacent); i++ {
		if value >= adjacent[i] {
			return 0
		}
	}
	return value + 1
}

func run9_1() {
	var input = read_lines("C:\\Users\\tothg\\Gege\\AOC2021\\res\\day_9.txt")

	volcanoMatrix := getVolcanoMatrix(input)
	printIntMatrix(volcanoMatrix)
	sum := 0
	for i := 0; i < len(volcanoMatrix); i++ {
		row := volcanoMatrix[i]
		for j := 0; j < len(row); j++ {
			adjFields := getAdjacentFields(i, j, volcanoMatrix)
			//fmt.Printf("%+v, %+v: %+v\n", i, j,adjFields)
			sum += getIfLowPoint(volcanoMatrix[i][j], adjFields)
		}
		println("")
	}
	println(sum)
}

type Basin struct {
	i, j int
	size int
}

func deleteVisited(lst *([]VolcanoField), elem VolcanoField) {
	e := -1
	for i := 0; i < len(*lst); i++ {
		if (*lst)[i].i == elem.i && (*lst)[i].j == elem.j {
			e = i
		}
	}
	if e != -1 {
		(*lst)[e] = (*lst)[len(*lst)-1]
		*lst = (*lst)[:len(*lst)-1]
	}
}

func removeVisitedFromAdj(adj []VolcanoField, visited map[VolcanoField]bool) []VolcanoField {
	for k, _ := range visited {
		deleteVisited(&adj, k)
	}
	return adj
}

func getSizeOfBasin(i, j int, volcanoMatrix [][]int, numb int, visited map[VolcanoField]bool) int {
	adjacents := getAdjacentCoords(i, j, volcanoMatrix)
	fmt.Printf("Current fields: %+v, %+v\n", i, j)
	fmt.Printf("Visited depth: %+v\n", visited)
	fmt.Printf("Adjecent before:  %+v\n", adjacents)
	adjacents = removeVisitedFromAdj(adjacents, visited)
	fmt.Printf("Adjecent fields after removal:  %+v\n", adjacents)
	fmt.Printf("Current depth: %+v\n", numb)
	fmt.Printf("\n")
	visited[VolcanoField{i, j}] = true
	for _, adj := range adjacents {
		value := volcanoMatrix[adj.i][adj.j]
		if value > volcanoMatrix[i][j] && value != 9 {
			//*visited=append(*visited, VolcanoField{i, j})
			ret := getSizeOfBasin(adj.i, adj.j, volcanoMatrix, numb+1, visited)
			numb = numb + ret
		}
	}
	return numb
}

func run9_2() {
	var input = read_lines("C:\\Users\\tothg\\Gege\\AOC2021\\res\\day_9.txt")
	volcanoMatrix := getVolcanoMatrix(input)
	printIntMatrix(volcanoMatrix)
	basinList := make([]Basin, 0)
	for i := 0; i < len(volcanoMatrix); i++ {
		row := volcanoMatrix[i]
		for j := 0; j < len(row); j++ {
			adjFields := getAdjacentFields(i, j, volcanoMatrix)
			//fmt.Printf("%+v, %+v: %+v\n", i, j,adjFields)
			low := getIfLowPoint(volcanoMatrix[i][j], adjFields)
			if low > 0 {
				basinList = append(basinList, Basin{i, j, 0})
			}
		}
		//println("")
	}

	basinSizes := make([]int, 0)
	for _, b := range basinList {
		visited := make(map[VolcanoField]bool, 0)
		getSizeOfBasin(b.i, b.j, volcanoMatrix, 0, visited)
		//fmt.Printf("Visited: %+v\n", visited)
		println(len(visited))
		basinSizes = append(basinSizes, len(visited))
	}
	sort.Ints(basinSizes)
	fmt.Printf("Basin sizes: %+v\n", basinSizes)
	fmt.Printf("Basin sizes: %+v\n", basinSizes[len(basinSizes)-1]*basinSizes[len(basinSizes)-2]*basinSizes[len(basinSizes)-3])

}
