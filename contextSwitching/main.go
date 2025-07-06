// In Go, context switching primarily occurs within the runtime’s scheduler, which is responsible for managing goroutines, the lightweight threads of execution in Go. Here's how context switching works in the Go programming language:

// Goroutines and the Go Scheduler
// Goroutines: Goroutines are functions or methods that run concurrently with other functions or methods. They are much lighter than traditional OS threads because they are managed by the Go runtime rather than the operating system.

// Go Scheduler: The Go scheduler is responsible for managing the execution of goroutines. It maps goroutines onto a smaller number of OS threads using a mechanism known as M
// scheduling. Here, M goroutines are multiplexed over N OS threads.

package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second) // Yield point
	}
}

func printLetters() {
	for i := 'A'; i <= 'E'; i++ {
		fmt.Println(string(i))
		time.Sleep(1 * time.Second) // Yield point
	}
}

func main() {
	go printNumbers()
	go printLetters()

	time.Sleep(6 * time.Second) // Allow time for goroutines to finish
}

// Explanation:
// Goroutines: In this example, printNumbers and printLetters are each run in their own goroutine.
// Yield Points: The time.Sleep(1 * time.Second) acts as a yield point where the Go scheduler can switch between the two goroutines.
// Context Switching: The Go scheduler will switch between the printNumbers and printLetters goroutines to interleave their output, giving the appearance that both functions are running concurrently.

// go routines over OS threads

// Memory Usage:
// Goroutines: Start with a small stack (~2 KB) that grows and shrinks dynamically.
// OS Threads: Typically require a large fixed stack size (~1 MB).

// Scheduling:
// Goroutines: Managed by Go's runtime with M scheduling, allowing many goroutines per OS thread.
// OS Threads: Managed by the OS kernel with 1:1 scheduling.

// Context Switching:
// Goroutines: Fast, user-space context switching with low overhead.
// OS Threads: Slower, kernel-space context switching with higher overhead.

// Scalability:
// Goroutines: Can easily scale to millions of concurrent tasks.
// OS Threads: Limited by higher resource consumption, scaling poorly for many threads.

// Concurrency Model:
// Goroutines: Simplified with built-in channels for communication.
// OS Threads: Requires explicit management of synchronization and communication.
