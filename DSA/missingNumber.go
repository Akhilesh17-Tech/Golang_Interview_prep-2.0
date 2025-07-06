package main

import "fmt"

func main() {
	var arr = []int{1, 2, 3, 4, 5, 7}
	missingNum := findMissingNum(arr)
	fmt.Println("The missing number is:", missingNum)
}

func findMissingNum(arr []int) int {
	// expectedSum := (arr[0] + arr[len(arr)-1]) * (len(arr) + 1) / 2
	n := len(arr) + 1
	expectedSum := (n * (n + 1)) / 2
	actualSum := 0
	for _, num := range arr {
		actualSum += num
	}
	return expectedSum - actualSum
}
