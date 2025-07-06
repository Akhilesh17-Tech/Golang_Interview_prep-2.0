package main

import (
	"fmt"
)

func main() {
	arr := []int{9, 9, 9, 9}
	fmt.Println(addOne(arr)) // Output should be [1, 2, 4, 0, 0]
}

func addOne(arr []int) []int {
	n := len(arr)
	carry := 1

	for i := n - 1; i >= 0; i-- {
		sum := arr[i] + carry
		if sum == 10 {
			arr[i] = 0
			carry = 1
		} else {
			arr[i] = sum
			carry = 0
		}
	}

	if carry == 1 {
		arr = append([]int{1}, arr...) // prepend element variadic arguments, arr... copying all the elements to new slice []int{1} and then assigning back to arr slice
	}
	return arr
}
