package main

import (
	"create-a-module/greetings"
	"fmt"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Sam")
	fmt.Println(message)
}
