package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 10, 11, 12, 17, 18, 19, 20}
	for _, n := range numbers {
		if isPrime(n) {
			fmt.Println("number is a prime number:", n)
		} else {
			fmt.Println("number is not a prime number:", n)
		}
	}
}
