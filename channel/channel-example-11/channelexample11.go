package main

import (
	"fmt"
)

func main() {
	queue := make(chan string, 2)

	queue <- "This is value one in the queue."
	queue <- "This is value two in the queue."

	close(queue)

	for element := range queue {
		fmt.Println(element)
	}
}
