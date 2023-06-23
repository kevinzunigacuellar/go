package main

import (
	"fmt"
	"os"
	"strings"
)

func Hello() string {
	return "Hello, world"
}

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
