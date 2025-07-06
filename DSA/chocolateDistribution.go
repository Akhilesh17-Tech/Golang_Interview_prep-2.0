// Chocolate Distribution Problem
// Given an array A[ ] of positive integers of size N, where each value represents the number of chocolates in a packet. Each packet can have a variable number of chocolates. There are M students, the task is to distribute chocolate packets among M students such that :
// 1. Each student gets exactly one packet.
// 2. The difference between maximum number of chocolates given to a student and minimum number of chocolates given to a student is minimum.

// Example 1:

// Input:
// N = 8, M = 5
// A = {3, 4, 1, 9, 56, 7, 9, 12}
// Output: 6
// Explanation: The minimum difference between maximum chocolates and minimum chocolates is 9 - 3 = 6 by choosing following M packets :{3, 4, 9, 7, 9}.

package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	arr := []int{3, 4, 1, 9, 56, 7, 9, 12}
	m := 5
	fmt.Println(findMinDiff(arr, m))
}

func findMinDiff(arr []int, m int) int {
	sort.Ints(arr) // 1,2,3,4,6,7,9
	minDiff := math.MaxInt64
	for i := 0; i < len(arr)-m+1; i++ {
		minElement := arr[i]
		maxElement := arr[i+m-1]
		diff := maxElement - minElement
		if diff < minDiff {
			minDiff = diff
		}

	}
	return minDiff

}
