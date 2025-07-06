package main

import "fmt"

func main() {

	arr := []int{1, 3, 2, 4, 1, 6, 3, 8, 9}

	peak := 0

	for i := 1; i < len(arr)-1; i++ {
		if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
			if arr[i] > peak {
				peak = arr[i]
			}
		}
	}
	fmt.Println(peak)

}
