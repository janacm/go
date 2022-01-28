package main

import "fmt"

// Arrays are fixed size. Slices are dynamically sized. Slices are more common in Go.
func arrays() {
	var a [10]string
	a[0] = "farts"
	a[1] = "toots"
	a[2] = "trumpet"
	a[3] = "cut the cheese"

	// Array literal
	b := [2]string{"smile", "smize"}

	// Array of struct literals
	structs := [2]struct {
		name string
		age int
	}{
		{"pradeega", 25},
		{"jeff", 30},
	}

	// Array of structs
	gtd_step2 := GTD{ // A struct
		nextAction: "define array type",
	}
	gtds := [3]GTD{
		GTD{ // Don't need to specify type
			nextAction: "define size",
		},
		gtd_step2, // Can use a var
		{ // Can use a literal
			nextAction: "supply values",
		},
	}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%v \n", structs)
	fmt.Printf("%v \n", gtds)
}

// Slices are dynamically sized collections which are built off of arrays
func slices() {
	// Define literal array
	array := [5]string{"smile", "smize", "size", "spies", "fries"}
	fmt.Printf("array: %s \n", array) // array: [smile CHANGED size spies fries]
	fmt.Printf("array: %s \n", array[:]) // Defaults to 0:len, aka no change

	slice := array[1:3] // [start:end], [1,3) - 3rd element is not included
	fmt.Printf("slice: %s \n", slice) // slice: [smize, size]
	slice[0] = "CHANGED"
	fmt.Printf("slice: %s \n", slice) // slice: [CHANGED size]
	fullSlice := array[:]
	fmt.Println("whole array slice: ", fullSlice)

	fmt.Printf("\n---Capacity VS Length---")
	fmt.Printf("Slice's capacity is the underlying arrays length, counting from the first element in the slice: %v \n", cap(slice))
	fmt.Printf("Slice's length is how many elements it contains: %v \n", len(slice))
	fmt.Printf("Underlying arrays length: %v \n", len(array))

	// We can extend a slice up till its capacity
	// slice = slice[1:] // Remove first element
	fmt.Printf("Slice literal")
	fmt.Printf("Slice literals are defined the same as array literals except without a size in []")
	a := []int{1, 2, 3} // Slice literal
	fmt.Printf("a = %v", a)

	//a = a[0:4] // Can't extend slices beyond capacity
	fmt.Printf("a = %v \n", a)

	// Can use `make` to create a slice with a specified len/cap
	makedSlice := make([]int, 5)
	printSlice(makedSlice, "makedSlice appended")

	// Can append slices
	makedSlice = append(makedSlice, 100, 101, 102) // slice grows beyond its capacity when appended
	printSlice(makedSlice, "makedSlice")

	// 2d slice
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	fmt.Println(board)

}
func printSlice(a []int, label string) {
	fmt.Printf("%v = %v, len=%v, cap=%v \n", label, a, len(a), cap(a))
}
