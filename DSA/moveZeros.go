package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 6, 8, 0, 0, 1, 0, 0, 0, 1, 1, 1, 0, 0}
	moveZeros(arr)
	fmt.Println(arr)
}

func moveZeros(arr []int) {
	// sort.Slice(arr, func(i, j int) bool {
	// 	return arr[i] > arr[j]
	// })

	left := 0
	right := len(arr) - 1

	for left <= right {
		if arr[left] < 1 && arr[right] > 0 {
			temp := arr[left]
			arr[left] = arr[right]
			arr[right] = temp
			left++
			right--
		} else if arr[left] > 0 {
			left++

		} else if arr[right] == 0 {
			right--
		} else {
			left++
			right--
		}
	}

}
