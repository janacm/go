package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile_byDay(day int) []byte {
	fmt.Printf("Reading: %v\n", day)
	var sb strings.Builder
	sb.WriteString("day")
	sb.WriteString(strconv.Itoa(day))
	sb.WriteString(".txt")
	filename := sb.String() // "day1.txt"

	file, err := os.ReadFile(filename)
	//file, err := os.Open(filename)
	check(err)
	return file
}

func readInputFile_asStringSlice(day int) []string {
	var sb strings.Builder
	sb.WriteString("day")
	sb.WriteString(strconv.Itoa(day))
	sb.WriteString(".txt")
	filename := sb.String() // "day1.txt"
	fmt.Printf("Reading: %v\n", filename)

	file, err := os.Open(filename)
	scanner := bufio.NewScanner(file)
	var slice []string
	for scanner.Scan() {
		slice = append(slice, scanner.Text())
	}
	check(err)
	fmt.Printf("Slice's length is how many elements it contains: %v \n", len(slice))
	return slice
}

func readInputFile_asIntegerSlice(day int) []int {
	var sb strings.Builder
	sb.WriteString("day")
	sb.WriteString(strconv.Itoa(day))
	sb.WriteString(".txt")
	filename := sb.String() // "day1.txt"
	fmt.Printf("Reading: %v\n", filename)

	file, err := os.Open(filename)
	scanner := bufio.NewScanner(file)
	var slice []int
	for scanner.Scan() {
		intInputLine, err := strconv.Atoi(scanner.Text())
		check(err)
		slice = append(slice, intInputLine)
	}
	check(err)
	fmt.Printf("Slice's length is how many elements it contains: %v \n", len(slice))
	return slice
}

func check(err error) {
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
}
