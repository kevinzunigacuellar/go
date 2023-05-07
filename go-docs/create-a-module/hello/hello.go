package main

import (
	"create-a-module/greetings"
	"fmt"
	"log"
)

func main() {
	// Get a greeting message and print it.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := greetings.Hello("Sam")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
