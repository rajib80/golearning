package main

import "fmt"

func main() {
	var aninteger = 100
	var pointertoaninteger *int

	pointertoaninteger = &aninteger

	fmt.Println(aninteger)
	fmt.Println(&aninteger)
	fmt.Println(pointertoaninteger)
	fmt.Println(*pointertoaninteger)
}
