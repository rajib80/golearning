package main

import "fmt"
import "sort"

func main() {
	strings := []string{"b", "x", "u"}
	sort.Strings(strings)
	fmt.Println(strings)

	integers := []int{3, 0, -1, 5, 9}
	sort.Ints(integers)
	fmt.Println(integers)

	isIntegersSorted := sort.IntsAreSorted(integers)
	fmt.Println("Is the integers sorted:", isIntegersSorted)
}
