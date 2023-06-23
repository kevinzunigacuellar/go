package main

import "fmt"

func main() {
	var myarr [3]int
	myarr[0] = 1
	myarr[1] = 2
	myarr[2] = 3

	var myarr2 = [3]int{1, 2, 3}

	for _, v := range myarr2 {
		println(v)
	}

	// slice
	scores := make([]int, 0, 10)
	scores = append(scores, 5)
	scores = append(scores, 1)
	fmt.Println(scores)

	sc2 := make([]int, 0, 10)
	sc2 = sc2[0:8]
	sc2[7] = 10
	fmt.Println(sc2)
}
