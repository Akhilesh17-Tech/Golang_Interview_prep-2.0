package main

import "fmt"

func fibonacci(n int) []int {
	fibSeq := make([]int, n)
	if n >= 1 {
		fibSeq[0] = 0
	}
	if n >= 2 {
		fibSeq[1] = 1
	}
	for i := 2; i < n; i++ {
		fibSeq[i] = fibSeq[i-1] + fibSeq[i-2]
	}
	return fibSeq
}

func main() {
	n := 10
	fibSeq := fibonacci(n)
	fmt.Println(fibSeq)
}
