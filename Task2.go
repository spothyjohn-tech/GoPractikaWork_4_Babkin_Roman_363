package main

import (
	"fmt"
	"sync"
)

func worker(wg *sync.WaitGroup, jobs <-chan int, results chan<- int) {
	defer wg.Done()
	for num := range jobs {
		results <- num * num
	}
}

func main() {
	const numJobs = 10
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	var wg sync.WaitGroup
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(&wg, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()
	for result := range results {
		fmt.Println(result)
	}
}
