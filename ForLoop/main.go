package main

import (
	"fmt"
)

func main() {
	fmt.Println("Slice and Array")

	// Arrays:
	// Fixed size.
	// Value type.
	// Length is part of the type.
	// Suitable for cases where the size is known and fixed.
	// len function for array lengths

	//creating an array

	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // array is fixed size when we copy an array into another it creates a new reference

	arr2 := array // arr2 is a copy of arr1
	arr2[0] = 10  // we are modifying arr2 won't affect the original array

	fmt.Println("array", array)
	fmt.Println("array2", arr2)

	// Slices:
	// Dynamic size.
	// Reference type.
	// Length and capacity.
	// More flexible and commonly used in Go.

	// two properties in slice
	// 1: length = Represents the number of elements currently present in the slice.
	// 2: capacity = The capacity of a slice is the number of elements the slice can hold without needing to allocate more memory.
	// 3: cap function is there to check capacity of a slice.

	slice := []int{1, 2, 3, 4, 5, 6} // dynamic shrink and grow size when we copy a slice it refers a reference

	slice2 := slice // slice2 refers to the same underlying array as slice1

	slice2[0] = 10

	fmt.Println("slice2", slice2)
	fmt.Println("slice", slice)

	// for index, value := range array1 {
	// 	fmt.Println(index, value)
	// }

	// for index, value := range slice {
	// 	fmt.Println(index, value)
	// }

	// Array
	arr := [3]int{1, 2, 3}
	fmt.Println("Array:", arr)

	// Slices
	newslice := []int{1, 2, 3}
	fmt.Println("Slice before append:", newslice)

	// Append to slice
	newslice = append(newslice, 4, 5)
	fmt.Println("Slice after append:", newslice)

	// Slicing an array
	arrSlice := arr[0:2]
	fmt.Println("Slice from array:", arrSlice)

	// Modify original array affects the slice
	arr[0] = 10
	fmt.Println("Modified array:", arr)
	fmt.Println("Slice from modified array:", arrSlice)

	// Slices share underlying array
	newslice2 := newslice
	newslice2[0] = 100
	fmt.Println("Original slice:", newslice)
	fmt.Println("Modified slice2:", newslice2)

	// slice with make
	slice5 := make([]int, 3, 5) // Length 3, Capacity 5
	fmt.Println("Slice made with make:", slice5)        // Output: [0 0 0]
    fmt.Println("Slice length:", len(slice5))           // Output: 3
    fmt.Println("Slice capacity:", cap(slice5))         // Output: 5

}
