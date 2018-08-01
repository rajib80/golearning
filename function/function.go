package main

import "fmt"

func main() {
	fmt.Println("The result of adding 99 and 57 is", add(99, 57))
	var result, remainder = divide(99, 37)
	fmt.Println("The result of 99 divide by 37 is", result)
	fmt.Println("The remainder of 99 divide by 37 is", remainder)
}

func add(number1 int, number2 int) int {
	return number1 + number2
}

func divide(number1 int, number2 int) (int, int) {
	return number1 / number2, number1 % number2
}
