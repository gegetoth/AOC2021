package main

import (
	"fmt"
	"sort"
	"strings"
)

var parenthesisMap = map[string]string{
	"(": ")",
	"{": "}",
	"[": "]",
	"<": ">",
	")": "(",
	"}": "{",
	"]": "[",
	">": "<",
}

func getCharsEnding(line string) (string, []string) {
	var openCharStack Stack
	missingOnes := make([]string, 0)
	for _, ch := range strings.Split(line, "") {
		if ch == "(" || ch == "{" || ch == "[" || ch == "<" {
			openCharStack.Push(ch)
		} else if ch == ")" || ch == "}" || ch == "]" || ch == ">" {
			lastOpen, notEmpty := openCharStack.Pop()
			if notEmpty && ch != parenthesisMap[lastOpen.(string)] {
				return ch, missingOnes
			}
		} else {
			return "", []string{"Error occurred."}
		}
	}

	for len(openCharStack) > 0 {
		openChar, notEmpty := openCharStack.Pop()
		if notEmpty == true {
			missingOnes = append(missingOnes, parenthesisMap[openChar.(string)])

		}
	}

	return "", missingOnes
}

var multiplier = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

func run10_1() {
	var input = read_lines("C:\\Users\\tothg\\Gege\\AOC2021\\res\\day_10.txt")

	corruptChars := make([]string, 0)
	sum := 0
	for _, line := range input {
		ch, _ := getCharsEnding(line)
		sum += multiplier[ch]
		corruptChars = append(corruptChars, ch)
	}
	fmt.Printf("List of corrupt chars: %+v\n", corruptChars)
	fmt.Printf("Score: %+v\n", sum)
}

var charVal = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
}

func autoCompleteScore(chars []string) int {
	score := 0
	for _, ch := range chars {
		score = score*5 + charVal[ch]
	}
	return score
}

func run10_2() {
	var input = read_lines("C:\\Users\\tothg\\Gege\\AOC2021\\res\\day_10.txt")

	scores := make([]int, 0)
	for _, line := range input {
		_, missingChars := getCharsEnding(line)
		fmt.Printf("Missing chars: %+v\n", missingChars)
		if len(missingChars) != 0 {
			score := autoCompleteScore(missingChars)
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	fmt.Printf("List of scores.: %+v\n", scores)
	fmt.Printf("List of scores.: %+v\n", scores[len(scores)/2])
}
