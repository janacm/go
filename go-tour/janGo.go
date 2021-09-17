package main

import "fmt"

/*
Go syntax
*/
func main() {
	//fmt.Println(add(3,4))
	//fmt.Println(getTwo())
	//fmt.Println(nakedReturn())
	//fmt.Println()
	//whatsMyType()
	//shortIf();
	//deferFxn()
	//goPointers()
	//structs()
	//arrays()
	//slices()
	//ranges()
	maps()

}

func maps() {
	var m map[string]string
	m = make(map[string]string)
	m["name"] = "pradeega"

	m2 := map[string]string{
		"this": "that",
	}
	fmt.Println(m)
	fmt.Println(m2)
}

func ranges() {
	pow := []int{1, 2, 4, 8, 16, 32}
	for i, v := range pow {
		fmt.Printf("2^%v = %v \n ", i, v)
	}
	fmt.Println()

	for i, _ := range pow {
		fmt.Printf("2^%v \n", i)
	}
	fmt.Println()

	for i := range pow {
		fmt.Printf("2^%v  \n", i)
	}
	fmt.Println()

	for _, v := range pow {
		fmt.Printf("2^ = %v \n ", v)
	}
}

// Slices are dynamically sized collections which are built off of arrays
func slices() {
	array := [5]string{"smile", "smize", "size", "spies", "fries"}
	slice := array[1:3]
	slice[0] = "CHANGED"
	fmt.Printf("array: %s \n", array) // Changing the slice, updates the underlying array
	fmt.Printf("slice: %s \n", slice)
	fmt.Println()

	// Capacity VS Length
	fmt.Printf("Slice's capacity is the underlying arrays length, counting from the first element in the slice: %v \n", cap(slice))
	fmt.Printf("Slice's length is how many elements it contains: %v \n", len(slice))
	fmt.Printf("Underlying arrays length: %v \n", len(array))

	// We can extend a slice up till its capacity, before it
	// slice = slice[1:] // Remove first element
	a := []int{1, 2, 3} // Slice literal
	fmt.Printf("a = %v", a)

	//a = a[0:4] // Can't extend slices beyond capacity
	fmt.Printf("a = %v \n", a)

	// Can use `make` to create a slice with a specified len/cap
	m := make([]int, 5)
	printSlice(m, "m appended")

	// Can append slices
	m = append(m, 100, 101, 102) // slice grows beyond its capacity when appended
	printSlice(m, "m")

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

// Arrays are fixed size. Slices are dynamically sized.
func arrays() {
	var a [10]string
	a[0] = "farts"
	a[1] = "toots"
	a[2] = "trumpet"
	a[3] = "cut the cheese"

	b := [2]string{"smile", "smize"}

	fmt.Println(a)
	fmt.Println(b)
}

// structs
func structs() {
	type Person struct {
		name  string
		email string
	}

	var (
		janac    = Person{"janac", "janac.meena@gmail.com"}
		pradeega = Person{"pradeega", "p@p.com"}
	)
	fmt.Println(Person{"jan", "bo@bo.com"})
	fmt.Println(janac.name)
	fmt.Println(pradeega.email)

}

/**
Go has pointers. Pointers can only hold
the memory address of a value. You define
a pointer by adding * before the type:

var a int // normal int
var a *int // a pointer to an int

We initialize the pointer by generating a memory address
from an existing variable using &
pointToB := &b

*/
func goPointers() {
	var b int = 2
	//var pointToB *int
	pointToB := &b
	fmt.Println(b)
	fmt.Println(&b)
	fmt.Println(pointToB)

	*pointToB = 21 // pointToB contains a memory address. *pointToB lets us access the variable related to that address.
	fmt.Println(b) // 21

}

func add(x, y int) int {
	return x + y
}

func getTwo() (a string, b string) {
	return "one", "two"
}
func nakedReturn() (a string, b string) {
	a = "hi"
	b = "HOLA AMIGOS"
	return
}

func someVars() int {
	a := 3
	var b = 3
	return a + b
}

func whatsMyType() {
	a := 30
	fmt.Printf("%T, %v, %a", a, a, a)
}

func shortIf() {
	b := 20
	if a := 30; a < b {
		fmt.Println("a < b")
	} else {
		fmt.Println("a > b")
	}
}

func deferFxn() {
	fmt.Println("hi")
	defer fmt.Println("Printed last")
	fmt.Println("bye")
}
