package main

import (
	"fmt"
	"strconv"
)

func Day1() {
	fmt.Println("Day 1")
	fileSlice := readInputFile_asStringSlice(1)
	count := 0
	for i := 0; i < len(fileSlice)-1; i++ {
		fmt.Println(fileSlice[i])
		currentInt, err := strconv.Atoi(fileSlice[i])
		check(err)
		nextInt, err2 := strconv.Atoi(fileSlice[i+1])
		check(err2)
		if currentInt < nextInt {
			count++
		}
	}
	fmt.Printf("count: %d\n", count)
}

func Day1_pt2() {
	fmt.Println("Day 1")
	fileSlice := readInputFile_asIntegerSlice(1)
	count := 0
	currentWindowSum := 0
	nextWindowSum := 0
	for i := 0; i < len(fileSlice)-3; i++ {
		currentWindowSum = add3(i, fileSlice)
		nextWindowSum = add3(i+1, fileSlice)

		if currentWindowSum < nextWindowSum {
			count++
		}

	}
	fmt.Printf("count: %d\n", count)
}

func add3(start int, fileSlice []int) int {
	sum := 0
	for j := 0; j < 3; j++ {
		sum += fileSlice[start+j]
	}
	return sum
}
