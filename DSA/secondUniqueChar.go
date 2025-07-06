package main

import (
	"fmt"
)

func main() {
	str := "nfbcagvuyfabcg"

	// Create a slice of size 26 to store counts for each letter ('a' to 'z')
	count := make([]int, 26)

	// Iterate over each character in the string to count occurrences
	for _, v := range str {
		count[v-'a']++
	}

	// Variable to track how many unique characters we have found
	uniqueCount := 0

	// Iterate over the string again to find the second unique character
	for _, v := range str {
		if count[v-'a'] == 1 {
			uniqueCount++
			if uniqueCount == 2 {
				fmt.Printf("Second unique character: %c\n", v)
				break
			}
		}
	}
}
