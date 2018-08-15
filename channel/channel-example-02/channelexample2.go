package main

import (
	"fmt"
)

func sum(numbers []int, channel chan int) {
	result := 0
	for _, number := range numbers {
		result = result + number
	}

	channel <- result
}

func main() {
	values := []int{7, 2, 8, -9, 4, 0}

	sumchannel := make(chan int)

	go sum(values[:len(values)/2], sumchannel)
	go sum(values[len(values)/2:], sumchannel)

	sum1, sum2 := <-sumchannel, <-sumchannel

	fmt.Println(sum1)
	fmt.Println(sum2)
	fmt.Println(sum1 + sum2)
}
