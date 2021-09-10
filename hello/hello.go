package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

func main() {
	// Get a greeting message and print it.
	names := []string{"Pradeega", "Pragoobie", "Bagoodie"}
	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
