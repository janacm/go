package main

import "fmt"

/*
Go syntax
*/
func main() {
	fmt.Println("-- Begin --\n")
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
	//maps()
	//goroutines()
	interfaces()
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
	fmt.Printf("array: %s \n", array) // array: [smile CHANGED size spies fries]
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

// Arrays are fixed size. Slices are dynamically sized.
func arrays() {
	var a [10]string
	a[0] = "farts"
	a[1] = "toots"
	a[2] = "trumpet"
	a[3] = "cut the cheese"

	// Array literal
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

type GTD struct {
	nextAction string
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
pointToB = 2; // ERROR: Can't do this since pointers only hold memory addresses, not values
However, if you access the pointer using *, then you can:
*pointToB = 2; // Sets B to 2

*/
func goPointers() {
	fmt.Println("Initialize var a = 1")
	var a int = 1
	var pointToA *int
	pointToA = &a
	//pointToA = 2; // Can't do this since pointers only hold memory addresses, not values
	fmt.Printf("a: %v\n", a)
	fmt.Printf("&a: %v\n", &a)
	fmt.Printf("pointToA: %v\n", pointToA)

	// pointToA contains a memory address. *pointToA lets us access the variable pointed by that address.
	// This is called dereferencing.
	// Think of it like a pointer contains a reference, and you dereference it to get the value.
	*pointToA = 21
	fmt.Printf("a after setting pointToA=21: %v", a)

	// Dereferencing a nil value causes a panic:
	//var nilPointer *GTD = nil
	//fmt.Println(*nilPointer) // panic: runtime error

	// Struct Pointers
	fmt.Println()
	fmt.Println()
	fmt.Printf("---Struct Pointers---")
	fmt.Println()
	var gtd = GTD{
		"Learn pointers",
	}
	pointToGtd := &gtd
	fmt.Printf("gtd: %v\n", gtd)
	fmt.Printf("pointToGtd: %v\n", pointToGtd)
	fmt.Printf("pointToGtd.nextAction: %v\n", pointToGtd.nextAction) // Notice we don't have to add the * before pointToGtd in order to access the value
	fmt.Printf("*pointToGtd: %v\n", *pointToGtd)
	fmt.Printf("*pointToGtd.nextAction: %v\n", "invalid indirect of pointToGtd.nextAction\n")

	fmt.Printf("---Nil Checks Pointers---\n")
	// Nil checks
	if &gtd.nextAction != nil {
		fmt.Printf("gtd: %v\n", gtd)
	}

	fmt.Println("Dereferencing Nil vars")

	pointToGtd.setAction("Learn pointer receivers")
	fmt.Printf("pointToGtd.nextAction after update: %v\n", pointToGtd.nextAction) // Incorrect: old value remains!
	pointToGtd.setAction_pointerReceiver("Set value using pointer receivers")
	fmt.Printf("pointToGtd.nextAction after update: %v\n", pointToGtd.nextAction) // Correct: New value

	setAction_noPointers(gtd, "Try to set value without pointers")
	fmt.Printf("pointToGtd.nextAction after update: %v\n", pointToGtd.nextAction) // Incorrect: old value remains!

	setAction_pointerArg(&gtd, "Set value using pointer arg")
	fmt.Printf("pointToGtd.nextAction after update: %v\n", pointToGtd.nextAction) // Incorrect: old value remains!

}

func (g GTD) setAction(action string) {
	g.nextAction = action
}

func (g *GTD) setAction_pointerReceiver(action string) {
	g.nextAction = action
}

func setAction_noPointers(g GTD, action string) {
	g.nextAction = action
}

func setAction_pointerArg(g *GTD, action string) {
	g.nextAction = action
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
	fmt.Printf("%T, %v, %v", a, a, a)
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
