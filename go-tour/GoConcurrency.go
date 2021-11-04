package main

import (
	"fmt"
	"time"
)

func say(s string) {
	fmt.Printf("%s\n", s)
}

func sayPeriodically(s string, times int, seconds time.Duration) {
	for i := 0; i < times; i++ {
		time.Sleep(seconds * time.Second)
		say(s)
	}
}

func execFxn(fxn func(string), s string) {
	fxn(s)
}

func printInts(prefix string, numOfTimes int) {
	for i := 0; i < numOfTimes; i++ {
		fmt.Printf("%s %d\n", prefix, i)
	}
}

func printIntsWithChannel(prefix string, numOfTimes int, c chan int) {
	for i := 0; i < numOfTimes; i++ {
		fmt.Printf("PIWC: %s %d\n", prefix, i)
		c <- i
	}
}

func goroutines() {
	fmt.Println("-goroutines start-")

	// Goroutines are lightweight threads that can be used to run multiple functions concurrently
	// Note that these only produce output if there is a delay in the execution of the subsequent functions
	go printInts("Goroutine 1", 100)
	go printInts("Goroutine 2", 100)
	channels()

	fmt.Println("goroutines end")
}

// Channels are used to communicate between goroutines, and let you send and receive values from multiple goroutines.
// Channels are typed by the values they carry.
func channels() {
	c := make(chan int) // make a channel of ints.
	go printIntsWithChannel("Channel 1", 100, c)
	go printIntsWithChannel("Channel 2", 100, c)
	x, y := <-c
	fmt.Println("x:", x, "y:", y)
}

// FYI: You didn't figure out how this works yet.
