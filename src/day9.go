package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
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

type VolcanoField struct {
	i, j, value int
}

func getAdjacentCoords(i, j int, volcanoMatrix [][]int) []VolcanoField {
	adjacent := make([]VolcanoField, 0)
	if i > 0 && j > 0 && i < len(volcanoMatrix)-1 && j < len(volcanoMatrix[0])-1 {
		adjacent = append(adjacent, VolcanoField{i - 1, j, volcanoMatrix[i-1][j]}, VolcanoField{i, j + 1, volcanoMatrix[i][j+1]}, VolcanoField{i + 1, j, volcanoMatrix[i+1][j]}, VolcanoField{i, j - 1, volcanoMatrix[i][j-1]})
	} else if i == 0 {
		if j == 0 {
			adjacent = append(adjacent, VolcanoField{i, j + 1, volcanoMatrix[i][j+1]}, VolcanoField{i + 1, j, volcanoMatrix[i+1][j]})
		} else if j == len(volcanoMatrix[0])-1 {
			adjacent = append(adjacent, VolcanoField{i, j - 1, volcanoMatrix[i][j-1]}, VolcanoField{i + 1, j, volcanoMatrix[i+1][j]})
		} else {
			adjacent = append(adjacent, VolcanoField{i, j + 1, volcanoMatrix[i][j+1]}, VolcanoField{i + 1, j, volcanoMatrix[i+1][j]}, VolcanoField{i, j - 1, volcanoMatrix[i][j-1]})
		}
	} else if i == len(volcanoMatrix)-1 {
		if j == 0 {
			adjacent = append(adjacent, VolcanoField{i - 1, j, volcanoMatrix[i-1][j]}, VolcanoField{i, j + 1, volcanoMatrix[i][j+1]})
		} else if j == len(volcanoMatrix[0])-1 {
			adjacent = append(adjacent, VolcanoField{i, j - 1, volcanoMatrix[i][j-1]}, VolcanoField{i - 1, j, volcanoMatrix[i-1][j]})
		} else {
			adjacent = append(adjacent, VolcanoField{i - 1, j, volcanoMatrix[i-1][j]}, VolcanoField{i, j + 1, volcanoMatrix[i][j+1]}, VolcanoField{i, j - 1, volcanoMatrix[i][j-1]})
		}
	} else if j == 0 {
		adjacent = append(adjacent, VolcanoField{i - 1, j, volcanoMatrix[i-1][j]}, VolcanoField{i, j + 1, volcanoMatrix[i][j+1]}, VolcanoField{i + 1, j, volcanoMatrix[i+1][j]})
	} else {
		adjacent = append(adjacent, VolcanoField{i, j - 1, volcanoMatrix[i][j-1]}, VolcanoField{i - 1, j, volcanoMatrix[i-1][j]}, VolcanoField{i + 1, j, volcanoMatrix[i+1][j]})
	}

	return adjacent
}

func getIfLowPoint(value int, adjacent []VolcanoField) int {
	for i := 0; i < len(adjacent); i++ {
		if value >= adjacent[i].value {
			return 0
		}
	}
	return value + 1
}

func getBasins(volcanoMatrix [][]int) (int, []Basin) {
	basinList := make([]Basin, 0)
	sum := 0
	for i := 0; i < len(volcanoMatrix); i++ {
		row := volcanoMatrix[i]
		for j := 0; j < len(row); j++ {
			adjFields := getAdjacentCoords(i, j, volcanoMatrix)
			//fmt.Printf("%+v, %+v: %+v\n", i, j,adjFields)
			low := getIfLowPoint(volcanoMatrix[i][j], adjFields)
			sum += low
			if low > 0 {
				basinList = append(basinList, Basin{i, j, 0})
			}
		}
		//println("")
	}
	return sum, basinList
}

func run9_1() {
	var input = read_lines("C:\\Users\\tothg\\Gege\\AOC2021\\res\\day_9.txt")

	volcanoMatrix := getVolcanoMatrix(input)
	printIntMatrix(volcanoMatrix)
	sum, _ := getBasins(volcanoMatrix)
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
	visited[VolcanoField{i, j, volcanoMatrix[i][j]}] = true
	for _, adj := range adjacents {
		value := volcanoMatrix[adj.i][adj.j]
		if value > volcanoMatrix[i][j] && value != 9 {
			ret := getSizeOfBasin(adj.i, adj.j, volcanoMatrix, numb+1, visited)
			numb = numb + ret
		}
	}
	return numb
}

func getAdjacentsOfBasin(wg *sync.WaitGroup, basin Basin, volcanoMatrix [][]int, basinSizes chan int) {
	defer wg.Done()
	visited := make(map[VolcanoField]bool, 0)
	getSizeOfBasin(basin.i, basin.j, volcanoMatrix, 0, visited)
	//fmt.Printf("Visited: %+v\n", visited)
	println(len(visited))
	basinSizes <- len(visited)
	println("Worker has finished with processing the basin.")
}

func run9_2() {
	var input = read_lines("C:\\Users\\tothg\\Gege\\AOC2021\\res\\day_9.txt")
	volcanoMatrix := getVolcanoMatrix(input)
	printIntMatrix(volcanoMatrix)
	_, basinList := getBasins(volcanoMatrix)
	var wg sync.WaitGroup

	basinSizes := make(chan int, len(basinList))
	for _, b := range basinList {
		wg.Add(1)
		go getAdjacentsOfBasin(&wg, b, volcanoMatrix, basinSizes)
	}

	println("Wit for the workers to finish")
	wg.Wait()
	close(basinSizes)

	basinS := make([]int, len(basinSizes))
	for elem := range basinSizes {
		basinS = append(basinS, elem)
	}
	sort.Ints(basinS)
	fmt.Printf("Basin sizes: %+v\n", basinS[len(basinS)-1]*basinS[len(basinS)-2]*basinS[len(basinS)-3])

}
