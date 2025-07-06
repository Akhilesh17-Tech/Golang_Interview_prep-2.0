package main

import (
	"fmt"
	"sync"
	"time"
)

// A worker pool is a concurrency pattern that allows you to limit the number of goroutines (workers) that handle a certain type of task. It helps in managing resource utilization effectively by controlling the number of concurrent tasks running.
// The worker pool pattern is useful when you have a large number of tasks that need to be processed concurrently, but you don't want to create too many goroutines at once. This can help prevent resource exhaustion and improve performance.

// jobs <-chan int Purpose: This channel is used to receive jobs that need to be processed by the worker.
// results chan<- int Purpose: This channel is used to send results produced by the worker after processing a job.

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(time.Second) // Simulating work
		fmt.Printf("Worker %d finished job %d\n", id, j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	var wg sync.WaitGroup

	// Starting workers
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Sending jobs to the workers
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Waiting for all workers to finish
	wg.Wait()
	close(results)

	// Collecting results
	for result := range results {
		fmt.Println("Result:", result)
	}
}
