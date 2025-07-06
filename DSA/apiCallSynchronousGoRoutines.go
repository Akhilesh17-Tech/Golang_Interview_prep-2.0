// synchronous communication between go routines

package main

import (
	"fmt"
	"sync"
	"time"
)

// Simulate an API call with a function
func apiCall(name string, next chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	// Simulate API delay
	time.Sleep(2 * time.Second)
	fmt.Println("Result from", name)
	if next != nil {
		// Signal the next goroutine to start
		next <- struct{}{}
		close(next)
	}
}

func processResult(result string) {
	fmt.Println("Processing:", result)
}

func main() {
	var wg sync.WaitGroup

	// Create channels to control the execution flow
	firstDone := make(chan struct{})
	secondDone := make(chan struct{})

	// Step 1: Run the first API call in a goroutine
	wg.Add(1)
	go func() {
		apiCall("API 1", firstDone, &wg)
	}()

	// Step 2: Process result and then trigger the second API call after the first one completes
	wg.Add(1)
	go func() {
		<-firstDone
		processResult("Result from API 1")
		apiCall("API 2", secondDone, &wg)
	}()

	// Step 3: Run API 3 and API 4 independently after API 2 starts
	wg.Add(2)
	go func() {
		<-secondDone
		apiCall("API 3", nil, &wg)
	}()
	go func() {
		<-secondDone
		apiCall("API 4", nil, &wg)
	}()

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All API calls completed.")
}
