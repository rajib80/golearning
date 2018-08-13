package main

import (
	"fmt"
)

func fibonacci(number int, channel chan int) {
	x, y := 0, 1
	for i := 0; i < number; i++ {
		channel <- x
		x, y = y, x+y
	}
	close(channel)
}

func main() {
	channel := make(chan int)
	go fibonacci(15, channel)
	for i := range channel {
		fmt.Println(i)
	}
}
