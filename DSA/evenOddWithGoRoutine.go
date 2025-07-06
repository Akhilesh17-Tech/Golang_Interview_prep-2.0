// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func printOdd(oddCh, evenCh chan struct{}, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 1; i <= 10; i += 2 {
// 		<-oddCh
// 		fmt.Println(i)
// 		evenCh <- struct{}{}
// 	}
// }

// func printEven(oddCh, evenCh chan struct{}, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 2; i <= 10; i += 2 {
// 		<-evenCh
// 		fmt.Println(i)
// 		oddCh <- struct{}{}
// 	}
// }

// func main() {
// 	oddCh := make(chan struct{}, 1) // initialize the oddCh channel with a buffer size of 1. This is necessary because the printOdd goroutine is trying to receive from oddCh before anything is sent to it. By initializing the channel with a buffer size of 1, we ensure that the first send operation on oddCh will not block.

// 	evenCh := make(chan struct{})
// 	var wg sync.WaitGroup

// 	wg.Add(2)
// 	go printOdd(oddCh, evenCh, &wg)
// 	go printEven(oddCh, evenCh, &wg)

// 	oddCh <- struct{}{} // Start the sequence with odd numbers
// 	wg.Wait()
// }

package main

import (
	"fmt"
	"sync"
)

func main() {

	// By initializing oddCh with a capacity of 1, you ensure that the initial send operation does not block, allowing the oddNumber goroutine to start first
	oddCh := make(chan bool, 1) // buffered with capacity 1 asynchronous natures it won't block receiver channel and vice versa, oddCh is buffered with capacity 1, this operation does not block. The oddNumber goroutine can now start its work without having to wait for the evenNumber goroutine to be ready.
	evenCh := make(chan bool)   // unbuffered without capacity synchronous nature it will block receiver if sender hasn't send or vice versa

	var wg sync.WaitGroup

	wg.Add(2)

	go oddNumber(oddCh, evenCh, &wg)

	go evenNumber(oddCh, evenCh, &wg)

	oddCh <- true

	wg.Wait()

}

func evenNumber(oddCh, evenCh chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		<-evenCh
		fmt.Println(i)
		oddCh <- true
	}
}

func oddNumber(oddCh, evenCh chan bool, wg *sync.WaitGroup) {
	defer wg.Done() // defer used for resource cleanup also decrement the count of go routine after complition of current function execution
	for i := 1; i <= 10; i += 2 {
		<-oddCh
		fmt.Println(i)
		evenCh <- true
	}

}

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	oddCh := make(chan bool)  // Unbuffered channel for signaling odd number printing
// 	evenCh := make(chan bool) // Unbuffered channel for signaling even number printing

// 	// Goroutine for printing odd numbers
// 	go func() {
// 		for i := 1; i <= 9; i += 2 {
// 			<-oddCh        // Wait for the signal from the main function or the even number goroutine
// 			fmt.Println(i) // Print the odd number
// 			evenCh <- true // Signal the even number goroutine to continue
// 		}
// 		close(evenCh) // Close the evenCh channel when done
// 	}()

// 	// Goroutine for printing even numbers
// 	go func() {
// 		for i := 2; i <= 10; i += 2 {
// 			<-evenCh
// 			fmt.Println(i) // Print the even number
// 			oddCh <- true  // Signal the odd number goroutine to continue
// 		}
// 		close(oddCh) // Close the oddCh channel when done
// 	}()

// 	// Start the sequence with the odd number goroutine
// 	oddCh <- true

// 	time.Sleep(2 * time.Second)
// }
