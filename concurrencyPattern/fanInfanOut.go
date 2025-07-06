// Fan-Out Pattern
// In the fan-out pattern, you have a single source of work that is distributed to multiple goroutines to be processed concurrently.

// Fan-In Pattern
// The fan-in pattern is where you collect the results of multiple goroutines and combine them into a single result or stream.

package main

import (
	"fmt"
	"sync"
)

func producer(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c)
}

func consumer(id int, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range c {
		fmt.Printf("Consumer %d received %d\n", id, n)
	}
}

func main() {
	c := make(chan int)
	var wg sync.WaitGroup

	go producer(c)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go consumer(i, c, &wg)
	}

	wg.Wait()
}
