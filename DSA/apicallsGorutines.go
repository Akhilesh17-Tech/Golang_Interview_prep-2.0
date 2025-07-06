package main

import (
	"fmt"
	"sync"
	"time"
)

// Simulate an API call with a function
func apiCall(name string, wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done()
	// Simulate API delay
	time.Sleep(2 * time.Second)
	result := fmt.Sprintf("Result from %s", name)
	fmt.Println(result)
	if ch != nil {
		ch <- result // Send result to channel if required
	}
}

func processResult(ch <-chan string) {
	result := <-ch
	fmt.Println("Processing:", result)
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan string)

	// Step 1: First API call
	wg.Add(1)
	go func() {
		defer close(ch) // Close the channel when done
		apiCall("API 1", &wg, ch)
	}()

	// Step 2: After first API call finishes, start the second one and process the result
	wg.Add(1)
	go func() {
		processResult(ch)          // Process result from API 1
		apiCall("API 2", &wg, nil) // Run API 2 after processing
	}()

	// Step 3: Independent API 3 and API 4 calls
	wg.Add(2)
	go apiCall("API 3", &wg, nil)
	go apiCall("API 4", &wg, nil)

	// Wait for all goroutines to complete
	wg.Wait()

	fmt.Println("All API calls completed.")
}


