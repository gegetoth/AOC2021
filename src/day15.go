package main

import (
	"fmt"
	"strconv"
	"strings"
)

var MAX_INT = 2147483647

func getRiskLevelMatrix(input []string) [][]int {
	riskLevelMatrix := make([][]int, 0)

	for _, line := range input {
		riskStr := strings.Split(line, "")
		row := make([]int, len(riskStr))
		for i := 0; i < len(riskStr); i++ {
			val, _ := strconv.Atoi(riskStr[i])
			row[i] = val
		}
		riskLevelMatrix = append(riskLevelMatrix, row)
	}

	return riskLevelMatrix
}

type RiskMatrixField struct {
	i    int
	j    int
	risk int
	cost int
}

func getRiskLevelFieldMatrix(riskLevelMatrix [][]int) [][]RiskMatrixField {
	riskLevelFieldMatrix := make([][]RiskMatrixField, len(riskLevelMatrix))
	for i := 0; i < len(riskLevelMatrix); i++ {
		row := make([]RiskMatrixField, len(riskLevelMatrix[0]))
		for j := 0; j < len(riskLevelMatrix[0]); j++ {
			row[j] = RiskMatrixField{i, j, riskLevelMatrix[i][j], MAX_INT}
		}
		riskLevelFieldMatrix[i] = row
	}
	return riskLevelFieldMatrix
}

func getAdjacenRiskFields(i, j int, matrix [][]RiskMatrixField) [][2]int {
	adjacent := make([][2]int, 0)
	if i > 0 && j > 0 && i < len(matrix)-1 && j < len(matrix[0])-1 {
		adjacent = append(adjacent, [2]int{i - 1, j}, [2]int{i, j + 1}, [2]int{i + 1, j}, [2]int{i, j - 1})
	} else if i == 0 {
		if j == 0 {
			adjacent = append(adjacent, [2]int{i, j + 1}, [2]int{i + 1, j})
		} else if j == len(matrix[0])-1 {
			adjacent = append(adjacent, [2]int{i, j - 1}, [2]int{i + 1, j})
		} else {
			adjacent = append(adjacent, [2]int{i, j + 1}, [2]int{i + 1, j}, [2]int{i, j - 1})
		}
	} else if i == len(matrix)-1 {
		if j == 0 {
			adjacent = append(adjacent, [2]int{i - 1, j}, [2]int{i, j + 1})
		} else if j == len(matrix[0])-1 {
			adjacent = append(adjacent, [2]int{i, j - 1}, [2]int{i - 1, j})
		} else {
			adjacent = append(adjacent, [2]int{i - 1, j}, [2]int{i, j + 1}, [2]int{i, j - 1})
		}
	} else if j == 0 {
		adjacent = append(adjacent, [2]int{i - 1, j}, [2]int{i, j + 1}, [2]int{i + 1, j})
	} else {
		adjacent = append(adjacent, [2]int{i, j - 1}, [2]int{i - 1, j}, [2]int{i + 1, j})
	}

	return adjacent
}

func runRiskDijkstra(riskLevelFieldMatrix *[][]RiskMatrixField) {
	visitedMap := make(map[string]bool, 0)
	possibleNextSteps := make(map[string]*RiskMatrixField, 0)
	current := [2]int{0, 0}
	(*riskLevelFieldMatrix)[0][0].cost = 0
	possibleNextSteps[strconv.Itoa(current[0])+","+strconv.Itoa(current[1])] = &(*riskLevelFieldMatrix)[current[0]][current[1]]

	for current[0] != len(*riskLevelFieldMatrix)-1 || current[1] != len((*riskLevelFieldMatrix)[0])-1 {
		delete(possibleNextSteps, strconv.Itoa(current[0])+","+strconv.Itoa(current[1]))
		i := current[0]
		j := current[1]
		currentStrKey := strconv.Itoa(current[0]) + "," + strconv.Itoa(current[1])
		visitedMap[currentStrKey] = true

		adjList := getAdjacenRiskFields(i, j, *riskLevelFieldMatrix)

		for _, adj := range adjList {
			if _, ok := visitedMap[strconv.Itoa(adj[0])+","+strconv.Itoa(adj[1])]; !ok {
				possibleNextSteps[strconv.Itoa(adj[0])+","+strconv.Itoa(adj[1])] = &(*riskLevelFieldMatrix)[adj[0]][adj[1]]

				if (*riskLevelFieldMatrix)[current[0]][current[1]].cost+(*riskLevelFieldMatrix)[adj[0]][adj[1]].risk < (*riskLevelFieldMatrix)[adj[0]][adj[1]].cost {
					(*riskLevelFieldMatrix)[adj[0]][adj[1]].cost = (*riskLevelFieldMatrix)[current[0]][current[1]].cost + (*riskLevelFieldMatrix)[adj[0]][adj[1]].risk
				}
			}
		}

		current = getNextField(possibleNextSteps, current)
	}

}

