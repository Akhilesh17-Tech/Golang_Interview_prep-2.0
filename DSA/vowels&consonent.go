package main

import (
	"fmt"
	"strings"
)

func countVowelsAndConsonents(s string) {
	vowels := "aeiou"
	vowelsMap := make(map[string]int)
	consonentsMap := make(map[string]int)

	for _, v := range strings.ToLower(s) {
		if strings.Contains(vowels, string(v)) {
			vowelsMap[string(v)]++
		} else if v >= 'a' && v <= 'z' {
			consonentsMap[string(v)]++
		}
	}

	fmt.Println("Vowels:", vowelsMap)
	fmt.Println("Consonents:", consonentsMap)
}

func main() {
	countVowelsAndConsonents("Akhileshkushwah, World!")
}
