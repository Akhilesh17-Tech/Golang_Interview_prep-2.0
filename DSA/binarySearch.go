package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 6, 7, 8, 9, 10} //
	target := 5

	fmt.Println(binarySearchWithTargetFoundOrNot(arr, target))
	// fmt.Println(binarySearch(arr, 1))

}

func binarySearchWithTargetFoundOrNot(arr []int, target int) int {
	low := 0
	high := len(arr)

	if target > arr[high-1] {
		return high
	}

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}

	}
	return low
}

// requires sorted array
func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
