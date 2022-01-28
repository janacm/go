package main

import (
	"fmt"
	"time"
)


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
	//interfaces()
	//bits()
	dateAndtime()
}

func dateAndtime() {
	now := time.Now()
	fmt.Println(now.String()) // 2022-01-19 16:55:44.443302 -0500 EST m=+0.000143126

	fmt.Printf("time.Kitchen: %v\n", time.Kitchen)
	kitchenTime, _ := time.Parse(time.Kitchen, "4:20PM") // time.Parse returns a time.Time, from a layout and a string that fits in that layout
	fmt.Printf("kitchenTime: %v\n", kitchenTime.String()) //kitchenTime: 0000-01-01 16:20:00 +0000 UTC, year/m/d is 0 since unspecified values default to zero

	time.Parse(time.RFC822Z, now.String())
	fmt.Printf("time.RFC822Z: %v\n", now.String()) // time.RFC822Z: 2022-01-19 21:04:12.728706 -0500 EST m=+0.000367251

	fmt.Printf("now.Unix(): %v\n", now.Unix()) // now.Unix(): 1579793344
	fmt.Printf("now.UTC(): %v\n", now.UTC()) // now.UTF(): 2022-01-19 16:55:44.443302 +0000 UTC
	fmt.Printf("now + 1hr: %v\n", now.Add(time.Hour * 1)) // now.UTF(): 2022-01-19 16:55:44.443302 +0000 UTC
	fmt.Printf("now + 1 day: %v\n", now.Add(time.Minute * 1440)) // 1440 mins in a day

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