func getNextField(steps map[string]*RiskMatrixField, current [2]int) [2]int {
	min := MAX_INT
	var next [2]int
	for _, v := range steps {
		if min > v.cost {
			min = v.cost
			next = [2]int{v.i, v.j}
		}
	}
	return next
}

func run15_1() {
	var input = read_lines("C:\\Users\\tothg\\Gege\\AOC2021\\res\\day_15.txt")
	riskLevelMatrix := getRiskLevelMatrix(input)
	printIntMatrix(riskLevelMatrix)
	riskLevelFieldMatrix := getRiskLevelFieldMatrix(riskLevelMatrix)
	runRiskDijkstra(&riskLevelFieldMatrix)

	fmt.Printf("%+v\n", riskLevelFieldMatrix)
	fmt.Printf("%+v", riskLevelFieldMatrix[len(riskLevelFieldMatrix)-1][len(riskLevelMatrix[0])-1])
}

func createHorizontalBlock(base [][]RiskMatrixField) [][]RiskMatrixField {
	bigMap := make([][]RiskMatrixField, 0)

	for ii := 0; ii < len(base); ii++ {
		bigRow := make([]RiskMatrixField, 0)
		for j := 0; j < 5; j++ {
			row := base[ii]
			for jj := 0; jj < len(row); jj++ {
				newRisk := 0
				if row[jj].risk+(j) > 9 {
					newRisk = row[jj].risk + (j) - 9
				} else {
					newRisk = row[jj].risk + (j)
				}
				field := RiskMatrixField{ii, (len(base))*j + jj, newRisk, MAX_INT}
				bigRow = append(bigRow, field)
			}
		}
		bigMap = append(bigMap, bigRow)
	}
	return bigMap
}

func vertically(horizontalBlock [][]RiskMatrixField) [][]RiskMatrixField {
	bigMap := make([][]RiskMatrixField, 0)

	for ii := 0; ii < 5; ii++ {
		for i := 0; i < len(horizontalBlock); i++ {
			bigRow := make([]RiskMatrixField, 0)
			row := horizontalBlock[i]
			for j := 0; j < len(row); j++ {
				newRisk := 0
				if row[j].risk+(ii) > 9 {
					newRisk = row[j].risk + (ii) - 9
				} else {
					newRisk = row[j].risk + (ii)
				}
				field := RiskMatrixField{(len(horizontalBlock))*ii + i, j, newRisk, MAX_INT}
				bigRow = append(bigRow, field)
			}
			bigMap = append(bigMap, bigRow)
		}
	}
	return bigMap
}

func run15_2() {
	var input = read_lines("C:\\Users\\tothg\\Gege\\AOC2021\\res\\day_15.txt")
	riskLevelMatrix := getRiskLevelMatrix(input)

	//printIntMatrix(riskLevelMatrix)
	riskLevelFieldMatrix := getRiskLevelFieldMatrix(riskLevelMatrix)
	horizontal := createHorizontalBlock(riskLevelFieldMatrix)
	//fmt.Printf("%+v\n", horizontal)
	bigMap := vertically(horizontal)
	//fmt.Printf("%+v\n", bigMap)

	runRiskDijkstra(&bigMap)

	//fmt.Printf("%+v\n", riskLevelFieldMatrix)
	fmt.Printf("%+v", bigMap[len(bigMap)-1][len(bigMap[0])-1])
}
