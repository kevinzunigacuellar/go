package iterations

import "strings"

func Repeat(char string, times int) (repeated string) {
	for i := 0; i < times; i++ {
		repeated += char
	}
	return repeated
}

func RepeatStrings(char string, times int) string {
	return strings.Repeat(char, times)
}
