package main

import "fmt"

// rain water trapping
// maximum water we can trap

func main() {
	arr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	// [0,1,0,2,1,0,1,3,2,1,2,1]
	left := make([]int, len(arr))
	right := make([]int, len(arr))

	left[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		left[i] = max(left[i-1], arr[i]) // auxiliary left array left max
	}

	right[len(arr)-1] = arr[len(arr)-1]
	for i := len(arr) - 2; i >= 0; i-- {
		right[i] = max(right[i+1], arr[i]) // auxiliary right array with right max
	}

	result := 0

	for i := 0; i < len(arr); i++ {
		result += min(left[i], right[i]) - arr[i] // at this position we can store maximum water
	}

	fmt.Println(left)
	fmt.Println(right)
	fmt.Println(result)

}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
