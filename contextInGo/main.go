// In Go, the context package provides a way to manage deadlines, cancellations, and request-scoped values across API boundaries and between processes. Contexts are used to propagate cancellation signals and carry request-scoped data through a program.

// Why Use Context?
// Cancellation: Context allows you to cancel long-running operations when they are no longer needed, saving resources.
// Timeouts and Deadlines: Context can specify timeouts and deadlines, ensuring operations don't run indefinitely.
// Request-Scoped Values: Context can carry values scoped to a request, allowing you to pass request-specific information through your program.
// Key Functions in the context Package
// context.Background(): Returns an empty context. It's typically used at the top level.
// context.TODO(): Returns an empty context when it's unclear which context to use.
// context.WithCancel(parent): Returns a copy of the parent context that can be canceled.
// context.WithDeadline(parent, deadline): Returns a copy of the parent context that is canceled at the specified deadline.
// context.WithTimeout(parent, timeout): Returns a copy of the parent context that is canceled after the specified timeout.
// context.WithValue(parent, key, val): Returns a copy of the parent context with a key-value pair.

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine cancelled")
				return
			default:
				fmt.Println("Working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

// Applications

// Handling Request Timeouts : When making network requests or performing operations that may take time, you can use a context with a timeout to ensure that the operation does not exceed a specified duration.
// Cancelling Long-Running Operations : You can use context to cancel long-running operations when they are no longer needed or when the user terminates the request.
// Passing Request-Scoped Values : Context is useful for passing request-scoped data (e.g., user ID, authentication tokens) through function calls and across API boundaries.
// Coordinating Multiple Goroutines : When you have multiple goroutines that need to be coordinated, you can use context to signal when all goroutines should stop processing.
// HTTP Request Handling : In web servers, context is often used to handle request cancellation and timeouts, allowing for clean handling of HTTP request lifecycle.
