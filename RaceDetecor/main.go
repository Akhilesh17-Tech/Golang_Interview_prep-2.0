package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 1000; i++ {
			counter++
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			counter++
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Counter:", counter)
}


// go run -race main.go 