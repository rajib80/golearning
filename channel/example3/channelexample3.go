package main

import "fmt"

func main() {
	aBufferedChannel := make(chan string, 2)

	aBufferedChannel <- "buffered"
	aBufferedChannel <- "channel"

	fmt.Println(<-aBufferedChannel)
	fmt.Println(<-aBufferedChannel)
}
