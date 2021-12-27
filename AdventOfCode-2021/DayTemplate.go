package main

import (
	"fmt"
)

//goland:noinspection GoUnusedExportedFunction
func Day0() {
	fmt.Println("Day 0")
	fileSlice := readInputFile_asStringSlice(0)
	fmt.Printf("%v\n", fileSlice)
	fmt.Printf("Ans: %d", 0)
}
