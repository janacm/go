package main

import (
	"fmt"
	"strconv"
	"strings"
)

//goland:noinspection GoUnusedExportedFunction
/**
Question:
	How many zeroes in a column of binary numbers?

Strategy
- How to get first bit of binary number?
	- Check if number is even or odd
	- Least significant bit (LSB), which is the rightmost bit of the number
	- If even, then LSB is 0
	- Store in array
	- Do a right shift and repeat until an array of total counts of zero

1 1 1 0 = 14
8 4 2 1
1 1 1 1 = 15
- Just invert the gamma rate binary string to get epsilon rate
*/
const numLength = 5

func Day3() {
	fmt.Println("Day 3")
	fileSlice := readInputFile_asStringSlice(3)
	var maxPossibleCount = len(fileSlice) / 2

	//columnTotalZeros := [numLength]int{0,0,0,0,0,0,0,0,0,0,0,0}
	var columnTotalZeros []int

	//for i, decNum := range fileSlice {
	//	fileSlice[i] = decNumz
	//}
	for _, decNum := range fileSlice {
		shiftedNum, _ := strconv.ParseInt(decNum, 2, 64)
		for j := numLength - 1; j >= 0; j-- {
			fmt.Printf("j: %d shiftedNum: %d , %b\n", j, shiftedNum, shiftedNum)
			if isEven(shiftedNum) {
				columnTotalZeros[j]++
			}
			shiftedNum = shiftedNum >> 1
		}
	}
	fmt.Printf("Column totals: %v\n", columnTotalZeros)

	gammaRateBinary := [numLength]int{}
	epsilonRateBinary := [numLength]int{}

	for i := 0; i < numLength; i++ {
		if columnTotalZeros[i] < maxPossibleCount {
			gammaRateBinary[i] = 1
		} else {
			epsilonRateBinary[i] = 1
		}
	}

	fmt.Printf("gammaRateBinary: %b\n", gammaRateBinary)
	fmt.Printf("epsilonRateBinary: %b\n", epsilonRateBinary)

	stringedGammaRate := intArrayToString(gammaRateBinary)
	stringedEpisilonRate := intArrayToString(epsilonRateBinary)
	fmt.Printf("stringedGammaRate: %s\n", stringedGammaRate)
	fmt.Printf("stringedEpisilonRate: %s\n", stringedEpisilonRate)

	gammaRate, _ := strconv.ParseInt(stringedGammaRate, 2, 64)
	epsilonRate, _ := strconv.ParseInt(stringedEpisilonRate, 2, 64)
	fmt.Printf("gammaRate: %d\n", gammaRate)
	fmt.Printf("epsilonRate: %d\n", epsilonRate)

	fmt.Printf("Ans in decimal is gamma * epsilon: %d", gammaRate*epsilonRate)
}

func Day3_pt2() {
	fmt.Println("Day 3")
	fileSlice := readInputFile_asStringSlice(3)
	var maxPossibleCount = len(fileSlice) / 2
	fmt.Printf("maxPossibleCount: %d\n", maxPossibleCount)
	columnTotalZeros := [numLength]int{}
	//oxygenRating := 0
	//c02Rating := 0

	for _, decNum := range fileSlice {
		shiftedNum, _ := strconv.ParseInt(decNum, 2, 64)
		for j := numLength - 1; j >= 0; j-- {
			fmt.Printf("j: %d shiftedNum: %d , %b\n", j, shiftedNum, shiftedNum)
			if isEven(shiftedNum) {
				columnTotalZeros[j]++
			}
			shiftedNum = shiftedNum >> 1
		}
	}
	fmt.Printf("Column totals: %v\n", columnTotalZeros)

	gammaRateBinary := [numLength]int{}
	epsilonRateBinary := [numLength]int{}

	for i := 0; i < numLength; i++ {
		if columnTotalZeros[i] < maxPossibleCount {
			gammaRateBinary[i] = 1
		} else if columnTotalZeros[i] == maxPossibleCount {
			gammaRateBinary[i] = 1
		} else {
			epsilonRateBinary[i] = 1
		}
	}

	fmt.Printf("gammaRateBinary: %b\n", gammaRateBinary)
	fmt.Printf("epsilonRateBinary: %b\n", epsilonRateBinary)

	stringedGammaRate := intArrayToString(gammaRateBinary)
	stringedEpisilonRate := intArrayToString(epsilonRateBinary)
	fmt.Printf("stringedGammaRate: %s\n", stringedGammaRate)
	fmt.Printf("stringedEpisilonRate: %s\n", stringedEpisilonRate)

	gammaRate, _ := strconv.ParseInt(stringedGammaRate, 2, 64)
	epsilonRate, _ := strconv.ParseInt(stringedEpisilonRate, 2, 64)
	fmt.Printf("gammaRate: %d\n", gammaRate)
	fmt.Printf("epsilonRate: %d\n", epsilonRate)

	fmt.Printf("Ans in decimal is gamma * epsilon: %d", gammaRate*epsilonRate)
}
func isEven(num int64) bool {
	return num%2 == 0
}

func intArrayToString(arr [numLength]int) string {
	var stringArr []string
	for _, num := range arr {
		stringArr = append(stringArr, strconv.Itoa(num))
	}
	return strings.Join(stringArr, "")
}
