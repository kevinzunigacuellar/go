package arrays

func Sum(numbers []int) int {
	var result int
	for _, num := range numbers {
		result += num
	}
	return result
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(arrs ...[]int) []int {
	var sums []int
	for _, arr := range arrs {
		if len(arr) == 0 {
			sums = append(sums, 0)
			continue
		}
		tails := arr[1:]
		sums = append(sums, Sum(tails))
	}
	return sums
}
