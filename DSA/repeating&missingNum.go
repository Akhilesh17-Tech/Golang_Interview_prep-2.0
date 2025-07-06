package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 2} // we are supposed to 1 - n numbers that's why second loop starts from 1 not 0 if 0 to n-1 we want to find then it should start with 0 to n 
	fmt.Println(findMissingAndRepeating(arr))

}

func findMissingAndRepeating(arr []int) (int, int) {

	repeating := -1
	missing := -1

	map1 := make(map[int]int)

	for i := 0; i < len(arr); i++ { //[1:1,2:2,3:1]
		map1[arr[i]]++
	}

	for i := 1; i <= len(arr); i++ {
		if map1[i] == 2 {
			repeating = i
		} else if map1[i] == 0 {
			missing = i
		}
		if repeating != -1 && missing != -1 {
			break
		}
	}

	return repeating, missing

}
