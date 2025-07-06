package main

import (
	"fmt"
	"sort"
)

func sumop(arr []int) int {
	// Sort the array in ascending order
	sort.Ints(arr)

	totalSum := 0
	products := []int{}

	// Perform operations until there are fewer than two elements left
	for len(arr) > 1 {
		// Take the two largest elements (last two in a sorted ascending order)
		a := arr[len(arr)-1]
		b := arr[len(arr)-2]

		// Calculate their product
		product := a * b

		// Store the product
		products = append(products, product)

		// Remove the last two elements
		arr = arr[:len(arr)-2]
	}

	// Sum up the remaining element in the array (if any)
	for _, num := range arr {
		totalSum += num
	}

	// Sum up all the products
	for _, product := range products {
		totalSum += product
	}

	return totalSum
}

func main() {
	arr := []int{1, 3, 5, 4, 2}
	fmt.Println("Maximum Sum:", sumop(arr)) // Output: 27
}
