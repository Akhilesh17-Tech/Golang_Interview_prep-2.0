package main

import "fmt"

func main() {

	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	leftMax := 0
	rightMax := 0

	left := 0
	right := len(height) - 1

	res := 0

	for left <= right {
		if height[left] <= height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				res += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				res += rightMax - height[right]
			}
			right--
		}
	}

	fmt.Println("total water can be filled = ", res)
}
