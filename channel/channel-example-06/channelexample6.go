package main

import (
	"fmt"
)

func ping(pings chan<- string, message string) {
	pings <- message
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string)
	pongs := make(chan string)

	go ping(pings, "Message Passed")
	go pong(pings, pongs)

	fmt.Println(<-pongs)
}
