package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Started working...")
	time.Sleep(5000 * time.Millisecond)
	fmt.Println("Done")

	done <- true
}

func main() {
	channel := make(chan bool, 1)
	go worker(channel)

	<-channel
}
