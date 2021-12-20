package main

import "fmt"

/**
Interfaces are a set of methods that a type must implement.
There is no explicit "implements" keyword. A type implements an interface by
defining a method with the same name and signature.
*/

type Comparable interface {
    Compare(Comparable) bool
}

// Defining this method means that GTD implements Comparable
func (g GTD) Compare(comparable Comparable) bool {
	//return g.nextAction == comparable.(GTD).nextAction
	other := comparable.(GTD) // Type assertion: comparable is a GTD
	return g.nextAction == other.nextAction
}

func interfaces() {
    var task1 Comparable
    task1 = GTD{
		"create an interface",
	}
	task2 := GTD{
		"implement interface",
	}
	fmt.Println(task1.Compare(task2))
}

func theEmptyInterface() {
	var any interface{}
	any = "The empty interface can hold values of any type"
	printIt(any)

	any = 3
	printIt(any)

	any = true
	printIt(any)
}
func printIt(i interface{}) {
    fmt.Println(i)
}