package main

import (
	"strconv"
	"strings"
)

type TransparentPaperField struct {
	x, y int
}

type FoldDirection int

// Declare related constants for each weekday starting with index 0
const (
	Up   FoldDirection = iota // EnumIndex = 0
	Left                      // EnumIndex = 1
)

func (w FoldDirection) String() string {
	return [...]string{"Up", "Left"}[w]
}

type Fold struct {
	direction FoldDirection
	value     int
}

func getDotsAndFolds(input []string) ([]TransparentPaperField, []Fold, int, int) {
	dots := make([]TransparentPaperField, 0)
	height := 0
	length := 0
	folds := make([]Fold, 0)
	isMap := true
	for _, line := range input {
		if line == "" {
			isMap = false
			break
		}

		if isMap {
			coords := strings.Split(line, ",")
			x, _ := strconv.Atoi(coords[0])
			y, _ := strconv.Atoi(coords[1])
			dots = append(dots, TransparentPaperField{x, y})

			if x > height {
				height = x
			}

			if y > length {
				length = y
			}

		} else {
			instruction := strings.Split(strings.Split(line, " ")[2], "=")
			fold := Fold{}
			if instruction[0] == "y" {
				fold.direction = Up
			} else {
				fold.direction = Left
			}
			v, _ := strconv.Atoi(instruction[1])
			fold.value = v
			folds = append(folds, fold)
		}
	}

	return dots, folds, height + 1, length + 1
}

func getTransparentPaper(dots []TransparentPaperField, height, length int) [][]string {
	paper := make([][]string, height)

	for i := 0; i < height; i++ {
		row := make([]string, length)
		for j := 0; j < length; j++ {
			row[j] = "."
		}
		paper[i] = row
	}

	for _, dot := range dots {
		paper[dot.x][dot.y] = "#"
	}

	return paper
}

func run13_1() {
	var input = read_lines("C:\\Users\\tothg\\Gege\\AOC2021\\res\\day_13_example.txt")
	dots, _, height, length := getDotsAndFolds(input)
	paper := getTransparentPaper(dots, height, length)

	printStrMatrix(paper)
}
