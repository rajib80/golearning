package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("Worker", id, "Started job", job)
		time.Sleep(2 * time.Second)
		fmt.Println("Worker", id, "Finished job", job)
		results <- job * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 5; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 100; j++ {
		jobs <- j
	}
	close(jobs)

	for r := 1; r <= 100; r++ {
		<-results
	}
}
