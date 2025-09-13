package main

import (
	"fmt"
	"sync"
)

func worker(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range ch {
		fmt.Printf("Рабочий %d получил задачу %d\n", id, task)
	}
}

func main() {
	var wg sync.WaitGroup

	const workerCount = 3
	tasks := []int{1, 2, 3, 4, 5, 6, 7}

	workerChans := make([]chan int, workerCount)

	for i := 0; i < workerCount; i++ {
		workerChans[i] = make(chan int)
		wg.Add(1)
		go worker(i+1, workerChans[i], &wg)
	}

	for i, task := range tasks {
		target := i % workerCount
		workerChans[target] <- task
	}

	for i := 0; i < workerCount; i++ {
		close(workerChans[i])
	}
	wg.Wait()
}
