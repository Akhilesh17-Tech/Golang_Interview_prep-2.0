package main

import "fmt"

func main() {
	// Initialize the counts for 0, 1, and 2
	count_0 := 0
	count_1 := 0
	count_2 := 0

	arr := []int{2, 0, 1, 1, 2, 0}

	// First pass: count the occurrences of 0, 1, and 2
	for i := 0; i < len(arr); i++ {
		switch arr[i] {
		case 0:
			count_0++ // Increment count of 0s
		case 1:
			count_1++ // Increment count of 1s
		case 2:
			count_2++ // Increment count of 2s
		}
	}

	// Append the correct number of 0s
	for i := 0; i < count_0; i++ {
		arr[i] = 0
	}

	// Append the correct number of 1s
	for i := count_0; i < count_0+count_1; i++ {
		arr[i] = 1
	}

	// Append the correct number of 2s
	for i := count_0 + count_1; i < len(arr); i++ {
		arr[i] = 2
	}

	fmt.Println(arr)

	arr1 := []int{1, 1, 2, 0, 2, 1, 2, 0, 0, 2, 1} // length = 11 [1,1,1,0,2,0,0,2,2], [0,1,1,1,2,0,0,2,2] [0,1,1,1,2,0,0,2,2], [0,1,1,1,0,0,2,2,2], [0,0,1,1,1,0,2,2,2]

	sortArray(arr1)

	fmt.Println(arr1)
}

// Dutch National Flag algorithm
func sortArray(arr []int) {
	low := 0
	mid := 0             // after first iteration 1 2
	high := len(arr) - 1 // 11

	for mid <= high { // 0 < 11  , 1 < 11, 2 < 11, 2 < 9 3 < 9, 4 < 9, 4 < 8, 4 < 7 5 < 7
		if arr[mid] == 0 { //  1 == 0, 1==0 , 2 == 0 , 1==0 , 0 == 0,0 == 0
			arr[low], arr[mid] = arr[mid], arr[low] // a[0], a[3] = a[3], a[0] = > a[0] = 0, a[3] = 1 // a[1], a[4] = a[4],a[1], a[1] = 0, a[4] = 1 
			low++                                   // 1 2
			mid++                                   // 4 5
		} else if arr[mid] == 1 { // 1 == 1, 1==1 , 2==1, 1==1
			mid++ // 1 2 3 already 5 -> 6
		} else {
			arr[mid], arr[high] = arr[high], arr[mid] // a[2], a[10] = a[10], a[2] = > a[2] = 1, a[10] = 2 //a[4] ,a[9] = a[9],a[4] = > a[4] = 2, a[9] = 2 // a[4] ,a[8] = a[8], a[4] = > a[4] = 0, a[8] = 2
			high--                                    // 9 8 7
		}
	}
}
