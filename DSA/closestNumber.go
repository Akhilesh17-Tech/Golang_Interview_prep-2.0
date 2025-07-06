// package main

// import (
// 	"fmt"
// 	"math"
// )

// // Function to find the closest number in an array to a target
// func findClosest(arr []int, target int) int {
// 	if len(arr) == 0 {
// 		panic("Array is empty")
// 	}

// 	closest := arr[0]

// 	for _, num := range arr {
// 		// Check if the current number is closer to the target
// 		if math.Abs(float64(num-target)) < math.Abs(float64(closest-target)) {
// 			closest = num
// 		}
// 	}

// 	return closest
// }

// func main() {
// 	arr := []int{2,12,45,1, 3,56,71, 7, 8, 9, 10, 15}
// 	target := 5

// 	closest := findClosest(arr, target)
// 	fmt.Printf("The closest number to %d is %d\n", target, closest)
// }

package main

import (
	"fmt"
	"math"
	"sort"
)

// Function to find the closest number in a sorted array to a target
func findClosestTwoPointers(arr []int, target int) int {
	// Initialize two pointers
	left, right := 0, len(arr)-1
	closest := arr[0]
	minDiff := math.MaxFloat64
	sort.Ints(arr)

	// Loop until the pointers cross
	for left <= right {
		mid := left + (right-left)/2
		diff := math.Abs(float64(arr[mid] - target))

		// Update closest if we find a smaller difference
		if diff < minDiff {
			minDiff = diff
			closest = arr[mid]
		}

		// Move pointers
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return closest
}

func main() {
	arr := []int{2,12,45,1, 3,56,71, 7, 8, 9, 10, 15}
	target := 5
	closest := findClosestTwoPointers(arr, target)
	fmt.Printf("The closest number to %d is %d\n", target, closest)
}
