package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	// Convert the string to lowercase and remove spaces
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")
	
	// Initialize two pointers
	left, right := 0, len(s)-1

	// Compare characters while moving the pointers towards the center
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	str := "A man a plan a canal Panama"
	// str = "nayan"
	if isPalindrome(str) {
		fmt.Printf("'%s' is a palindrome.\n", str)
	} else {
		fmt.Printf("'%s' is not a palindrome.\n", str)
	}
}
