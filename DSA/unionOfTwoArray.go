package main

import (
	"fmt"
	"sort"
)

func union(arr1, arr2 []int) []int {
	// Create a map to track unique elements
	elementMap := make(map[int]bool)

	// Add elements from arr1 to the map
	for _, value := range arr1 {
		elementMap[value] = true
	}

	// Add elements from arr2 to the map
	for _, value := range arr2 {
		elementMap[value] = true
	}

	// Collect unique elements from the map into a slice
	var result []int
	for key := range elementMap {
		result = append(result, key)
	}

	return result
}

func main() {
	arr1 := []int{1, 2, 3, 4}
	arr2 := []int{3, 4, 5, 6}

	result := union(arr1, arr2)

	sort.Ints(result)

	fmt.Println(result) // Output: [1 2 3 4 5 6]
}
