// Panic in Go
// A panic can be triggered in a few different ways:

// Explicitly: Using the panic() function.
// Implicitly: Due to an unexpected error such as an out-of-bounds array access or a nil pointer dereference.

//In Go, a panic is a runtime error that occurs when a program encounters an unexpected situation, such as an out-of-range value, a nil pointer dereference, or a division by zero. When a panic occurs, the program's execution is halted, and the error is propagated up the call stack until it is caught by a deferred function or the program terminates.

//To recover from a panic, you can use the recover function, which returns the value passed to the panic function. The recover function can only be called within a deferred function.

package main

import "fmt"

func main() {
	fmt.Println("Starting the program...")

	// Defer a function that will recover from a panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// Trigger a panic
	panic("Something went wrong!")

	fmt.Println("This line will not be executed because of the panic.")
}
