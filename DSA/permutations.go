package main

import "fmt"

func permute(str []rune, l, r int) {
	if l == r {
		fmt.Println(string(str))
	} else {
		for i := l; i <= r; i++ {
			str[l], str[i] = str[i], str[l] // Swap characters at index l and i
			permute(str, l+1, r)            // Recur for the next index
			str[l], str[i] = str[i], str[l] // Backtrack: swap back to the original configuration
		}
	}
}

func main() {
	str := "ABCD"
	permute([]rune(str), 0, len(str)-1)
}
