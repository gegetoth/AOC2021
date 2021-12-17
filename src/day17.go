package main

import (
	"fmt"
	"strconv"
)

func getNumSeries(deepest int) int {
	deepest = Abs(deepest)
	height := 0
	for i := deepest; i > 0; i-- {
		height += i
	}
	return height
}

func xRange(value int, from int, to int) bool {
	value = Abs(value)
	height := 0
	for i := value; i > 0; i-- {
		height += i
		if height >= from && height <= to {
			return true
		} else if height > to {
			return false
		}
	}
	return false
}

func run17_1() {
	h := getNumSeries(109)
	fmt.Printf("Height: %+v\n", h)
}

func getPossibleXValues(x1, x2 int) []int {
	xValues := make([]int, 0)
	for i := 0; i <= x2; i++ {
		if xRange(i, x1, x2) {
			xValues = append(xValues, i)
		}
	}
	return xValues
}

func getYValues(maxYVelo int, y2 int) []int {
	yValues := make([]int, 0)
	for i := maxYVelo; i >= y2; i-- {
		yValues = append(yValues, i)
	}
	return yValues
}

func calcPathWithin(x, y, x1, x2, y1, y2 int) bool {
	xc := 0
	yc := 0
	xVelo := x
	yVelo := y
	for xc <= x2 && yc >= y2 {
		if xc >= x1 && xc <= x2 && yc <= y1 && yc >= y2 {
			fmt.Printf("Match values: %+v, %+v\n", xc, yc)
			return true
		}
		xc = xc + xVelo
		yc = yc + yVelo

		if xVelo > 0 {
			xVelo--
		}

		yVelo--
	}
	return false
}

func run17_2() {
	x1 := 156
	x2 := 202
	y1 := -69
	y2 := -110

	//x1:=20
	//x2:=30
	//y1:=-5
	//y2:=-10
	xValues := getPossibleXValues(x1, x2)
	fmt.Printf("X values: %+v\n", xValues)
	yValues := getYValues(109, y2)
	fmt.Printf("Y values: %+v\n", yValues)
	possibeCombo := make(map[string]bool, 0)
	for _, x := range xValues {
		for _, y := range yValues {
			if calcPathWithin(x, y, x1, x2, y1, y2) {
				possibeCombo[strconv.Itoa(x)+","+strconv.Itoa(y)] = true
			}
		}
	}
	fmt.Printf("Possible combos: %+v\n", possibeCombo)
	fmt.Printf("Possible combos length: %+v\n", len(possibeCombo))
}
