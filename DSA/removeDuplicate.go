package main

import (
	"fmt"
	"sort"
)

func removeDuplicate(arr []int) []int {
	var result = []int{}
	// short array using method
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		if i == 0 || arr[i] != arr[i-1] {
			result = append(result, arr[i])
		}
	}
	return result

}

func main() {
	arr := [...]int{4, 4, 3, 4, 4, 4, 5, 5, 1, 2, 5, 6, 8, 5, 6, 7}

	res := removeDuplicate(arr[:])

	fmt.Println(res)
}
