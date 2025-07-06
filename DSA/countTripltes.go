// Given an array Arr consisting of N distinct integers. The task is to count all the triplets such that sum of two elements equals the third element.
// Example 1:

// Input:
// N = 4
// arr[] = {1, 5, 3, 2}
// Output: 2
// Explanation: There are 2 triplets:
//  1 + 2 = 3 and 3 +2 = 5

package main

import "fmt"

func main() {
	arr := []int{1, 5, 3, 2}
	countTriplets(arr)
}

func countTriplets(arr []int) {
	// Find the maximum value in the array
	maxVal := 0
	for _, v := range arr {
		if v > maxVal {
			maxVal = v
		}
	}

	// Create an indexes slice with length maxVal + 1
	indexes := make([]int, maxVal+1)

	for _, v := range arr {
		indexes[v] = 1
	}

	fmt.Println(indexes)

	// Count the triplets
	count := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			sum := arr[i] + arr[j]
			if sum <= maxVal && indexes[sum] == 1 {
				count++
			}
		}
	}

	fmt.Printf("Number of triplets: %d\n", count)
}

