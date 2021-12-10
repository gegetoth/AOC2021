package main

import (
	"bufio"
	"log"
	"os"
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

func printIntMatrix(intMatrix [][]int) {
	for i := 0; i < len(intMatrix); i++ {
		row := intMatrix[i]
		for j := 0; j < len(row); j++ {
			print(row[j])
		}
		println("")
	}
}

func printStrMatrix(intMatrix [][]string) {
	for i := 0; i < len(intMatrix); i++ {
		row := intMatrix[i]
		for j := 0; j < len(row); j++ {
			print(row[j])
		}
		println("")
	}
	println("")
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
