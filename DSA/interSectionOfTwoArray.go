package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 6, 7, 8, 9}
	arr2 := []int{3, 4, 5, 7, 9, 10}
	fmt.Println(intersectionOfTwoArray(arr1, arr2))

}

func intersectionOfTwoArray(arr1, arr2 []int) []int {

	elements := make(map[int]bool)
	result := []int{}

	for _, v := range arr1 {
		elements[v] = true
	}

	for _, v := range arr2 {
		if elements[v] {
			result = append(result, v)
			delete(elements, v)
		}
	}

	return result

}
