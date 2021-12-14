package main

import (
	"fmt"
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
			continue
		}

		if isMap {
			coords := strings.Split(line, ",")
			x, _ := strconv.Atoi(coords[0])
			y, _ := strconv.Atoi(coords[1])
			dots = append(dots, TransparentPaperField{x, y})

			if x > length {
				length = x
			}

			if y > height {
				height = y
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
		paper[dot.y][dot.x] = "#"
	}

	return paper
}

func flipPaperUpDown(paper [][]string) [][]string {
	paperCopy := make([][]string, 0)
	for _, row := range paper {
		rowCopy := row[:]
		paperCopy = append(paperCopy, rowCopy)
	}

	for i := 0; i < len(paper); i++ {
		for j := 0; j < len(paper[0]); j++ {
			paperCopy[i][j] = paper[len(paper)-1-i][j]
		}
	}
	return paperCopy
}

func flipPaperRightToLeft(paper [][]string) [][]string {
	paperCopy := make([][]string, 0)

	for _, row := range paper {
		rowCopy := row[:]
		paperCopy = append(paperCopy, rowCopy)
	}

	for i := 0; i < len(paper); i++ {
		for j := 0; j < len(paper[0]); j++ {
			paperCopy[i][j] = paper[i][len(paper[0])-1-j]
		}
	}
	return paperCopy
}

func foldPaper(fold Fold, paper [][]string) [][]string {
	var afterFold [][]string
	if fold.direction == Up {
		//Completely useless section
		//if int(math.Round(float64(len(paper)/2)))-1 > fold.value {
		//	fmt.Printf("Before Flip: %+v\n", fold.value)
		//	paper=flipPaperUpDown(paper)
		//	fold.value=len(paper)-fold.value-1
		//	fmt.Printf("After Flip: %+v\n", fold.value)
		//}
		afterFold = foldUp(fold, paper)
	} else {
		//Completely useless section
		//if int(math.Round(float64(len(paper[0])/2)))-1 > fold.value {
		//	fmt.Printf("Before Flip: %+vS: %+v\n", fold.value, len(paper[0]))
		//	paper=flipPaperRightToLeft(paper)
		//	fold.value=len(paper[0])-fold.value-1
		//	fmt.Printf("After Flip: %+v\n", fold.value)
		//
		//}
		afterFold = foldLeft(fold, paper)
	}

	return afterFold
}

func foldUp(fold Fold, paper [][]string) [][]string {
	paperCopy := make([][]string, 0)

	// Copy from the original map to the target map
	for _, row := range paper {
		rowCopy := row[:]
		paperCopy = append(paperCopy, rowCopy)
	}

	for i := fold.value; i < len(paper); i++ {
		for j := 0; j < len(paper[0]); j++ {
			if paper[i][j] == "#" {

				paperCopy[fold.value-(i-fold.value)][j] = paper[i][j]
			}
		}
	}

	return paperCopy[:fold.value]
}

func foldLeft(fold Fold, paper [][]string) [][]string {
	paperCopy := make([][]string, 0)

	// Copy from the original map to the target map
	for _, row := range paper {
		rowCopy := row[:fold.value]
		paperCopy = append(paperCopy, rowCopy)
	}

	for i := 0; i < len(paper); i++ {
		for j := fold.value; j < len(paper[0]); j++ {
			if paper[i][j] == "#" {
				paperCopy[i][fold.value-(j-fold.value)] = paper[i][j]
			}
		}
	}

	return paperCopy[:][:]
}

func run13_1() {
	var input = read_lines("C:\\Users\\tothg\\Gege\\AOC2021\\res\\day_13.txt")
	dots, folds, height, length := getDotsAndFolds(input)
	paper := getTransparentPaper(dots, height, length)
	fold := folds[0]
	//printStrMatrix(paper)

	afterFold := foldPaper(fold, paper)
	//printStrMatrix(afterFold)

	count := 0
	for i := 0; i < len(afterFold); i++ {
		for j := 0; j < len(afterFold[0]); j++ {
			if afterFold[i][j] == "#" {
				count++
			}
		}
	}
	fmt.Printf("%+v", count)
}

func run13_2() {
	var input = read_lines("C:\\Users\\tothg\\Gege\\AOC2021\\res\\day_13.txt")
	dots, folds, height, length := getDotsAndFolds(input)
	paper := getTransparentPaper(dots, height, length)
	//printStrMatrix(paper)

	afterFold := paper
	for _, fold := range folds {
		afterFold = foldPaper(fold, afterFold)
	}
	printStrMatrix(afterFold)

}
