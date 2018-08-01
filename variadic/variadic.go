package main

import "fmt"

func printnumbers(numbers ...int) {
	fmt.Println(numbers)
}

func main() {
	printnumbers(1, 2)
	printnumbers(1, 2, 3)
	aslice := []int{1, 2, 3, 4, 5}
	printnumbers(aslice...)
	printnumbers()
}
