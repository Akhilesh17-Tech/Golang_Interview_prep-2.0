package main

import "fmt"

func longestSubStringWithoutRepeatingChar(str string) (int, string) {
	maxLen := 0
	var maxSubStr string

	for i := 0; i < len(str); i++ {
		charMap := make(map[byte]int) // Reset the map for each starting position
		subStr := ""

		for j := i; j < len(str); j++ {
			// If the character is already in the map, break the loop
			if charMap[str[j]] == 1 {
				break
			}
			// Add the character to the substring and map
			subStr += string(str[j])
			charMap[str[j]] = 1

			// Update the maximum length and substring if the current one is longer
			if len(subStr) > maxLen {
				maxLen = len(subStr)
				maxSubStr = subStr
			}
		}
	}

	return maxLen, maxSubStr
}

func main() {
	maxlen, subStr := longestSubStringWithoutRepeatingChar("abcabcbb")
	fmt.Println("Length:", maxlen, "Substring:", subStr)

	maxlen1, subStr1 := optimalSolution("abcabcbb")
	fmt.Println("Length:", maxlen1, "Substring:", subStr1)
}

func optimalSolution(str string) (int, string) {
	charMap := make(map[byte]int)

	left, right := 0, 0
	maxLen := 0
	maxSubStr := ""
	length := len(str)

	for right < length {
		// If the character has been seen before and is within the current window
		if index, found := charMap[str[right]]; found && index >= left {
			left = index + 1
		}

		// Update the character's last seen index
		charMap[str[right]] = right

		// Calculate the current window length
		subLen := right - left + 1

		// Update the maximum length and substring if the current one is longer
		if subLen > maxLen {
			maxLen = subLen
			maxSubStr = str[left : right+1]
		}

		right++
	}

	return maxLen, maxSubStr
}
