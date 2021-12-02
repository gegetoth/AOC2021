package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Command struct {
	direction string
	value     int
}

type PositionMap struct {
	horizontal int
	depth      int
	aim        int
}

func (positionMap *PositionMap) commandProcessor(command Command) {
	switch command.direction {
	case "forward":
		positionMap.horizontal += command.value
		positionMap.depth += command.value * positionMap.aim
	case "up":
		positionMap.aim -= command.value
	case "down":
		positionMap.aim += command.value
	default:
		fmt.Println("Invalid direction")
	}
}

func getPosition(input []string) int {
	positionMap := PositionMap{0, 0, 0}

	for _, each_ln := range input {
		ln_split := strings.Split(each_ln, " ")
		direction := ln_split[0]
		value, _ := strconv.Atoi(ln_split[1])

		positionMap.commandProcessor(Command{direction, value})
	}
	return positionMap.horizontal * positionMap.depth
}

func day2_1() {
	var input []string = read_lines("C:\\Users\\tothg\\Gege\\AOC2021\\res\\day_2.txt")
	fmt.Println(getPosition(input))
}

func day2_2() {
	var input []string = read_lines("C:\\Users\\tothg\\Gege\\AOC2021\\res\\day_2.txt")
	fmt.Println(getPosition(input))
}
