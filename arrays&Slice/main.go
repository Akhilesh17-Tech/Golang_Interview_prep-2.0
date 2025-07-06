// Arrays
// Fixed Size: The size is part of the type and cannot be changed after declaration.
// Value Type: Assigning an array to another copies all elements (deep copy).
// Memory: All elements are stored together in memory.
// Syntax: [n]Type (e.g., [5]int).


// Slices
// Dynamic Size: Can grow or shrink using append.
// Reference Type: Assigning a slice to another variable points to the same underlying array (shallow copy).
// Memory: Slice is a descriptor with a pointer to an underlying array, length, and capacity.
// Syntax: []Type (e.g., []int).


package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello, World!")

	//creating an array with fixed size value type
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr := array

	array[0] = 10000
	arr[0] = 2000

	fmt.Println(array)
	fmt.Println(arr)

	// slice is reference type in go
	slice1 := []int{1, 2, 3, 4}
	slice2 := slice1
	slice1[0] = 100
	slice2[1] = 200

	fmt.Println(slice1)
	fmt.Println(slice2)

	// for loop range function
	for index, value := range array {
		fmt.Println(index, value)
	}

	// Declare and Initialize an Array with Default Values
	var array1 [5]int
	fmt.Println("array1:", reflect.TypeOf(array1).Kind(), array1)

	// Declare and Initialize an Array with Specific Values
	var array2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println("array2:", array2)

	// Declare and Initialize an Array with Specific Values Using `...`
	var array3 = [...]int{1, 2, 3, 4, 5}
	fmt.Println("array3:", array3)

	// Short Variable Declaration with Array
	array4 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("array4:", array4)

	// Partial Initialization
	var array5 = [5]int{1, 2}
	fmt.Println("array5:", array5)

	// Initialize Specific Elements
	var array6 = [5]int{1: 10, 3: 30}
	fmt.Println("array6:", array6)

	sl := make([]int, 5, 10)
	fmt.Println("sl:", reflect.TypeOf(sl).Kind(), sl)
	sl = append(sl, 10)
	fmt.Println("sl:", cap(sl), sl)

	var arrs = new([5]int)
	var arrs1 = make([]int, 5)
	fmt.Println(*arrs)
	fmt.Println(arrs1)




	 // Shrinking a slice by reslicing
	 slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	 fmt.Println("Original slice:", slice)
 
	 // Shrink to first 4 elements
	 slice = slice[:4] // using reslicing to shrink the slice
	 // This does not change the underlying array, just the slice's view of it.
	 // The underlying array still exists, but the slice now only refers to the first 4 elements.
	 fmt.Println("Shrunk slice:", slice) // Output: [1 2 3 4]





	 // Growing a slice by appending 
	 slice := []int{1, 2, 3, 4}
	 fmt.Println("Original slice:", slice)
 
	 // Grow by adding elements 
	 slice = append(slice, 5, 6)
	 fmt.Println("Grown slice:", slice) // Output: [1 2 3 4 5 6]
     

// 	 Length (len(slice)):
// The number of elements the slice currently holds.

// Capacity (cap(slice)):
// The number of elements in the underlying array, counting from the start of the slice to the end of the array. This is the maximum size the slice can grow before a new array needs to be allocated. 



// =====================
// Arrays & Slices Interview Q&A
// =====================

// 1. What is an array in Go?
// An array is a fixed-size, ordered collection of elements of the same type. Its size is part of its type.

// 2. How do you declare and initialize an array?
// Example: var arr [3]int = [3]int{1, 2, 3}

// 3. Is the size of an array part of its type?
// Yes. [3]int and [4]int are different types.
// Reason: In Go, the array size is part of the type definition. For example, [3]int and [4]int are two distinct types.
// You cannot assign or compare arrays of different sizes, even if their element types are the same.
// Example:
// var a [3]int
// var b [4]int
// a = b // Compile-time error: cannot use b (type [4]int) as type [3]int in

// 4. What happens when you assign one array to another?
// All elements are copied (deep copy).

// 5. Are arrays value types or reference types in Go?
// Value types. Assignment copies the entire array.

// 6. How do you partially initialize an array?
// Example: var arr = [5]int{1, 2} // [1 2 0 0 0]
// If you access the third element (index 2), it returns the default value (0 for int), not an error.
// Example:
// fmt.Println(arr[2]) // output: 0

// 7. How do you initialize specific elements of an array?
// Example: var arr = [5]int{1: 10, 3: 30} // [0 10 0 30 0]

// 8. How do you iterate over an array?
// for i, v := range arr { ... }

// 9. What is the default value of an array element?
// The zero value of the element type (e.g., 0 for int) and false for bool and "" for string.

// 10. How is memory managed for arrays in Go?
// All elements are stored contiguously in memory.

// 11. What is a slice in Go?
// A slice is a dynamically-sized, flexible view into the elements of an array.

// 12. How do you declare and initialize a slice?
// Example: s := []int{1, 2, 3}

// 13. How do slices differ from arrays?
// Slices are dynamic, reference types; arrays are fixed-size, value types.

// 14. Are slices value types or reference types?
// Reference types. Assignment copies the slice header, not the underlying data.
// for example:
// slice1 := []int{1, 2, 3}
// slice2 := slice1
// slice1[0] = 100
// fmt.Println(slice1) // Output: [100 2 3]
// fmt.Println(slice2) // Output: [100 2 3]

// 15. What happens when you assign one slice to another?
// Both refer to the same underlying array.

// 16. How do you append elements to a slice?
// Use append: s = append(s, 4, 5)

// 17. How do you shrink a slice?
// By reslicing: s = s[:n]

// 18. How do you remove an element from a slice?
// s = append(s[:i], s[i+1:]...)

// 19. What are length and capacity of a slice?
// Length: number of elements in the slice. Capacity: number of elements from the start of the slice to the end of the underlying array.

// 20. How does Go handle slice capacity when appending elements?
// If capacity is exceeded, a new underlying array is allocated and elements are copied by the goruntime.

// 21. How do you create a slice with make?
// s := make([]int, length, capacity)

// 22. What is the underlying array of a slice?
// The actual array in memory that the slice references.

// 23. How do you copy a slice?
// Use copy: copy(dst, src)

// 24. What happens if you modify a slice after copying it?
// If you copy with copy(), the slices are independent. If you just assign, both refer to the same data.

// 25. How do you iterate over a slice?
// for i, v := range s { ... }

// 26. What are the key differences between arrays and slices in Go?
// Arrays: fixed size, value type. Slices: dynamic size, reference type.

// 27. When would you use an array vs a slice?
// Use arrays for fixed-size collections, slices for dynamic collections.

// 28. How does memory allocation differ between arrays and slices?
// Arrays allocate all elements up front. Slices allocate a descriptor and reference an underlying array.

// 29. How does passing arrays and slices to functions differ?
// Arrays are copied (value type). Slices pass a reference (changes affect the original data).

// =====================
// End
}





