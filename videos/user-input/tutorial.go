package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Type the year you were born: ")
	scanner.Scan()
	input, err := strconv.ParseInt(scanner.Text(), 10, 64)
	bornYear := int(input)
	if err != nil {
		log.SetPrefix("ERROR: ")
		log.SetFlags(0)
		log.Fatal("Please type a valid year")
	}

	currentYear := time.Now().Year()
	fmt.Printf("This year you will be %v \n", currentYear-bornYear)
}
