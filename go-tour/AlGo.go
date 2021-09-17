package main

import (
	"fmt"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	return map[string]int{"x": 1}
}

func main() {
	ans := WordCount("a a a b b c")
	if ans["a"] == 3 && ans["b"] == 2 && ans["c"] == 1 {
		fmt.Println("Works")
	} else {
		fmt.Println("You fucked up")
	}
}
