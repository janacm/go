package main

import "fmt"

func main() {
	//fmt.Println(add(3,4))
	//fmt.Println(getTwo())
	//fmt.Println(nakedReturn())
	//fmt.Println()
	whatsMyType()


}

func add(x, y int) int {
	return x + y
}

func getTwo()(a string, b string) {
	return "one", "two"
}
func nakedReturn()(a string, b string) {
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
	a := 30;
	fmt.Printf("%T, %v, %a", a, a, a)
}
