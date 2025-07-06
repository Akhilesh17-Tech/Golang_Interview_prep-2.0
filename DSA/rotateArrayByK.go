package main

import (
	"fmt"
)

// RotateLeft rotates the array to the left by k positions
func RotateLeft(arr []int, k int) []int {
	n := len(arr)
	k = k % n // In case k is greater than array size
	reverse(arr, 0, k-1)
	reverse(arr, k, n-1)
	reverse(arr, 0, n-1)
	return arr
}

// RotateRight rotates the array to the right by k positions
func RotateRight(arr []int, k int) []int {
	n := len(arr)
	k = k % n // In case k is greater than array size
	reverse(arr, 0, n-k-1)
	reverse(arr, n-k, n-1)
	reverse(arr, 0, n-1)
	return arr
}

// reverse reverses elements from start to end indices
func reverse(arr []int, start, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6,7}
	k := 3

	fmt.Println("Original array:", arr)

	rotatedLeft := RotateLeft(append([]int(nil), arr...), k) // Use a copy to preserve original array
	fmt.Println("Array after left rotation by", k, "positions:", rotatedLeft)

	rotatedRight := RotateRight(append([]int(nil), arr...), k) // Use a copy to preserve original array
	fmt.Println("Array after right rotation by", k, "positions:", rotatedRight)
}
