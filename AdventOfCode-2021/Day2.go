package main

import (
	"fmt"
	"strconv"
	"strings"
)

//goland:noinspection GoUnusedExportedFunction
func Day2() {
	fmt.Println("Day 2")
	fileSlice := readInputFile_asStringSlice(2)
	depth := 0
	horPos := 0
	for i := 0; i < len(fileSlice); i++ {
		fmt.Println(fileSlice[i])
		dir, amount := parseDirection(fileSlice[i])
		switch dir {
		case "forward":
			horPos += amount
		case "down":
			depth += amount
		case "up":
			depth -= amount
		}
	}
	fmt.Printf("Final position: %d, %d\n", horPos, depth)
	fmt.Printf("Ans: %d", horPos*depth)
}

func parseDirection(direction string) (string, int) {
	split := strings.Split(direction, " ")
	amount, err := strconv.Atoi(split[1])
	Check(err)
	return split[0], amount
}

func Day2_pt2() {
	fmt.Println("Day 2 pt2")
	fileSlice := readInputFile_asStringSlice(2)
	depth := 0
	horPos := 0
	aim := 0
	for i := 0; i < len(fileSlice); i++ {
		fmt.Println(fileSlice[i])
		dir, amount := parseDirection(fileSlice[i])
		switch dir {
		case "forward":
			horPos += amount
			depth += aim * amount
		case "down":
			aim += amount
		case "up":
			aim -= amount
		}
	}
	fmt.Printf("Final position: %d, %d\n", horPos, depth)
	fmt.Printf("Ans: %d", horPos*depth)
}
