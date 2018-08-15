package main

import (
	"fmt"
)

func main() {
	message := make(chan string)

	go func() {
		message <- "ping"
	}()

	//fmt.Println(message)
	//fmt.Println(<-message)

	msg := <-message
	fmt.Println(msg)
}
