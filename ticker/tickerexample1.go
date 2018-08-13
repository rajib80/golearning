package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("Ticked at", t)
		}
	}()

	time.Sleep(2600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
