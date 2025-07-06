// Given an unsorted array arr of size n that contains only non negative integers, find a sub-array (continuous elements) that has sum equal to s. You mainly need to return the left and right indexes(1-based indexing) of that subarray.

// In case of multiple subarrays, return the subarray indexes which come first on moving from left to right. If no such subarray exists return an array consisting of element -1.

// Examples:

// Input: arr[] = [1,2,3,7,5], n = 5, s = 12
// Output: 2 4
// Explanation: The sum of elements from 2nd to 4th position is 12.

// required
// sum, flag, n , start : sum=0, flag=0, n=len(arr), start=0

package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 7, 5} // 1,3
	target := 12
	subArraySum(arr1, target)
}

func subArraySum(arr []int, target int) {
	n := len(arr)
	flag := 0
	start := 0
	sum := 0
	for i := 0; i < n; i++ {
		sum = sum + arr[i]
		for sum > target {
			sum = sum - arr[start]
			start++
		}
		if sum == target {
			flag = 1
			fmt.Println([]int{start, i})
			break
		}

	}
	if flag == 0 {
		fmt.Println(-1)
	}
}
