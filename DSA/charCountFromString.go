package main

import (
	"fmt"
	"strings"
)

func main() {
	var name string = "akhilesh"
	// countChar(name)
	countCharacters(name)

}

func countChar(s string) {

	map1 := make(map[string]int)

	for _, v := range s {
		if map1[string(v)] > 0 {
			map1[string(v)] += 1
		} else {
			map1[string(v)] = 1
		}
	}
	fmt.Println(map1)
}

func countCharacters(s string) {
	var count [256]int

	s = strings.ToLower(s)

	for _, char := range s {
		count[char]++
	}

	for i := 0; i < 256; i++ {
		if count[i] > 0 {
			fmt.Printf("'%c' : %d\n", i, count[i])
		}
	}
}
