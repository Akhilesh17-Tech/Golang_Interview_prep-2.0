package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 0, 2, 13, 4, 25, 6, 3, 4, 5}

	max := 0
	currLen := 0

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[i+1] {
			currLen++
		} else {
			if currLen > max {
				max = currLen
			}
			currLen = 0
		}
	}

	// After the loop, check if the last subarray was the longest
	if currLen > max {
		max = currLen
	}

	// Add 1 to max to account for the first element in the subarray
	fmt.Println(max + 1)

}
