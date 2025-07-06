package main

import (
	"fmt"
	"strconv"
)

func compressString(str string) string {
	if len(str) == 0 {
		return str
	}

	var compressed string
	count := 1

	for i := 0; i < len(str); i++ {
		if i < len(str)-1 && str[i] == str[i+1] {
			count++
		} else {
			compressed += string(str[i]) + strconv.Itoa(count)
			count = 1
		}
	}

	if len(compressed) >= len(str) {
		return str
	}

	return compressed
}

func main() {
	fmt.Println(compressString("aabcccccaaa"))
	fmt.Println(compressString("abc"))
}
