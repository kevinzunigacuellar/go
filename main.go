package main

import "fmt"

/*
	version 1
	func main() {
		var s, sep string // declare an empty string
		// this for loop takes a initial value, a condition, and a post statement

		intial value: i := 1
		condition: i < len(os.Args)
		post statement: i++

		for i := 1; i < len(os.Args); i++ {
			// s is created in each iteration of the loop
			s += sep + os.Args[i]
			sep = " "
		}
		fmt.Println(s)
	}

	version 2
	func main() {
		var s, sep string // declare an empty string
		// this for loop takes a range expression
		for _, arg := range os.Args[1:] {
			// s is recreated in each iteration of the loop
			s += sep + arg
			sep = " "
		}
		fmt.Println(s)
	}

	version 3
	func main() {
		fmt.Println(strings.Join(os.Args[1:], " "))
	}



// Exercise 1.1: Modify the echo program to also print os.Args[0], the name of the command that invoked it.

func main() {
	fmt.Println(strings.Join(os.Args[:], " "))
}



// Exercise 1.2: Modify the echo program to print the index and value of each of its arguments, one per line.

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
}

// TODO: Exercise 1.3 (page 8)

*/

// func main() {
// 	counts := make(map[string]int)
// 	input := bufio.NewScanner(os.Stdin)
// 	for input.Scan() {
// 		counts[input.Text()]++
// 	}

// 	for line, n := range counts {
// 		if n > 1 {
// 			fmt.Printf("%d\t%s\n", n, line)
// 		}
// 	}
// }

func main() {
	fruits := []string{}
	fruits = append(fruits, "apple")
	fruits = append(fruits, "banana")
	fruits = append(fruits, "orange")
	// fmt.Println(fruits)

	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}

}
