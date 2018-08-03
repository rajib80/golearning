package main

import "fmt"

func main() {
	evenNumbers := make([]int, 5, 10)

	value := 1
	for i := 0; i < len(evenNumbers); i++ {
		evenNumbers[i] = value * 2
		value++
	}

	fmt.Println(evenNumbers)
	fmt.Println(evenNumbers[2:4])
}
