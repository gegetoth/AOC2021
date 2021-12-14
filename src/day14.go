package main

import (
	"fmt"
	"strings"
	"sync"
)

type PolymerInsertion struct {
	pair    string
	element string
}

func getPolymerInput(input []string) ([]string, map[string]string) {
	polymerInsertions := make(map[string]string, 0)
	var initPolymer []string
	isInsertion := false
	for _, line := range input {
		if isInsertion {
			splits := strings.Split(line, " -> ")
			//polymerInsertions=append(polymerInsertions, PolymerInsertion{splits[0], splits[1]})
			polymerInsertions[splits[0]] = splits[1]
		} else if line == "" {
			isInsertion = true
		} else {
			initPolymer = strings.Split(line, "")
		}
	}
	return initPolymer, polymerInsertions
}

func insertionStep(polymer []string, insertions map[string]string) []string {
	newPolymer := make([]string, 0)
	for i := 0; i < len(polymer)-1; i++ {
		double := strings.Join(polymer[i:i+2], "")
		if insertion, ok := insertions[double]; ok {
			if i == 0 {
				newPolymer = append(newPolymer, polymer[0], insertion, polymer[i+1])
			} else {
				newPolymer = append(newPolymer, insertion, polymer[i+1])
			}
		}
	}
	return newPolymer
}

func run14_1() {
	var input = read_lines("C:\\Users\\tothg\\Gege\\AOC2021\\res\\day_14_example.txt")
	initPolymer, polymerInsertions := getPolymerInput(input)

	newP := initPolymer
	for i := 1; i <= 10; i++ {
		newP = insertionStep(newP, polymerInsertions)
		fmt.Printf("Steps: %+v\n", i)
		fmt.Printf("Polimer: %+v\n", newP)
	}

	comonentMap := make(map[string]int, 0)
	for _, e := range newP {
		if _, ok := comonentMap[e]; ok {
			comonentMap[e] += 1
		} else {
			comonentMap[e] = 1
		}
	}
	fmt.Printf("Component Map: %+v\n", comonentMap)
	min := -1
	max := -1

	for _, e := range comonentMap {
		if min == -1 && max == -1 {
			min = e
			max = e
		}
		if min > e {
			min = e
		}

		if max < e {
			max = e
		}
	}
	fmt.Printf("Min: %+v\n", min)
	fmt.Printf("Max: %+v\n", max)
	fmt.Printf("Diff: %+v\n", max-min)
}

func genPolymer(wg *sync.WaitGroup, polymer []string, insertions map[string]string, steps int, subs chan []string) {
	defer wg.Done()
	newP := polymer
	for i := 1; i <= steps; i++ {
		newP = insertionStep(newP, insertions)
		fmt.Printf("Steps: %+v\n", i)
		//fmt.Printf("Polimer: %+v\n", newP)
	}
	subs <- newP
}

func run14_2() {
	var input = read_lines("C:\\Users\\tothg\\Gege\\AOC2021\\res\\day_14_example.txt")
	initPolymer, polymerInsertions := getPolymerInput(input)

	//newP:=initPolymer

	//subs:=make([][]string, 0)
	var wg sync.WaitGroup

	subs := make(chan []string, 3)
	for i := 0; i < len(initPolymer)-1; i++ {
		wg.Add(1)
		go genPolymer(&wg, initPolymer[i:i+2], polymerInsertions, 40, subs)
	}

	println("Wit for the workers to finish")
	wg.Wait()
	close(subs)

	//for i := 1; i <= 40; i++ {
	//	newP = insertionStep(newP, polymerInsertions)
	//	fmt.Printf("Steps: %+v\n", i)
	//	//fmt.Printf("Polimer: %+v\n", newP)
	//}

	basinS := make([]string, 0)
	for elem := range subs {
		basinS = append(basinS, elem...)
	}

	fmt.Printf("Component Map: %+v\n", basinS)

	comonentMap := make(map[string]int, 0)
	for _, e := range basinS {
		if _, ok := comonentMap[e]; ok {
			comonentMap[e] += 1
		} else {
			comonentMap[e] = 1
		}
	}
	fmt.Printf("Component Map: %+v\n", comonentMap)
	min := -1
	max := -1

	for _, e := range comonentMap {
		if min == -1 && max == -1 {
			min = e
			max = e
		}
		if min > e {
			min = e
		}

		if max < e {
			max = e
		}
	}
	fmt.Printf("Min: %+v\n", min)
	fmt.Printf("Max: %+v\n", max)
	fmt.Printf("Diff: %+v\n", max-min)

}
