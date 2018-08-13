package main

import "fmt"

func afunction(inputstring string) {
	for i := 0; i < 5; i++ {
		fmt.Println(inputstring, ":", i)
	}
}

func main() {
	afunction("direct call")

	go afunction("goroutine 1")
	go afunction("goroutine 2")
	go afunction("goroutine 3")
	go afunction("goroutine 4")

	go func(msg string) {
		fmt.Println(msg)
	}("this is another function call from goroutine")

	fmt.Scanln() //This Scanln requires we press a key before the program exits.
	fmt.Println("done")
}
