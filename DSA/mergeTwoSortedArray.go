package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1 := []int{11, 1, 3, 5, 7, 9}
	arr2 := []int{14, 2, 4, 6, 8, 10, 12, 13}
	sort.Ints(arr1)
	sort.Ints(arr2)
	mergedArray := mergeSortedArrays(arr1, arr2)
	fmt.Println(mergedArray)

	n := len(arr1)
	m := len(arr2)

	mergeSortWithoutExtraSpace(arr1, arr2, n, m)
	fmt.Println(arr1, arr2)
}

// using extra space
func mergeSortedArrays(arr1 []int, arr2 []int) []int {
	merged := make([]int, 0, len(arr1)+len(arr2)) // Initialize with capacity, not length
	left, right := 0, 0

	// Merge the arrays
	for left < len(arr1) && right < len(arr2) {
		if arr1[left] < arr2[right] {
			merged = append(merged, arr1[left])
			left++
		} else {
			merged = append(merged, arr2[right])
			right++
		}
	}

	// Append any remaining elements from arr1
	for left < len(arr1) {
		merged = append(merged, arr1[left])
		left++
	}

	// Append any remaining elements from arr2
	for right < len(arr2) {
		merged = append(merged, arr2[right])
		right++
	}

	return merged
}

// better approach without using extra space

// Input:

// n = 4, arr1[] = [1 4 8 10]
// m = 5, arr2[] = [2 3 9]

// Output:

// arr1[] = [1 2 3 4]
// arr2[] = [8 9 10]

func mergeSortWithoutExtraSpace(arr1 []int, arr2 []int, n int, m int) {
	left := n - 1
	right := 0

	// Swap elements if needed to maintain order across both arrays
	for left >= 0 && right < m {
		if arr1[left] > arr2[right] {
			arr1[left], arr2[right] = arr2[right], arr1[left]
		}
		left--
		right++
	}

	// Sort both arrays again to ensure both are sorted
	sort.Ints(arr1)
	sort.Ints(arr2)
}
