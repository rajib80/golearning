package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var counter uint64

	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&counter, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	finalCounter := atomic.LoadUint64(&counter)
	fmt.Println("Counter:", finalCounter)
}
