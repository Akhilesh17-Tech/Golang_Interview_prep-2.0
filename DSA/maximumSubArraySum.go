package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{-2, -3, 4, -1, -2, 1, 5, -3}
	maxSum := math.MinInt32
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			subArray := arr[i : j+1]
			fmt.Println(subArray)
			sum := 0
			for k := i; k <= j; k++ {
				sum += arr[k]
			}
			if sum > maxSum {
				maxSum = sum
			}
		}

	}
	fmt.Println(maxSum)

	// O(N^2)
	for i := 0; i < len(arr); i++ {
		sum := 0
		for j := i; j < len(arr); j++ {
			sum += arr[j]
			if sum >= 0 && sum > maxSum {
				maxSum = sum
			}
		}

	}
	fmt.Println(maxSum)
	OptimalSolution()
}

// Kadan's algorithm to find the maximum sub array sum
func OptimalSolution() {
	arr := []int{-2, -3, 4, -1, -2, 1, 5, -3}
	maxSum := math.MinInt32
	start := 0
	ansStart := -1
	ansEnd := -1
	currentSum := 0
	for i, v := range arr {
		if currentSum == 0 {
			start = i
		}
		currentSum += v
		if currentSum < 0 {
			currentSum = 0
		}
		if v > maxSum {
			maxSum = v
		}
		if currentSum > maxSum {
			maxSum = currentSum
			ansStart = start
			ansEnd = i
		}
	}

	fmt.Println(maxSum)
	fmt.Println(ansStart, ansEnd)

}
