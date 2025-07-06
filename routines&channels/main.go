package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter = 0
	mu      sync.Mutex
)

func increment(wg *sync.WaitGroup, i int) {
	defer wg.Done()
	// mu.Lock()
	counter++
	fmt.Println(i)
	// mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(&wg, i)
	}
	wg.Wait()
	fmt.Println("final counter : ", counter)

	// channels buffered
	ch := make(chan int, 5) // Create a buffered channel with capacity 5

	go producer(ch, 10)
	go consumer(ch)

	// channels unbuffered

	ch1 := make(chan int)
	go producer1(ch1)
	go consumer1(ch1)
	time.Sleep(2 * time.Second)

}

func producer(ch chan int, num int) {
	for i := 0; i < num; i++ {
		ch <- i
		fmt.Println("Produced:", i)
	}
	close(ch)
}

func consumer(ch chan int) {
	for {
		select {
		case x, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed")
				return
			}
			fmt.Println("Consumed:", x)
		}
	}
}

// producer using unbuffered channel

func producer1(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("Produced:", i)
	}
	close(ch)
}

func consumer1(ch chan int) {
	for {
		select {
		case x, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed")
				return
			}
			fmt.Println("Consumed:", x)
		}
	}
}





