package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func read_lines(location string) []string {

	// os.Open() opens specific file in
	// read-only mode and this return
	// a pointer of type os.
	file, err := os.Open(location)

	if err != nil {
		log.Fatalf("failed to open")

	}

	// The bufio.NewScanner() function is called in which the
	// object os.File passed as its parameter and this returns a
	// object bufio.Scanner which is further used on the
	// bufio.Scanner.Split() method.
	scanner := bufio.NewScanner(file)

	// The bufio.ScanLines is used as an
	// input to the method bufio.Scanner.Split()
	// and then the scanning forwards to each
	// new line using the bufio.Scanner.Scan()
	// method.
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	// The method os.File.Close() is called
	// on the os.File object to close the file
	file.Close()

	// and then a loop iterates through
	// and prints each of the slice values.
	//for _, each_ln := range text {
	//	fmt.Println(each_ln)
	//}
	return text
}

func getIntMatrix(input []string) [][]int {
	var intMatrix [][]int
	for _, each_ln := range input {
		chars := []rune(each_ln)
		intArr := make([]int, len(chars))
		for i := 0; i < len(chars); i++ {
			value, _ := strconv.Atoi(string(chars[i]))
			intArr[i] = value
		}
		intMatrix = append(intMatrix, intArr)
	}

	printIntMatrix(intMatrix)

	return intMatrix
}

func getMatrix(input []string) [][]string {
	var intMatrix [][]string
	for _, each_ln := range input {
		chars := []rune(each_ln)
		intArr := make([]string, len(chars))
		for i := 0; i < len(chars); i++ {
			intArr[i] = string(chars[i])
		}
		intMatrix = append(intMatrix, intArr)
	}

	printStrMatrix(intMatrix)

	return intMatrix
}

func printIntMatrix(intMatrix [][]int) {
	for i := 0; i < len(intMatrix); i++ {
		row := intMatrix[i]
		for j := 0; j < len(row); j++ {
			fmt.Printf("%+v", row[j])
		}
		fmt.Printf("\n")
	}
}

func printStrMatrix(intMatrix [][]string) {
	for i := 0; i < len(intMatrix); i++ {
		row := intMatrix[i]
		for j := 0; j < len(row); j++ {
			fmt.Printf("%+v", row[j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

type Stack []interface{}

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

type Queue []interface{}

// IsEmpty: check if stack is empty
func (s *Queue) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Queue) Push(str interface{}) {
	*s = append(*s, str) // Simply append the new value to the end of the queue
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Queue) Get() (interface{}, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		element := (*s)[0] // Index into the slice and obtain the element.
		*s = (*s)[1:]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
