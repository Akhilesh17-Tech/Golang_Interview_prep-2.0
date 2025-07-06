// The pipeline pattern is a design pattern that involves organizing a sequence of processing stages, where each stage receives data, processes it, and passes it along to the next stage. In concurrent programming, the pipeline pattern is particularly effective for structuring tasks that need to be processed in a specific order, allowing for parallel processing and efficient use of resources.

package main

import (
	"fmt"
)

func generator(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	// Setting up the pipeline
	gen := generator(2, 3, 4)
	sq := square(gen)

	// Consuming the output
	for n := range sq {
		fmt.Println(n)
	}
}
