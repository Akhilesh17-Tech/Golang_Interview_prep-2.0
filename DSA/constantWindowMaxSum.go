package main

import "fmt"

func maxSumOfContiguousElements(arr []int, k int) int {
	if len(arr) < k {
		fmt.Println("Array length is less than the window size.")
		return 0
	}

	// Calculate the initial sum of the first window of size k
	maxSum := 0
	for i := 0; i < k; i++ {
		maxSum += arr[i]
	}

	// Set the current sum to the maxSum initially
	currentSum := maxSum

	// Slide the window over the array
	for i := k; i < len(arr); i++ {
		// Slide the window: subtract the element going out and add the element coming in
		currentSum += arr[i] - arr[i-k]

		// Update maxSum if the current sum is higher
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}

func main() {
	arr := []int{-1, 2, 3, 3, 4, 5, -1} // Example array
	k := 4                              // Window size

	maxSum := maxSumOfContiguousElements(arr, k)
	fmt.Printf("Maximum sum of %d contiguous elements is: %d\n", k, maxSum)

	maxSum1 := maxSumOfContiguousElements1(arr, k)
	fmt.Printf("Maximum sum of %d contiguous elements is: %d\n", k, maxSum1)
}

// better approach

func maxSumOfContiguousElements1(arr []int, k int) int {
	if len(arr) < k {
		fmt.Println("Array length is less than the window size.")
		return 0
	}

	maxSum := 0

	// Iterate through the array to find all subarrays of size k
	for i := 0; i <= len(arr)-k; i++ {
		// Calculate the sum of the current subarray of size k
		currentSum := 0
		for j := i; j < i+k; j++ {
			currentSum += arr[j]
		}

		// Update maxSum if the current sum is higher
		if i == 0 || currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}
