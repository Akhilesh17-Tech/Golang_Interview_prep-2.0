package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu    sync.RWMutex
	value int
}

func (c *SafeCounter) Increment() {
	c.mu.Lock() // exclusive lock for writing
	c.value++
	c.mu.Unlock()
}

func (c *SafeCounter) GetValue() int {
	c.mu.RLock() // shared lock for reading
	defer c.mu.RUnlock()
	return c.value
}

func main() {
	counter := &SafeCounter{}

	// Start multiple goroutines to read and write
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
			time.Sleep(1 * time.Millisecond) // Simulate some processing
		}()
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Current Value:", counter.GetValue())
			time.Sleep(1 * time.Millisecond) // Simulate some processing
		}()
	}

	wg.Wait()
	fmt.Println("Final Value:", counter.GetValue())
}
