package main

import (
	"fmt"
	"strings"
)

type Cave struct {
	name     string
	isBig    bool
	adjCaves map[string]*Cave
}

var caveMap = make(map[string]*Cave)

func readCaveMapInput(input []string) Cave {
	for _, line := range input {
		nodes := strings.Split(line, "-")
		from, ok := caveMap[nodes[0]]
		if !ok {
			from = &Cave{nodes[0], strings.ToUpper(nodes[0]) == nodes[0], make(map[string]*Cave, 0)}
		}
		to, ok := caveMap[nodes[1]]
		if !ok {
			to = &Cave{nodes[1], strings.ToUpper(nodes[1]) == nodes[1], make(map[string]*Cave, 0)}
		}
		caveMap[nodes[0]] = from
		caveMap[nodes[1]] = to
		from.adjCaves[nodes[1]] = to
		to.adjCaves[nodes[0]] = from
	}
	return *caveMap["start"]
}

var count = 0

func isSmallVisitedTwice(visitedCaves map[string]int) bool {
	for _, value := range visitedCaves {
		if value == 2 {
			return true
		}
	}
	return false
}

func DFS(cave Cave, path []Cave, visitedCaves map[string]int) {
	visitedCopy := make(map[string]int, len(visitedCaves))
	for k, v := range visitedCaves {
		visitedCopy[k] = v
	}
	if cave.name == "end" {
		path = append(path, cave)
		fmt.Printf("Path: %+v\n", path)
		count++
		return
	}

	_, ok := visitedCopy[cave.name]
	if cave.name == strings.ToLower(cave.name) && !ok {
		path = append(path, cave)
		visitedCopy[cave.name] = 1
		for _, adj := range cave.adjCaves {
			_, ok := visitedCopy[adj.name]
			if !ok {
				DFS(*adj, path, visitedCopy)
			}
		}
	} else if cave.name == strings.ToUpper(cave.name) {
		path = append(path, cave)
		for _, adj := range cave.adjCaves {
			_, ok := visitedCopy[adj.name]
			if !ok {
				DFS(*adj, path, visitedCopy)
			}
		}
	}

}

func run12_1() {
	var input = read_lines("C:\\Users\\tothg\\Gege\\AOC2021\\res\\day_12.txt")
	startCave := readCaveMapInput(input)

	fmt.Printf("Start cave: %+v\n", startCave)

	for _, v := range caveMap {
		fmt.Printf("Cave: %+v\n", v)
	}

	DFS(startCave, make([]Cave, 0), make(map[string]int, 0))
	fmt.Printf("Count: %+v\n", count)
}

func run12_2() {
	var input = read_lines("C:\\Users\\tothg\\Gege\\AOC2021\\res\\day_12_example_1.txt")
	startCave := readCaveMapInput(input)

	fmt.Printf("Start cave: %+v\n", startCave)

	for _, v := range caveMap {
		fmt.Printf("Cave: %+v\n", v)
	}

	DFS(startCave, make([]Cave, 0), make(map[string]int, 0))
	fmt.Printf("Count: %+v\n", count)
}
