package main

import "fmt"

func findSecondLargest(arr []int) int {
	var max = arr[0]
	var secondMax = arr[0]

	for _, value := range arr {
		if value > max {
			secondMax = max
			max = value
		} else if value > secondMax && secondMax != max {
			secondMax = value
		}
	}
	return secondMax

}

func main() {

	// var arrayName = [length]dataType{values}

	var arr = []int{1, 2, 3, 4, 5, 6, 7, 5}

	result := findSecondLargest(arr)
	fmt.Println(result)

}
