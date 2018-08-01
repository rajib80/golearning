package main

import (
	"fmt"
)

func filter(numbers []int, callback func(int) bool) []int {
	var result []int

	for _, number := range numbers {
		if callback(number) {
			result = append(result, number)
		}
	}

	return result
}

func main() {
	var filternumbers = filter([]int{1, 2, 3, 4}, func(number int) bool {
		return number > 1
	})

	fmt.Println(filternumbers)
}
